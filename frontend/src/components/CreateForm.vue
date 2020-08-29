<template>
  <div>
    <h3>Learning English</h3>
    <label for="japanese">Japanese</label>
    <textarea type="text" v-model="japanese" id="japanese" />
    <label for="english">Enlish</label>
    <textarea type="text" v-model="english" id="english" />
    <br><br>
    <button @click="post">Create</button>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'

import axios from 'axios'

@Component
export default class CreateForm extends Vue {
  japanese: string
  english: string

  constructor() {
    super()
    this.japanese = ''
    this.english = ''
  }

  post(): void {
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


</script>

<style scoped lang="sass">

</style>
