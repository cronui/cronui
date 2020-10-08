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
                      <input class="form-control" v-model="values.username" type="text" id="inputUsername"
                             autocomplete="new-password" v-bind:class="[$data.errors.username ? 'is-invalid' : null]"
                             placeholder="Your username" @change="validate('username')">
                      <div class="invalid-feedback">{{ errors.username }}</div>
                    </div>
                    <div class="form-group">
                      <label for="inputPassword" class="sr-only">Password</label>
                      <input class="form-control" v-model="values.password" type="password" id="inputPassword"
                             autocomplete="new-password" v-bind:class="[$data.errors.password ? 'is-invalid' : null]"
                             placeholder="Password" @change="validate('password')">
                      <div class="invalid-feedback">{{ errors.password }}</div>
                    </div>
                    <div class="form-group">
                      <label for="inputPasswordRepeat" class="sr-only">PasswordRepeat</label>
                      <input class="form-control" v-model="values.passwordRepeat" type="password"
                             id="inputPasswordRepeat"
                             autocomplete="new-password"
                             v-bind:class="[$data.errors.passwordRepeat ? 'is-invalid' : null]"
                             placeholder="Password (repeat)" @change="validate('passwordRepeat')">
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
import * as yup from 'yup';
import {installService} from '../service'
import {alertStore} from '../store'

const installFormSchema = yup.object().shape({
  username: yup.string().required().min(4).max(36),
  password: yup.string().required().min(6),
  passwordRepeat: yup.string().oneOf([yup.ref('password')], 'Passwords must match'),
})

export default {
  name: 'InstallPage',
  data() {
    return {
      values: {
        username: '',
        password: '',
        passwordRepeat: ''
      },
      errors: {},
    }
  },
  methods: {
    validate(field) {
      installFormSchema.validateAt(field, this.values)
          .then(() => {
            delete this.errors[field];
          })
          .catch(err => {
            this.errors[field] = err.message.capitalize();
          })
    },
    handleSubmit(e) {
      installFormSchema.validate(this.values, {abortEarly: false})
          .then(() => {
            this.errors = {};
            installService.postUser(this.values.username, this.values.password)
                .then(resp => {
                  console.log(resp);
                  alertStore.success("User with username '" + this.values.username + "' created.");
                  installService.clearStore();
                  this.$router.push('/login');
                })
                .catch((e) => alertStore.error("Install error: " + e))
          })
          .catch(err => {
            err.inner.forEach(error => {
              this.errors[error.path] = error.message.capitalize();
            });
          });
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
