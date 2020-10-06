<template>
  <div class="container">
    <div class="row justify-content-center">
      <div class="col-xl-10 col-lg-12 col-md-9">
        <div class="card border-0 shadow-lg my-5">
          <div class="card-body p-0">
            <div class="row">
              <div class="col-lg-6">
                <div class="p-5">
                  <div class="text-center">
                    <h1 class="h4 mb-4">Install CronUI</h1>
                    <p class="mb-4 text-muted small">It seems like there is no user account available. Please, create
                      the first user using form below.</p>
                  </div>
                  <form class="form-install m-auto text-left" @submit.prevent="handleSubmit">
                    <div class="form-group">
                      <label for="inputUsername" class="sr-only">Username</label>
                      <input v-model="username" type="text" id="inputUsername"
                             v-bind:class="['form-control', $data.errors.username ? 'is-invalid' : null]"
                             placeholder="Your username" required @change="validateUsername">
                      <div class="invalid-feedback">{{ errors.username }}</div>
                    </div>
                    <div class="form-group">
                      <label for="inputPassword" class="sr-only">Password</label>
                      <input v-model="password" type="password" id="inputPassword"
                             v-bind:class="['form-control', $data.errors.password ? 'is-invalid' : null]"
                             placeholder="Password" required @change="validatePassword">
                      <div class="invalid-feedback">{{ errors.password }}</div>
                    </div>
                    <div class="form-group">
                      <label for="inputPasswordRepeat" class="sr-only">PasswordRepeat</label>
                      <input v-model="passwordRepeat" type="password" id="inputPasswordRepeat"
                             v-bind:class="['form-control', $data.errors.passwordRepeat ? 'is-invalid' : null]"
                             placeholder="Password (repeat)" required @change="validatePasswordRepeat">
                      <div class="invalid-feedback">{{ errors.passwordRepeat }}</div>
                    </div>
                    <button class="btn btn-lg btn-primary btn-block mt-2" type="submit">Submit</button>
                  </form>
                </div>
              </div>
              <div class="col-lg-6 d-none d-lg-block bg-image"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {installService} from '../service'
import {alertStore} from '../store'

export default {
  name: 'InstallPage',
  data() {
    return {
      errors: {},
      username: 'admin', // TODO remove values
      password: 'hesloo',
      passwordRepeat: 'hesloo'
    }
  },
  methods: {
    validateUsername() {
      delete this.errors.username;
      if (this.username.length < 4) {
        this.errors.username = 'Please use at least 4 characters.';
      }
    },
    validatePassword() {
      delete this.errors.password;
      if (this.password.length < 6) {
        this.errors.password = 'Please use at least 6 characters.';
      }
    },
    validatePasswordRepeat() {
      delete this.errors.passwordRepeat;
      if (this.passwordRepeat !== this.password) {
        this.errors.passwordRepeat = 'Repeat the same password.';
      }
    },
    handleSubmit(e) {
      this.validateUsername()
      this.validatePassword()
      this.validatePasswordRepeat()

      if (Object.keys(this.errors).length === 0) {
        installService.postUser(this.username, this.password)
            .then(resp => {
              console.log(resp);
              alertStore.success("User with username '" + this.username + "' created.");
              installService.clearStore();
              this.$router.push('/login');
            })
            .catch((e) => alertStore.error("Install error: " + e))
      }
    }
  }
}
</script>

<style scoped>
.bg-image {
  background: url("https://source.unsplash.com/PRCgmkh_ggc/600x800") center;
  background-size: cover;
  border-top-right-radius: 0.25rem;
  border-bottom-right-radius: 0.25rem;
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
