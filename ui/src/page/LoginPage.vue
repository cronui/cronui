<template>
  <div class="container">
    <div class="row justify-content-center">
      <div class="col-xl-10 col-lg-12 col-md-9">
        <div class="card border-0 shadow-lg my-5">
          <div class="card-body p-0">
            <div class="row">
              <div class="col-lg-6 d-none d-lg-block bg-image"></div>
              <div class="col-lg-6">
                <div class="p-5">
                  <div class="text-center">
                    <h1 class="h4 mb-4">Please sign in</h1>
                  </div>
                  <form class="form-install m-auto text-left" @submit.prevent="handleSubmit">
                    <div class="form-group">
                      <label for="inputUsername" class="sr-only">Username</label>
                      <input class="form-control" v-model="values.username" type="text" id="inputUsername"
                             placeholder="Your username" autofocus v-bind:class="[$data.errors.username ? 'is-invalid' : null]">
                      <div class="invalid-feedback">{{ errors.username }}</div>
                    </div>
                    <div class="form-group">
                      <label for="inputPassword" class="sr-only">Password</label>
                      <input class="form-control" v-model="values.password" type="password" id="inputPassword"
                             placeholder="Password" v-bind:class="[$data.errors.password ? 'is-invalid' : null]">
                      <div class="invalid-feedback">{{ errors.password }}</div>
                    </div>
                    <button class="btn btn-lg btn-primary btn-block mt-2" type="submit">Submit</button>
                  </form>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {authService} from '../service'
import {alertStore} from '../store'

export default {
  name: 'LoginPage',
  data() {
    return {
      values: {
        username: '',
        password: ''
      },
      errors: {},
    }
  },
  methods: {
    handleSubmit(ev) {
      const {username, password} = this.values;
      authService.login(username, password)
          .then(resp => {
            this.errors = {};
            this.$router.push('/');
          })
          .catch(e => {
            if (e.field) {
              this.errors[e.field] = e.message;
            } else {
              alertStore.error("Auth error: " + e)
            }
          })
    }
  }
}
</script>


<style scoped>
.bg-image {
  background: url("https://source.unsplash.com/505eectW54k/600x800") center;
  background-size: cover;
  border-top-left-radius: 0.25rem;
  border-bottom-left-radius: 0.25rem;
}

.form-install .form-control, .form-install .btn {
  border-radius: 1.5rem;
  padding: 1rem;
  font-size: 0.8em;
  height: 3.2rem;
}

.form-install .invalid-feedback {
  padding-left: 1.1rem;
}
</style>
