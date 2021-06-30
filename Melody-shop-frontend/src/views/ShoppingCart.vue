<template>
  <section class="shopping-cart spad">
    <div class="container">
      <div class="row">
        <div class="col-lg-8">
          <div class="shopping__cart__table">
            <table>
              <thead>
                <tr>
                  <th>Product</th>
                  <th>Quantity</th>
                  <th>Total</th>
                  <th></th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(product, key) in products" :key="key">
                  <CartItem :item="product" :index="key" />
                </tr>
              </tbody>
            </table>
          </div>
          <div class="row">
            <div class="col-lg-6 col-md-6 col-sm-6">
              <div class="continue__btn">
                <router-link to="/shop">Continue Shopping</router-link>
              </div>
            </div>
          </div>
        </div>
        <div class="col-lg-4">
          <div class="cart__discount">
            <h6>Discount codes</h6>
            <div >
              <input type="text" placeholder="Coupon code" v-model="coupon" />
              <button  @click="addCoupon">Apply</button>
            </div>
          </div>
          <div class="cart__total">
            <h6>Cart total</h6>
            <ul>
              <li>
                Subtotal <span>{{ subTotal }}</span>
              </li>
              <li>
                Total <span>{{ totalItems }}</span>
              </li>
            </ul>
            <a href="#" class="primary-btn">Proceed to checkout</a>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
<script>
import { mapState, mapGetters } from "vuex";
import CartItem from "../components/CartItem.vue";
export default {
  name: "ShoppingCart",
  components: {
    CartItem,
  },
  data() {
    return {
      cloneCart: [],
      coupon:"",
    };
  },
  computed: {
    ...mapState("cart", ["products", "isLoading"]),
    ...mapGetters("cart", ["totalItems", "subTotal"]),
  },
  methods: {
    changeQuanti(value, index) {
      let convert = parseInt(value, 10);
      this.cloneCart[index].quantity = convert;
      if (value === "" || convert === 0) {
        this.cloneCart[index].quantity = 1;
      }
    },
    addCoupon() {
      console.log(this.coupon)
      this.$store.dispatch("cart/addCoupon", this.coupon);
    },
  },
  created() {
    this.cloneCart = [...this.products];
  },
};
</script>
