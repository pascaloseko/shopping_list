import axios from "axios";
import { IItem } from "../../types/interfaces";
import { tokenConfig } from "./authActions";
import { returnErrors } from "./errorActions";
import { GET_ITEMS, ADD_ITEM, DELETE_ITEM, ITEMS_LOADING } from "./types";

export const getItems = () => (
  dispatch: Function,
  getState: Function
  ) => {
  dispatch(setItemsLoading());
  axios
    .get("/api", tokenConfig(getState))
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

export const deleteItem = (id: number) => (
  dispatch: Function,
  getState: Function
  ) => {
    console.info(id)
  axios.delete(`/api/${id}`, tokenConfig(getState)).then((res) => {
    dispatch({
      type: DELETE_ITEM,
      payload: id,
    });
  });
};

export const addItem = (item: IItem) => (
  dispatch: Function,
  getState: Function
  ) => {
  axios
    .post("/api", item, tokenConfig(getState))
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
