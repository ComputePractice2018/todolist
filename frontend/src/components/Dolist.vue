<template>
  <b-row>
    <b-col sm="8" offset-sm="2">
      <input type="text" class="todo-input" placeholder="Добавьте задачи и нажмите Enter"
     v-model="new_task" @keyup.enter="add_task"/>
    <b-container>
      <p> {{ error }}</p>
        <table class="table table-striped table-dark">
          <transition-group name="fade" enter-active-class="animated fadeInUp" leave-active-class="animated fadeOutDown">
             <tr v-for="(item, index) in todosFiltered" :key="item.id" >
                 <td  class="w-10"><input class="mych" type="checkbox" v-model="item.success" @click="editSuccess(item)"></td>
                 <td class="w-100">
                    <div v-if="!item.editing" class="todo-item-label" :class="{ success : item.success }">{{ item.name }}</div>
                    <input v-else class="todo-item-edit" type="text" @blur="doneEdit(item)" v-model="item.name" @keyup.enter="doneEdit(item)" @keyup.esc="cancelEdit(item)" v-focus>
                  </td>
                  <td class="w-10">
                     <button @click="editTodo(item)"><ion-icon name="create"></ion-icon></button>
                  </td>
                  <td class="w-10">
                    <button class="btn-danger" @click="removeTodo(item,index)"><ion-icon name="close"></ion-icon></button>
                  </td>
              </tr>
           </transition-group>
        </table>
    </b-container>
    <hr class="myhr"/>
    <b-row>
    <b-col md="5" offset-md="1">
    </b-col>
    <b-col md="6">
        <span>Не выполнено задач: {{ remaining }}</span>
    </b-col>
</b-row>
<hr class="myhr"/>
    <b-row>
        <b-col lg="1" offset-lg="1" sm="1" offset-sm="1" cols="1" offset="1">
          <button :class="{ 'active': filter == 'all' }" @click="filter = 'all'">Все</button>
        </b-col>
        <b-col lg="2" offset-lg="3" sm="2" offset-sm="1" cols="2" offset="1">
          <button  :class="{ 'active': filter == 'active' }" @click="filter = 'active'">Активные</button>
        </b-col>
        <b-col lg="2" offset-lg="2" sm="2" offset-sm="1" cols="2" offset="1">
          <button  :class="{ 'active': filter == 'success' }" @click="filter = 'success'">Выполненные</button>
       </b-col>
    </b-row>
    </b-col>
  </b-row>
</template>

<script>
const axios = require('axios')
export default {
  name: 'Todolist',
  data () {
    return {
      new_task: '',
      beforeEditCache: '',
      filter: 'all',
      error: '',
      tasks: []
    }
  },
  computed: {
    remaining () {
      return this.tasks.filter(tasks => !tasks.success).length
    },
    anyRemaining () {
      return this.remaining !== 0
    },
    todosFiltered () {
      if (this.filter === 'all') {
        return this.tasks
      } else if (this.filter === 'active') {
        return this.tasks.filter(todo => !todo.success)
      } else if (this.filter === 'success') {
        return this.tasks.filter(todo => todo.success)
      }
      return this.todos
    },
    showClearsuccessButton () {
      return this.tasks.filter(todo => todo.success).length > 0
    }
  },
  mounted: function () {
    this.get_tasks()
  },
  directives: {
    focus: {
      inserted: function (el) {
        el.focus()
      }
    }
  },
  methods: { /*  первая загрузка*/ 
    get_tasks () {
      this.error = ''
      const url = 'api/todolist/task'
      axios.get(url).then(response => {
        this.tasks = response.data
      }).catch(response => {
        this.error = response.response.data
      })
    },
    add_task () { /* добавление */
      if (this.new_task.trim().length === 0) {
        return
      }
      this.error = ''
      const url = 'api/todolist/task'
      axios.post(url, {
        name: this.new_task
      }).then(response => {
        this.tasks.push({
          id: response.data.id,
          name: response.data.name,
          success: false,
          editing: false
        })
        this.new_task = ''
      }).catch(response => {
        this.error = response.response.data
      })
    },
    removeTodo (item, index) { /* удаление  */
      axios.delete('api/todolist/task/' + item.id)
        .then(response => {
          this.tasks.splice(index, 1)
        })
        .catch(error => {
          this.error = response.response.data
          console.log(this.error)
        })
    },
    doneEdit (item) { /*завершилось редактирование*/
      if (item.name.trim() === '') {
        item.name = this.beforeEditCache
        item.editing = false
        return /* выйти из редактирования если пустая задача */
      }
    axios.patch('api/todolist/task/' + item.id, {
        name: item.name,
        success: item.success
      })
        .then(response => {
          item.editing = false
        })
        .catch(error => {
          this.error = response.response.data
          console.log(this.error)
        })
    },
    editSuccess (item) {  /* успешна 1 задача или не успешна  роут такой же что для редактирования*/
      item.success = !item.success
      axios.patch('api/todolist/task/' + item.id, {
        name: item.name,
        success: item.success
      })
        .then(response => {})
        .catch(error => {
          this.error = response.response.data
          console.log(this.error)
        })
    },
    cancelEdit(item) { /* выход из редактирования */
      item.name = this.beforeEditCache
      item.editing = false
      this.beforeEditCache = ''
    },
    editTodo(item){ /* начало редактирования занесение кеш */
      this.beforeEditCache = item.name
      item.editing = true
    }
  }
}
</script>
<style >
@import url("https://cdnjs.cloudflare.com/ajax/libs/animate.css/3.5.2/animate.min.css");
.myhr{
   border: none;
    color: #000;
    background-color: #000;
    height: 2px;
}
  .todo-input {
    width: 100%;
    padding: 10px 18px;
    font-size: 18px;
    margin-bottom: 16px;
  }
  .todo-item {
    animation-duration: 0.3s;
  }
  .todo-item-label {
    padding: 10px;
    margin-left: 12px;
  }
  .todo-item-edit {
   padding: 10px 18px;
    font-size: 18px;
  }
  .success {
    text-decoration: line-through;
    color: grey;
  }
  button {
  font-size: 14px;
  color: #fff;
  background-color: #1d2124;
  border-color: #171a1d;
  border-radius: 5px;
  border:none;
  box-shadow: inset 0px 1px 0px #0a0d0e, 0px 5px 0px 0px #6f7274, 0px 10px 5px #fff;
  }
  .active {
    color: #fff;
    background-color: #1e7e34;
    border-color: #1c7430;
  }
  .fade-enter-active, .fade-leave-active {
    transition: opacity .2s;
  }
  .fade-enter, .fade-leave-to {
    opacity: 0;
  }
  .mydel{
  font-size: 14px;
  color: #fff;
  background-color: #dc3545;
  border-color: #dc3545;
  border-radius: 5px;
  border:none;
  }
  .mycont{
    display: flex;
    align-items: center;
    justify-content: center;
    height: 20px;
  }
  .mych{
  transform:scale(2);
  opacity:0.9;
  cursor:pointer;
  vertical-align: bottom;
  }
@media (min-width: 576px)
{
  button {
    font-size: 12px;
    }
}
</style>
