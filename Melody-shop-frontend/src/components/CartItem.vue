<template>
  <td class="product__cart__item">
    <div class="product__cart__item__pic">
      <img class="w-100 rounded" :src="item.images[0].src" alt="" />
    </div>
    <div class="product__cart__item__text">
      <h6>{{ item.name }}</h6>
      <h5>{{ item.price }}</h5>
    </div>
  </td>
  <td class="quantity__item">
    <div class="quantity">
      <div class="pro-qty-2">
        <input
          type="text"
          :value="item.quantity"
          @input="changeQuantity($event.target.value, index)"
          min="1"
          max="10"
        />
      </div>
    </div>
  </td>
  <td class="cart__price">{{ item.totalPrice }}</td>
  <td class="cart__close"><button class="btn" @click="deleteProduct(index)">X</button></td>
</template>
<script>
export default {
  name: "CartItem",
  props: ["item","index"],
  methods: {
    changeQuantity(value, index) {
      this.$store.dispatch("cart/updateProduct", {value,index});
    },
    deleteProduct(index){
        this.$store.dispatch("cart/deleteProduct",index)
    }
  },
};
</script>
