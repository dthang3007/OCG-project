import axios from "../axios";


export default {
  async login({ username, email, password }) {
    return axios
      .post(`/authen/login`, { username, email, password })
      .then((response) => {
        return response.data;
      });
  },

  async register(user) {
    return axios.post(`/authen/register`, user).then((response) => {
      return response.data;
    });
  },
};
