<template>
  <div class="text-center">
    <form class="form-signin">
      <img
        class="mb-4"
        src="https://cdn.pixabay.com/photo/2017/01/18/16/15/musical-note-1990141_960_720.jpg"
        alt=""
        width="72"
        height="72"
      />
      <h1 class="h3 mb-3 font-weight-normal">Please sign in</h1>
      <label for="inputEmail" class="sr-only">Email address</label>
      <input
        type="text"
        id="inputEmail"
        class="form-control"
        placeholder="Email address"
        required=""
        autofocus=""
      />
      <label for="inputPassword" class="sr-only">Password</label>
      <input
        type="password"
        id="inputPassword"
        class="form-control"
        placeholder="Password"
        required=""
      />
      <button class="btn btn-lg btn-primary btn-block mt-4" type="submit">
        Sign in
      </button>
      <div class="d-flex justify-content-between">
        <label>Do you have an account?</label>
        
        <router-link to="/register">Register</router-link>
      </div>

      <p class="mt-5 mb-3 text-muted">© 2020-2021</p>
    </form>
  </div>
</template>
<script>
export default {
  name: "Login",
 
};
</script>
<style>
.form-signin {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: 0 auto;
}
.form-signin a:hover{
color:red
}
</style>
