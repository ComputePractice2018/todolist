<template>
  <div class="hello">
    <h1>{{ title }}</h1>
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
      <button v-on:click="make_new_task" v-if="edit_index > -1">Закончить редактирование</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'Dolist',
  props: {
    title: String
  },
  data: function () {
    return {
      task_list: [
        {
          'id': '0',
          'name': 'Задача №1',
          'success': true
        },
        {
          'id': '1',
          'name': 'Задача №2',
          'success': false
        }
      ],
      new_task:
        {
          'id': '',
          'name': '',
          'success': ''
        },
      edit_index: -1
    }
  },
  methods: {
    add_task: function () {
      this.task_list.push(this.new_task)
    },
    remove_task: function (item) {
      this.task_list.splice(this.task_list.indexOf(item), 1)
    },
    edit_task: function (item) {
      this.edit_index = this.task_list.indexOf(item)
      this.new_task = this.task_list[this.edit_index]
    },
    make_new_task: function () {
      this.edit_index = -1
      this.new_task = {
        'id': '',
        'name': '',
        'success': ''
      }
    }
  }
}
</script>
<style>
</style>
