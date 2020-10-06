<template>
  <div class="d-flex flex-column align-content-center h-75">
    <form class="form-signin m-auto text-center" @submit.prevent="handleSubmit">
      <h1 class="h3 mb-3 font-weight-normal">Please sign in</h1>
      <label for="inputEmail" class="sr-only">Username</label>
      <input v-model="input.username" type="text" id="inputEmail" class="form-control" placeholder="Username" required autofocus>
      <label for="inputPassword" class="sr-only">Password</label>
      <input v-model="input.password" type="password" id="inputPassword" class="form-control" placeholder="Password" required>
      <button class="btn btn-lg btn-primary btn-block mt-2" type="submit">Sign in</button>
    </form>
  </div>
</template>

<script>
import {authService} from '../service'
import {alertStore} from '../store'

export default {
  name: 'LoginPage',
  data() {
    return {
      input: {
        username: '',
        password: ''
      }
    }
  },
  methods: {
    handleSubmit(e) {
      const {username, password} = this.input;
      authService.login(username, password)
          .catch((e) => alertStore.error("Auth error: " + e))
    }
  }
}
</script>


<style scoped>
.form-signin {
  width: 100%;
  max-width: 300px;
}

.form-signin .form-control {
  font-size: 1.1em;
}

.form-signin .form-control:focus {
  z-index: 2;
}

.form-signin input[type="text"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}

.form-signin input[type="password"] {
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
</style>
