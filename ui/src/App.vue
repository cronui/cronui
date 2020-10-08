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
@import 'https://stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css';
@import '~bootstrap';
@import '~bootstrap-vue';

html, body, #app {
  height: 100%;
}
</style>
