<template>
  <div
    class="offcanvas-menu-overlay"
    :class="{ active: isActive }"
    @click="toogleCanvas"
  ></div>
  <div class="offcanvas-menu-wrapper" :class="{ active: isActive }">
    <div class="offcanvas__option">
      <div class="offcanvas__links">
        <router-link to="/login">Sign in</router-link>
        <a href="#">FAQs</a>
      </div>
      <div class="offcanvas__top__hover">
        <span>Usd <i class="arrow_carrot-down"></i></span>
        <ul>
          <li>USD</li>
          <li>EUR</li>
          <li>USD</li>
        </ul>
      </div>
    </div>
    <div class="offcanvas__nav__option">
      <a href="#" class="search-switch"
        ><img src="../assets/icon/search.png" alt=""
      /></a>
      <a href="#"><img src="../assets/icon/heart.png" alt=""/></a>
      <a href="#"
        ><img src="../assets/icon/cart.png" alt="" /> <span>0</span></a
      >
      <div class="price">$0.00</div>
    </div>
    <div id="mobile-menu-wrap">
      <div class="slicknav_menu">
        <a
          href="#"
          aria-haspopup="true"
          role="button"
          tabindex="0"
          class="slicknav_btn slicknav_collapsed"
          style="outline: none;"
          ><span class="slicknav_menutxt">MENU</span
          ><span class="slicknav_icon"
            ><span class="slicknav_icon-bar"></span
            ><span class="slicknav_icon-bar"></span
            ><span class="slicknav_icon-bar"></span></span
        ></a>
        <nav
          class="slicknav_nav slicknav_hidden"
          aria-hidden="true"
          role="menu"
          style="display: none;"
        >
          <ul>
            <li><a href="./index.html" role="menuitem">Home</a></li>
            <li class="active">
              <a href="./shop.html" role="menuitem">Shop</a>
            </li>
            <li><a href="./blog.html" role="menuitem">Blog</a></li>
            <li><a href="./contact.html" role="menuitem">Contacts</a></li>
          </ul>
        </nav>
      </div>
    </div>

    <div class="offcanvas__text">
      <p>Free shipping, 30-day return or refund guarantee.</p>
    </div>
  </div>
  <header class="header">
    <div class="header__top">
      <div class="container">
        <div class="row">
          <div class="col-lg-6 col-md-7">
            <div class="header__top__left">
              <p>Free shipping, 30-day return or refund guarantee.</p>
            </div>
          </div>
          <div class="col-lg-6 col-md-5">
            <div class="header__top__right">
              <div class="header__top__links">
                <router-link to="/login">Sign in</router-link>

                <a href="#">FAQs</a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="container">
      <div class="row">
        <div class="col-lg-3 col-md-3">
          <div class="header__logo">
            <router-link to="/" style="color:black"
              ><i class="bi bi-music-note"></i
              ><strong>Melody Shop</strong></router-link
            >
          </div>
        </div>
        <div class="col-lg-6 col-md-6">
          <nav class="header__menu mobile-menu">
            <ul>
              <router-link
                to="/"
                v-slot="{ navigate, href, isActive, isExactActive }"
                custom
              >
                <li
                  :class="[
                    isActive && 'active',
                    isExactActive && 'router-link-exact-active',
                  ]"
                >
                  <a :href="href" @click="navigate">Home</a>
                </li>
              </router-link>
              <router-link
                to="/shop"
                v-slot="{ navigate, href, isActive, isExactActive }"
                custom
              >
                <li
                  :class="[
                    isActive && 'active',
                    isExactActive && 'router-link-exact-active',
                  ]"
                >
                  <a :href="href" @click="navigate">Shop</a>
                </li>
              </router-link>
              <router-link
                to="/about"
                v-slot="{ navigate, href, isActive, isExactActive }"
                custom
              >
                <li
                  :class="[
                    isActive && 'active',
                    isExactActive && 'router-link-exact-active',
                  ]"
                >
                  <a :href="href" @click="navigate">About</a>
                </li>
              </router-link>
            </ul>
          </nav>
        </div>
        <div class="col-lg-3 col-md-3">
          <div class="header__nav__option">
            <a href="#" class="search-switch"
              ><img src="../assets/icon/search.png" alt=""
            /></a>
            <a href="#"><img src="../assets/icon/heart.png" alt=""/></a>
            <router-link to="/shopping-cart"
              ><img src="../assets/icon/cart.png" alt="" /> <span>{{totalItems}}</span></router-link
            >
         
          </div>
        </div>
      </div>
      <div class="canvas__open" @click="toogleCanvas">
        <i class="bi bi-justify"></i>
      </div>
    </div>
  </header>
</template>

<script>
import { mapState, mapGetters } from "vuex";

export default {
  name: "HeaderNav",
  data() {
    return {
      isActive: false,
    };
  },
  methods: {
    toogleCanvas() {
      this.isActive = !this.isActive;
    },
  },
  computed: {
    ...mapState("cart", ["products", "isShowCartDropdown", "addToCartResult"]),
    ...mapGetters("cart", ["totalItems", "subTotal"]),
  },

  created() {
    this.$store.dispatch("cart/getProductsInCart");
  },
};
</script>
