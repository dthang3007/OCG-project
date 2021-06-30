<template>
  <!-- Hero Section Begin -->

  <!-- Hero Section End -->

  <!-- Banner Section Begin -->
    <div id="preloder" v-if="isLoading==true">
        <div class="loader"  ></div>
    </div>
  <section class="banner spad">
    <div class="container">
      <div class="row">
        <div class="col-lg-7 offset-lg-4">
          <div class="banner__item">
            <div class="banner__item__pic">
              <img
                src="https://www.sony.com.vn/image/fec35cd990f766e0d1e02531d9210cc5?fmt=pjpeg&wid=660&bgcolor=F1F5F9&bgc=F1F5F9"
                style="width:440px;height:440px"
                alt=""
              />
            </div>
            <div class="banner__item__text">
              <h2>HeadPhone</h2>
              <router-link href="#" to="/shop">Shop now</router-link>
            </div>
          </div>
        </div>
        <div class="col-lg-5">
          <div class="banner__item banner__item--middle">
            <div class="banner__item__pic">
              <img
                src="https://www.marshallheadphones.com/on/demandware.static/-/Sites-zs-master-catalog/default/dwaa4d3aca/images/marshall/speakers/acton-ii-bt/black/high-res/pos-marshall-acton-ii-bt-black-01.png"
                style="width:440px;height:440px"
                alt=""
              />
            </div>
            <div class="banner__item__text">
              <h2>Speaker</h2>
              <a href="#">Shop now</a>
            </div>
          </div>
        </div>
        <div class="col-lg-7">
          <div class="banner__item banner__item--last">
            <div class="banner__item__pic">
              <img
                src="https://anhducdigital.vn/media/product/1797_1_3.jpg"
                style="width:440px;height:440px"
                alt=""
              />
            </div>
            <div class="banner__item__text">
              <h2>Mp4 player case</h2>
              <a href="#">Shop now</a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
  <!-- Banner Section End -->

  <!-- Product Section Begin -->
  <section class="product spad">
    <div class="container">
      <div class="row">
        <div class="col-lg-12">
          <ul class="filter__controls">
            <li class="active" data-filter="*">Best Sellers</li>
            <li data-filter=".new-arrivals">New Arrivals</li>
            <li data-filter=".hot-sales">Hot Sales</li>
          </ul>
        </div>
      </div>
      <div class="row product__filter">
        <div
          v-for="(product, index) in products"
          :key="index"
          class="col-lg-3 col-md-6 col-sm-6 col-md-6 col-sm-6 mix new-arrivals"
        >
          <ProductItem :data="product"></ProductItem>
        </div>
      </div>
    </div>
  </section>
  <!-- Product Section End -->

  <!-- Categories Section Begin -->
  <!-- <section class="categories spad">
    <div class="container">
      <div class="row">
        <div class="col-lg-3">
          <div class="categories__text">
            <h2>
              Production Hot <br />
              <span>Speaker Collection</span> <br />
              Accessories
            </h2>
          </div>
        </div>
        <div class="col-lg-4">
          <div class="categories__hot__deal">
            <img src="../assets/product-sale.jpg" alt="" />
            <div class="hot__deal__sticker">
              <span>Sale Of</span>
              <h5>$29.99</h5>
            </div>
          </div>
        </div>
        <div class="col-lg-4 offset-lg-1">
          <div class="categories__deal__countdown">
            <span>Deal Of The Week</span>
            <h2>Multi-music production</h2>
            <div class="categories__deal__countdown__timer" id="countdown">
              <div class="cd-item">
                <span>{{time.days}}</span>
                <p>Days</p>
              </div>
              <div class="cd-item">
                <span>{{time.hours}}</span>
                <p>Hours</p>
              </div>
              <div class="cd-item">
                <span>{{time.minutes}}</span>
                <p>Minutes</p>
              </div>
              <div class="cd-item">
                <span>{{time.seconds}}</span>
                <p>Seconds</p>
              </div>
            </div>
            <a href="#" class="primary-btn">Shop now</a>
          </div>
        </div>
      </div>
    </div>
  </section> -->
  <!-- Categories Section End -->

  <!-- Instagram Section Begin -->
  <section class="instagram spad">
    <div class="container">
      <div class="row">
        <div class="col-lg-8">
          <img
            src="../assets/instagram/instagram-1.jpg"
            alt="..."
            class="instagram__pic__item"
          />
          <img
            src="../assets/instagram/instagram-2.jpg"
            alt="..."
            class="instagram__pic__item"
          />
          <img
            src="../assets/instagram/instagram-3.jpg"
            alt="..."
            class="instagram__pic__item"
          />
          <img
            src="../assets/instagram/instagram-4.jpg"
            alt="..."
            class="instagram__pic__item"
          />
          <img
            src="../assets/instagram/instagram-5.jpg"
            alt="..."
            class="instagram__pic__item"
          />
          <img
            src="../assets/instagram/instagram-6.jpg"
            alt="..."
            class="instagram__pic__item"
          />
        </div>
        <div class="col-lg-4">
          <div class="instagram__text">
            <h2>Instagram</h2>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
              eiusmod tempor incididunt ut labore et dolore magna aliqua.
            </p>
            <h3>#Melody_Shop</h3>
          </div>
        </div>
      </div>
    </div>
  </section>
  <!-- Instagram Section End -->

  <!-- Latest Blog Section Begin -->
  
</template>
<script>
import ProductItem from "../components/ProductItem.vue";
import axios from "../axios";
export default {
  components: { ProductItem },
  data() {
    return {
      products: [],
      time:{days:0,hours:0,minutes:0,seconds:0},
      isLoading:true,
    };
  },
  name: "Home",
  methods: {
  
  },
  async created() {
    let respon =  await axios.get("/product");
    this.products = respon.data.data.slice(0,8);
    this.isLoading=false
  },

};
</script>
<style scoped>
.instagram__pic__item {
  width: 33.33% !important;
  float: left !important;
  height: 261px !important;
}
</style>
