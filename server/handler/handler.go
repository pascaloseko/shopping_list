package handler

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber"
	"github.com/pascaloseko/shopping_list/server/database"
	"github.com/pascaloseko/shopping_list/server/model"
	"github.com/pascaloseko/shopping_list/server/util"
	"golang.org/x/crypto/bcrypt"
)

var user *model.User

// LoginPayload login body
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginResponse struct {
	Token string `json:"token"`
}

// CreateItem handler
func CreateItem(c *fiber.Ctx) {
	item := new(model.Item)
	if err := c.BodyParser(item); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	// insert item to db
	_, err := database.DB.Query("INSERT INTO items (name) VALUES ($1)", item.Name)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	// return item in json format
	if err := c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "Item added successfully",
		"item":    item,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
}

// GetAllItems handler
func GetAllItems(c *fiber.Ctx) {
	rows, err := database.DB.Query("SELECT * FROM items order by name")
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return
	}

	defer rows.Close()
	result := model.Items{}

	for rows.Next() {
		item := model.Item{}
		err := rows.Scan(&item.ID, &item.Name, &item.Date)
		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
			return
		}
		result.Items = append(result.Items, item)
	}

	if err := c.JSON(&fiber.Map{
		"success":  true,
		"products": result,
		"message":  "All items returned successfully",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
}

// GetSingleItem handler
func GetSingleItem(c *fiber.Ctx) {
	id := c.Params("id")
	item := model.Item{}

	row, err := database.DB.Query("SELECT * FROM items WHERE id = $1", id)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	defer row.Close()
	for row.Next() {
		switch err := row.Scan(&item.ID, &item.Name, &item.Date); err {
		case sql.ErrNoRows:
			log.Println("No rows were returned!")
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		case nil:
			log.Println(item.Name, item.Date)
		default:
			//   panic(err)
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}

	if err := c.JSON(&fiber.Map{
		"success": false,
		"message": "Successfully fetched item",
		"item":    item,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
}

// UpdateItem handler
func UpdateItem(c *fiber.Ctx) {
	id := c.Params("id")
	// Instantiate new item struct
	item := new(model.Item)
	//  Parse body into item struct
	if err := c.BodyParser(item); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	// Update item database
	_, err := database.DB.Query("UPDATE items SET name=$1, date=$2 WHERE id = $3", item.Name, item.Date, id)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	// Return item in JSON format
	if err := c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Item successfully updated",
		"item":    item,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating item",
		})
		return
	}
}

// DeleteItem handler
func DeleteItem(c *fiber.Ctx) {
	id := c.Params("id")
	// query item table in database
	_, err := database.DB.Query("DELETE FROM items WHERE id = $1", id)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return
	}
	// return item in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "item deleted successfully",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return
	}
}

// Register	user
func Register(c *fiber.Ctx) {
	user := new(model.User)
	check := isUser()
	if check {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "User already exists",
		})
		return
	}

	if err := c.BodyParser(user); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	// insert user to db
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) returning id, name, register_date"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	defer stmt.Close()
	if err := stmt.QueryRow(user.UserName, user.Email, Encrypt(user.Password)).Scan(&user.ID, &user.UserName, &user.Email, &user.RegisterDate); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	// return item in json format
	if err := c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "User added successfully",
		"item":    user,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
}

// LoginUser handler
func LoginUser(c *fiber.Ctx) {
	userPayload := new(model.User)
	if err := c.BodyParser(userPayload); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	err := database.DB.QueryRow("SELECT id, password, name, register_date FROM users WHERE email=$1", userPayload.Email).Scan(&user.ID, &user.UserName, &user.RegisterDate)
	if err != nil {
		c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userPayload.Password))
	if err != nil {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "invalid user credentials",
		})
		return
	}
	jwtWrapper := util.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "error signing token",
		})
		return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": tokenResponse,
	})

	return
}

// IsUser check if user is registered
func isUser() (available bool) {
	var num int
	database.DB.QueryRow("SELECT COUNT(*) from users WHERE email=$1 LIMIT 1", user.Email).Scan(&num)
	if num == 0 {
		available = false
	}
	available = true
	return
}

// Encrypt encypts a string with sha1 algorithm
func Encrypt(plaintext string) (cryptoText string) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	cryptoText = string(hashPassword)
	return
}
