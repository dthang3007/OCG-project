<template>
  <!-- <BreadCrumb />
  <Shop /> -->

  <HeaderNav />
  <router-view />
  <FooterVue />
</template>

<script>
import HeaderNav from "./components/HeaderNav";
// import BreadCrumb from "./components/BreadCrumb.vue";
import FooterVue from "./components/Footer.vue";
// import Shop from "./components/Shop.vue";
export default {
  name: "App",
  components: {
    HeaderNav,

    FooterVue,
  },
};
</script>
<style lang="scss">
@import "./sass/style.scss";
</style>
