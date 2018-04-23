<template>
  <v-app>
    <v-container fluid>
      <v-layout justify-center>
        <v-flex xs10 sm6 lg4>
          <v-card-title>
            <h2>Vue.js + GO TodoList</h2>
          </v-card-title>
        </v-flex>
      </v-layout>
      <v-layout justify-center>
        <v-flex xs10 sm6 lg4 class="text-xs-center">
          <v-form @submit.prevent="addTodo">
            <v-btn
              dark
              color="primary"
              type="submit"
            >Add Todo</v-btn>
            <v-text-field
              v-model.lazy.trim="todo"
              label="Todo"
              append-icon="add"
            />
          </v-form>
        </v-flex>
      </v-layout>
      <transition-group
        name="list"
        tag="div"
        class="justify-center"
      >
        <v-layout
          v-for="(item, i) in todos"
          :key="i"
          justify-center
          align-center
          class="my-3"
        >
          <v-flex xs10 sm6 lg4>
            <v-card>
              <v-card-text>
                <b class="mr-2">{{ i + 1}}.</b><strong>{{ item }}</strong>
              </v-card-text>
            </v-card>
          </v-flex>
        </v-layout>
      </transition-group>
    </v-container>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import api from './api/instance';

export default Vue.extend({
  data: () => ({
    todo: '' as string,
    todos: ['hellow', 'rowld', 'something'] as string[],
  }),
  methods: {
    async addTodo() {
      // api().post('/create', {});
      this.todos.push(this.todo);
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

.list-item
  display: inline-block;
  margin-right: 10px;

.list-enter-active, .list-leave-active
  transition: all 1s;

.list-enter, .list-leave-to
  opacity: 0;
  transform: translateY(30px);


</style>
