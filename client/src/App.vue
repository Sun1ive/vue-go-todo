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
          <v-card class="px-5">
            <v-form @submit.prevent="addTodo">
              <v-text-field
                v-model.lazy.trim="todo.title"
                label="Todo Title"
                prepend-icon="bookmark"
              />
              <v-text-field
                v-model.lazy.trim="todo.category"
                label="Todo Category"
                prepend-icon="list"
                required
              />
              <v-select
                :items="['true', 'false']"
                v-model="todo.isDone"
                prepend-icon="done"
                label="Completed todo ?"
                required
              />
              <v-text-field
                v-model.lazy.trim="query"
                prepend-icon="search"
                label="Поиск"
                required
              />
              <v-btn
                :disabled="todo.length < 3"
                light
                color="primary"
                class="white--text"
                type="submit"
              >Add Todo</v-btn>
            </v-form>
          </v-card>
        </v-flex>
      </v-layout>
      <transition-group
        name="list"
        tag="div"
        class="justify-center myLayout"
      >
        <v-layout
          v-for="(item, i) in paginated"
          :key="i"
          justify-center
          align-center
          class="my-3"
        >
          <v-flex xs10 sm6 lg4>
            <v-card>
              <v-card-text>
                <h2 v-if="item.isDone === 'false'">{{ i + 1 }}. Title: {{ item.title }}</h2>
                <h2 v-else>{{ i + 1 }}. Title: <s>{{ item.title }}</s></h2>
                <h3><b>Completed:</b> {{ item.isDone }}</h3>
                <h3><b>Category:</b> {{ item.category }}</h3>
                <v-btn
                  flat
                  fab
                  color="red"
                  float
                  @click="removeTodo(i, item.ID)"
                >
                  <v-icon>close</v-icon>
                </v-btn>
              </v-card-text>
            </v-card>
          </v-flex>
        </v-layout>
      </transition-group>
      <v-layout justify-center align-center>
        <v-pagination :length="pLength" v-model="page" />
      </v-layout>
    </v-container>
  </v-app>
</template>

<script lang="ts">
/* eslint-disable import/no-unresolved */
import Vue from 'vue';
import api from './api/instance';

interface Todos {
  title: string;
  category: string;
  isDone: string;
}

export default Vue.extend({
  data: () => ({
    todo: {
      title: '',
      category: '',
      isDone: '',
    } as Todos,
    todos: [] as Todos[],
    page: 1 as number,
    perPage: 5 as number,
    query: '' as string,
  }),
  computed: {
    pLength(): number {
      return Math.ceil(this.todos.length / this.perPage);
    },
    paginated(): Todos[] {
      if (this.query.length > 0) {
        return this.todos.filter(todo => todo.title.toLowerCase().includes(this.query.toLowerCase()));
      }
      return this.todos.slice().splice((this.perPage * this.page) - this.perPage, this.perPage);
    },
  },
  created() {
    this.fetch();
  },
  methods: {
    async addTodo() {
      try {
        // add new todo in db
        const { data } = await api().post('/create', {
          title: this.todo.title,
          category: this.todo.category,
          isDone: this.todo.isDone,
        });
        this.todos.push(this.todo);
        this.todo = {
          title: '',
          category: '',
          isDone: '',
        };
      } catch (error) {
        throw error;
      }
    },
    async removeTodo(index: number, id: number) {
      try {
        await api().delete('/delete', {
          data: {
            id,
          },
        });
        this.todos.splice(index, 1);
      } catch (error) {
        throw error;
      }
    },
    async fetch() {
      const { data } = await api().get('/todos');
      console.log('---', data);
      this.todos = data;
    },
  },
});
</script>


<style lang="stylus">
#app
  font-family Roboto

.card__text
  position relative
  .btn
    position absolute
    top 40%
    right 10px
    transform translateY(-50%)

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

.myLayout
  min-height 650px
</style>
