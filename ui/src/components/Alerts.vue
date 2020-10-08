<template>
  <div v-if="$data.length > 0" class="alerts">
    <div v-for="alert in $data" class="alert" v-bind:class="[alert.type]" role="alert">
      {{ alert.message }}
      <button type="button" class="close" data-dismiss="alert" aria-label="Close" @click="close(alert)">
        <span aria-hidden="true">&times;</span>
      </button>
    </div>
  </div>
</template>

<script>
import {alertStore} from '../store'

export default {
  data() {
    console.log(alertStore.getState());
    return alertStore.getState();
  },
  methods: {
    close(alert) {
      alertStore.remove(alert);
    }
  }
}
</script>

<style scoped>
.alerts {
  position: absolute;
  width: 100%;
  z-index: 100;
}
.alerts .alert {
  margin: 0 auto 1rem;
  width: 50%;
  min-width: 200px;
}
.alerts .alert:first-child {
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
</style>
