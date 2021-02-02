package handler

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber"
	"github.com/pascaloseko/shopping_list/server/database"
	"github.com/pascaloseko/shopping_list/server/model"
)

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
		log.Println(err)
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
