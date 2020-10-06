<template>
  <Alerts />
  <router-view />
</template>

<script>
import Alerts from './components/Alerts.vue';

export default {
  components: {
    Alerts
  }
}
</script>

<style>
@import '~bootstrap';
@import '~bootstrap-vue';

html, body, #app {
  height: 100%;
}
</style>
