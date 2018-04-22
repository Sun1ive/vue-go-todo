<template>
  <v-app>
    <v-container fluid>
      <v-layout justify-center>
        <v-flex xs10 sm6 lg4>
          <v-card>
            <v-card-title>
              <h2>Vue.js + GO TodoList</h2>
            </v-card-title>
          </v-card>
        </v-flex>
      </v-layout>
      <v-layout justify-center>
        <v-flex xs10 sm6 lg4 class="text-xs-center">
          <v-text-field v-model.lazy.trim="todo" />
          <v-btn
            dark
            color="primary"
            @click="addTodo"
          >Add Todo</v-btn>
        </v-flex>
      </v-layout>
    </v-container>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import api from './api/instance';

export default Vue.extend({
  data: () => ({
    todo: '' as string,
  }),
  methods: {
    async addTodo() {
      api().post('/create', {});
    },
  },
});
</script>


<style lang="stylus">
#app
  font-family Roboto

.card__title
  text-align center
  justify-content center
</style>
