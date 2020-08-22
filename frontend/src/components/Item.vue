<template>
  <tbody>
    <tr v-for="(data, index) in datas" :key="data.index">
      <td>{{index + 1}}</td>
      <td>{{data.japanese}}</td>
      <td>{{data.english}}</td>
    </tr>
  </tbody>
</template>

<script>
export default {
  props: ['datas'],
  created: function() {
    console.log(this.datas)
  }
}
</script>
