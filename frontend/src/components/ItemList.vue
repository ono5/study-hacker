<template>
  <table class="langage-table">
    <thead>
      <tr>
        <th>No.</th>
        <th>日本語</th>
        <th>English</th>
        <th></th>
        <th></th>
      </tr>
    </thead>
      <Item :datas="items" />
  </table>
</template>

<script lang="ts">
import axios from 'axios'
import Item from './Item.vue'

import { Component, Vue } from 'vue-property-decorator'
import { Languages } from './models/items'

@Component({
  components: {
    Item
  }
})
export default class ItemList extends Vue{
  items: Languages[] = []

  created() {
    axios.get('http://localhost:8080')
    .then(response => {
      // console.log("Test")
      this.items = response.data
    })
  }
}
</script>

<style scoped>
.langage-table {
  margin-left: auto;
  margin-right: auto;
}
</style>
