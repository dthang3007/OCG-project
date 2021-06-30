import axios from "../axios";

export default {
  async getProducts({ page, limit, sort, order, search, categoryId }) {
    return axios
      .get(
        `/product?page=${page}&limit=${limit}&sort=${sort}&order=${order}&q=${search}&category=${categoryId}`
      )
      .then((response) => {
        console.log(response);
        return {
          data: response.data,
        };
      });
  },

  async getCategories() {
    return axios.get(`/categories`).then((response) => {
      return response.data;
    });
  },

  async getProductById(productId) {
    return axios.get(`/product/${productId}`).then((response) => {
      return response.data;
    });
  },
};
