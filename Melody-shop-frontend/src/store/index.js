import { createStore } from "vuex";
import users from "./modules/users";
import cart from "./modules/cart.module"
import products from "./modules/products.module"

const debug = process.env.NODE_ENV !== "production";

const store = createStore({
  strict: debug,
  modules: {
    users,
    cart,
    products,
  },
});

export default store;
