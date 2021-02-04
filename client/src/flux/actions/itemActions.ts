import axios from "axios";
import { IItem } from "../../types/interfaces";
import { returnErrors } from "./errorActions";
import { GET_ITEMS, ADD_ITEM, DELETE_ITEM, ITEMS_LOADING } from "./types";

export const getItems = () => (dispatch: Function) => {
  dispatch(setItemsLoading());
  axios
    .get("/api")
    .then((res) =>
      dispatch({
        type: GET_ITEMS,
        payload: res.data,
      })
    )
    .catch((err) => {
      dispatch(returnErrors(err.response.data, err.response.status));
    });
};

export const deleteItem = (id: number) => (dispatch: Function) => {
  axios.delete(`/api/${id}`).then((res) => {
    dispatch({
      type: DELETE_ITEM,
      payload: id,
    });
  });
};

export const addItem = (item: IItem) => (dispatch: Function) => {
  axios
    .post("/api", item)
    .then((res) => {
      dispatch({ type: ADD_ITEM, payload: res.data.item });
    })
    .catch((err) => {
      dispatch(returnErrors(err.response.data, err.response.status));
    });
};

export const setItemsLoading = () => {
  return {
    type: ITEMS_LOADING,
  };
};
