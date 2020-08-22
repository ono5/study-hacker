<template>
  <table class="langage-table">
    <thead>
      <tr>
        <th>No.</th>
        <th>日本語</th>
        <th>English</th>
      </tr>
    </thead>
      <Item :datas="items" />
  </table>
</template>

<script>
import axios from 'axios'
import Item from './Item'

export default {
  components: {
    Item
  },
  data() {
    return {
      items: []
    }
  },
  created: function() {
    axios.get('http://localhost:8080')
    .then(response => {
      this.items = response.data
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
