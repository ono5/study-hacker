<template>
  <table class="langage-table">
    <thead>
      <tr>
        <th>No.</th>
        <th>日本語</th>
        <th>English</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(data, index) in datas" :key="data.index">
        <td>{{index + 1}}</td>
        <td>{{data.japanese}}</td>
        <td>{{data.english}}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import axios from 'axios'
export default {
  data() {
    return {
      datas: []
    }
  },
  created: function() {
    axios.get('http://localhost:8080')
    .then(response => {
      console.log(response)
      this.datas = response.data
    })
  },
}
</script>

<style scoped>
.langage-table {
  margin-left: auto;
  margin-right: auto;
}
</style>
