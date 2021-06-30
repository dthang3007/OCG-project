<template>
  <div class="product__item">
    <div style="w" class="card bg-dark text-white border-0" @click="toDetailPage">
      <img  class="card-img w-100" :src="data.images[0]?.src" alt="Card image" />
      <div class="card-img-overlay"></div>
    </div>
    <div class="product__item__text">
      <h6>{{ data.name }}</h6>
      <button class="add-cart" @click="addProductToCart">+ Add To Cart</button>
      <div class="rating"></div>
      <h5>{{ currency(data.price) }}</h5>
    </div>
  </div>
</template>
<script>

import {currency} from "../util"
export default {
  
  name: "ProductItem",
  props: ["data"],

  methods: {
    toDetailPage() {
      this.$router.push({
        name: "ShopDetail",
        params: { name: this.data.name, id: this.data.ID },
      });
    },
    addProductToCart() {
     this.$store.dispatch("cart/addProductToCart", this.data);
    },
    currency

    // addToCart(item) {
    //   if (!localStorage.getItem("cart")) {
    //     localStorage.setItem("cart", JSON.stringify([]));
    //   }
    //   let cartItems = JSON.parse(localStorage.getItem("cart"));
    //   if (!cartItems.length) {
    //     item.quantity = 1;
    //     cartItems.push(item);
    //     localStorage.setItem("cart", JSON.stringify(cartItems));
    //     return;
    //   }
    //   for (let i = 0; i < cartItems.length; i++) {
    //     // console.log(item.ID)
    //     // console.log(cartItems)
    //     if (cartItems[i].ID == item.ID) {

    //       cartItems[i].quantity += 1;
    //       localStorage.setItem("cart", JSON.stringify(cartItems));
    //       return;
    //     }
    //   }
    //   console.log(item)
    //   item.quantity = 1;
    //   cartItems.push(item);
    //   localStorage.setItem("cart", JSON.stringify(cartItems));

    //   return;
    // },
  },
};
</script>
