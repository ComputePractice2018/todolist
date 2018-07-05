<template>
  <div class="home">
    <Dolist title="Список задач"/>
  </div>
</template>

<script>
// @ is an alias to /src
import Dolist from '@/components/Dolist.vue'

export default {
  name: 'home',
  components: {
    Dolist
  }
}
</script>
