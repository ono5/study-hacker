<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <CreateForm />
    <br>
    <br>
    <item-list />
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import CreateForm from '@/components/CreateForm'
import ItemList from '@/components/ItemList'

export default {
  name: 'Home',
  components: {
    // HelloWorld
    CreateForm,
    ItemList
  }
}
</script>
