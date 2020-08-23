<template>
  <div>
    <h3>Learning English</h3>
    <label for="japanese">Japanese</label>
    <textarea type="text" v-model="japanese" id="japanese" />
    <label for="english">Enlish</label>
    <textarea type="text" v-model="english" id="english" />
    <br><br>
    <button @click="create">Create</button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      japanese: "",
      english: ""
    }
  },
  methods: {
    create() {
      axios.post("http://localhost:8080/create", {
        "japanese": this.japanese,
        "english": this.english
      })
      .then(response => {
        console.log({response})
        window.location.reload()
      })
      .catch(err => console.log({err}))
      this.japanese = ''
      this.english = ''
    }
  }
}
</script>

<style scoped lang="sass">

</style>
