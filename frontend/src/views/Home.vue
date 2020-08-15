<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <CreateForm />
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import CreateForm from '@/components/CreateForm'

export default {
  name: 'Home',
  components: {
    // HelloWorld
    CreateForm
  }
}
</script>
