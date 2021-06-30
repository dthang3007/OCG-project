<template>
  <div class="text-center">
    <form class="form-register" @submit="login">
      <img
        class="mb-4"
        src="https://cdn.pixabay.com/photo/2017/01/18/16/15/musical-note-1990141_960_720.jpg"
        alt=""
        width="72"
        height="72"
      />
      <h1 class="h3 mb-3 font-weight-normal">Register</h1>
      <label for="inputFirstName" class="sr-only">First Name</label>
      <input
        type="text"
        id="inputFirstName"
        class="form-control"
        placeholder="First name"
        required=""
        autofocus=""
        v-model="user.first_name"
      />
      <label for="inputLastName" class="sr-only">Last Name</label>
      <input
        type="text"
        id="inputLastName"
        class="form-control"
        placeholder="Last name"
        required=""
        autofocus=""
        v-model="user.last_name"
      />

      <label for="inputEmail" class="sr-only">Email address</label>
      <input
        type="email"
        id="inputEmail"
        class="form-control"
        placeholder="Email address"
        required=""
        autofocus=""
                v-model="user.email"

      />
      <label for="inputPassword" class="sr-only">Password</label>
      <input
        type="password"
        id="inputPassword"
        class="form-control"
        placeholder="Password"
        required=""
        v-model="user.password"
      />
      <label for="inputConfirmPassword" class="sr-only">Confirm Password</label>
      <input
        type="password"
        id="inputConfirmPassword"
        class="form-control"
        placeholder="Confirm Password"
        required=""
        v-model="user.confirmPass"
      />

      <button class="btn btn-lg btn-primary btn-block mt-4" type="submit">
        Sign in
      </button>
      <div class="d-flex justify-content-between mt-3">
        <label>Do you have an account?</label>

        <router-link to="/login">Sign in</router-link>
      </div>
      <p class="mt-5 mb-3 text-muted">©2021-2022</p>
    </form>
  </div>
</template>
<script>
import {mapState} from "vuex"
export default {
  
  name: "Login",
  data() {
    return {
      user: {
        first_name: "",
        last_name: "",
        email: "",
        password: "",
        password_confirm: "",
      },
    };
  },
   computed: mapState("users", ["isLoginSuccess", "loginMessage"]),

  created() {
    if (this.isLoginSuccess) {
      this.$router.replace("/user");
    }
  },

  methods: {
    async login(user) {
      this.isLoading = true;

      await this.$store.dispatch("users/login", {
        email: user.email,
        password: user.password,
      });

      this.isLoading = false;

      if (this.isLoginSuccess) {
        this.$router.push("/user");
      }
    },
  },

};
</script>
<style>
.form-register {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: 0 auto;
}
.form-register a:hover {
  color: red;
}
</style>
