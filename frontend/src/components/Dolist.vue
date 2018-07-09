<template>
  <div class="hello">
    <h1>{{ title }}</h1>
    <h3 v-if="error">Ошибка: {{error}}</h3>
    <table>
      <tr>
        <th>ID</th>
        <th>Название</th>
        <th>Выполнена</th>
      </tr>
      <tr v-for="task in task_list" v-bind:key="task.id">
       <td>{{task.id}}</td>
       <td>{{task.name}}</td>
       <td>{{task.success}}</td>
       <td><button
        v-on:click="remove_task(task)">Удалить задачу</button></td>
        <td><button
        v-on:click="edit_task(task)">Редактировать задачу</button></td>
      </tr>
    </table>
     <h3 v-if="edit_index == -1">Новая задача</h3>
    <form>
      <p> ID:<input type="text" v-model="new_task.id"></p>
      <p> Название:<input type="text" v-model="new_task.name"></p>
      <p>Выполнение:<input type="checkbox" v-model="new_task.success"></p>
      <button v-if="edit_index == -1" v-on:click="add_task">Добавить</button>
      <button v-on:click="end_of_edition" v-if="edit_index > -1">Закончить редактирование</button>
    </form>
  </div>
</template>

<script>
const axios = require('axios')
export default {
  name: 'Dolist',
  props: {
    title: String
  },
  data: function () {
    return {
      task_list: [],
      new_task:
        {
          'id': '',
          'name': '',
          'success': ''
        },
      edit_index: -1,
      error: ''
    }
  },
  mounted: function () {
    this.get_tasks()
  },
  methods: {
    get_tasks: function () {
      this.error = ''
      const url = 'api/task/getList'
      axios.get(url).then(response => {
        this.task_list = response.data
      }).catch(response => {
        this.error = response.response.data
      })
    },
    add_task: function () {
      this.error = ''
      const url = 'api/task/add'
      axios.post(url, this.new_task).then(response => {
        this.task_list.push(this.new_task)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    remove_task: function (item) {
      this.error = ''
      const url = '/api/task/del/' + this.task_list.indexOf(item)
      axios.delete(url).then(response => {
        this.task_list.splice(this.task_list.indexOf(item), 1)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    edit_task: function (item) {
      this.edit_index = this.task_list.indexOf(item)
      this.new_task = this.task_list[this.edit_index]
    },
    end_of_edition: function () {
      this.error = ''
      const url = 'api/task/' + this.edit_index
      axios.put(url, this.new_task).then(response => {
        this.edit_index = -1
        this.new_task = {
          'id': '',
          'name': '',
          'success': ''
        }
      }).catch(response => {
        this.error = response.response.data
      })
    }
  }
}
</script>
<style>
</style>
