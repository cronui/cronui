import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

String.prototype.capitalize = function() {
    return this.charAt(0).toUpperCase() + this.slice(1)
}

createApp(App)
    .use(router)
    .mount('#app')
