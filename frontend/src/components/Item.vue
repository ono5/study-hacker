<template>
  <tbody>
    <tr v-for="(data, index) in datas" :key="data.index">
      <td>{{index + 1}}</td>
      <td>{{data.japanese}}</td>
      <td>{{data.english}}</td>
      <td><button class="update">更新</button></td>
      <td><button class="delete">削除</button></td>
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

<style lang="scss" scoped>
  $deleteBtnColor: tomato;

  .delete {
    background-color: $deleteBtnColor;
  }
</style>
