<template>
  <div v-if="$data.type != null" class="alerts">
    <div v-bind:class="['alert', $data.type]" role="alert">
      {{ $data.message }}
    </div>
  </div>
</template>

<script>
import {alertStore} from '../store'

export default {
  data() {
    return alertStore.getState()
  }
}
</script>

<style scoped>
.alerts {
  position: absolute;
  width: 100%;
}
.alerts .alert {
  margin: 0 auto;
  width: 50%;
  min-width: 200px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
</style>
