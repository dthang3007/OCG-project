<template>
  <div id="preloder" v-if="isLoading == true">
    <div class="loader"></div>
  </div>
  <section class="shop spad">
    <div class="container">
      <div class="row">
        <div class="col-lg-3">
          <div class="shop__sidebar">
            <div class="shop__sidebar__search">
              <div class="input-group mb-3">
                <input
                  type="text"
                  placeholder="Search..."
                  v-model="searchKey"
                   @keypress.enter="getProductBySearch"
                   class="form-control" 
                />
                <button type="submit" @click="getProductBySearch"><i class="bi bi-search"></i></button>
              </div>
            </div>
            <div class="shop__sidebar__accordion">
              <div class="accordion" id="accordionExample">
                <div class="card">
                  <div class="card-heading">
                    <a data-toggle="collapse" data-target="#collapseOne"
                      >Categories</a
                    >
                  </div>
                  <div
                    id="collapseOne"
                    class="collapse show"
                    data-parent="#accordionExample"
                  >
                    <div class="card-body">
                      <div class="shop__sidebar__categories">
                        <ul class="nice-scroll">
                          <li>
                            <a href="#" @click="getProductByCatergory(1)"
                              >Loa
                            </a>
                          </li>
                          <li>
                            <a href="#" @click="getProductByCatergory(2)"
                              >HeadPhone</a
                            >
                          </li>
                          <li>
                            <a href="#" @click="getProductByCatergory(3)"
                              >Mp4
                            </a>
                          </li>
                          <li>
                            <a href="#" @click="getProductByCatergory(4)"
                              >Guitar</a
                            >
                          </li>
                        </ul>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="col-lg-9">
          <div class="shop__product__option">
            <div class="row">
              <div class="col-lg-6 col-md-6 col-sm-6">
                <div class="shop__product__option__left">
                  <p>
                    Showing 1–{{ products.length }} of {{ totalItems }} results
                  </p>
                </div>
              </div>
              <div class="col-lg-6 col-md-6 col-sm-6">
                <div class="shop__product__option__right">
                  <p>Sort by Price:</p>
                  <select @change="getProductBySort($event.target.value)">
                    <option value="">default</option>
                    <option value="asc">Low to High</option>
                    <option value="desc">High to Low</option>
                  </select>
                </div>
              </div>
            </div>
          </div>
          <div class="row">
            <div
              class="col-lg-4 col-md-6 col-sm-6 "
              v-for="(product, index) in products"
              :key="index"
            >
              <ProductItem :data="product"></ProductItem>
            </div>
          </div>
          <div class="row">
            <div class="col-lg-12">
              <Pagination
                :length="totalItems"
                :pageSize="limit"
                :pageIndex="pageIndex"
                @change="changePage"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
<script>
import { mapState, mapGetters } from "vuex";
import ProductItem from "../components/ProductItem.vue";
import Pagination from "../components/Pagination.vue";
export default {
  name: "Shop",
  components: {
    ProductItem,
    Pagination,
  },
  data() {
    return {
      searchKey: "",
    };
  },
  computed: {
    ...mapState("products", [
      "isLoading",
      "products",
      "totalItems",
      "pageIndex",
      "limit",
    ]),
    ...mapGetters("products", [
      "sortDropdownValue",
      "itemStartIndex",
      "itemEndIndex",
    ]),
  },

  methods: {
    // async switchPage(index) {
    //   this.isLoading = true;
    //   window.scrollTo(0, 0);
    //   let response = await axios.get(`/product?q=loa&category&sort&order`);
    //   (this.products = response.data.data), (this.meta = response.data.meta);
    //   this.isLoading = false;
    // },
    changePage(pageIndex) {
      this.$store.dispatch("products/getProducts", { pageIndex });
    },

    getProductBySort(value) {
      console.log(value);
      this.$store.dispatch("products/getProducts", {
        order: "price",
        sort: value,
      });
    },
    getProductByCatergory(value) {
      this.$store.dispatch("products/getProducts", { category: value });
    },
    getProductBySearch() {
      this.$store.dispatch("products/getProducts", { search: this.searchKey });
    },
  },
  created() {
    this.$store.dispatch("products/getProducts", {});
  },
};
</script>
