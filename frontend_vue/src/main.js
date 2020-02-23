// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Axios from 'axios'
import { BootstrapVue, BootstrapVueIcons } from 'bootstrap-vue/dist/bootstrap-vue.esm'
import App from './App'
import router from './router'
import store from './store'

Vue.config.productionTip = false
Vue.prototype.$http = Axios

Vue.use(BootstrapVue)
Vue.use(BootstrapVueIcons)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: { App }
})
