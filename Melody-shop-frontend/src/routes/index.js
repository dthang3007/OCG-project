import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home";
import Shop from "../views/Shop";
import About from "../views/About";
import ShopDetail from "../views/ShopDetail";
import Login from "../components/Login";
import Register from "../components/Register";
import ShoppingCart from "../views/ShoppingCart"
const routes = [
  { path: "/", component: Home },
  { path: "/shop", component: Shop },
  { path: "/about", component: About },
  { path: "/login", component: Login },
  { path: "/register", component: Register },
  { path: "/detail/:id", component: ShopDetail, name: "ShopDetail" },
  {path:"/shopping-cart",component:ShoppingCart}
];
const router = createRouter({
  history: createWebHistory(),
  scrollBehavior(){
    return {
      top:0
    }
  },
  routes,
  linkActiveClass: "active",
  linkExactActiveClass: "exact-active",
});
export default router;
