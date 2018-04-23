import Vue from 'vue';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.css';
import App from './App.vue';

Vue.config.productionTip = false;
Vue.use(Vuetify, {
  theme: {
    primary: '#1EB082',
  },
});

new Vue({
  render: h => h(App),
}).$mount('#app');
