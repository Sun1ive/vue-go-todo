import Vue from 'vue';
import App from './App.vue';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.css';

Vue.config.productionTip = false;
Vue.use(Vuetify, {
  theme: {
    primary: '#ff995e', // #E53935
  },
});

new Vue({
  render: h => h(App),
}).$mount('#app');
