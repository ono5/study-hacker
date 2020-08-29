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

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { Languages } from './models/items'

@Component
export default class Item extends Vue {

  // typeにLanguagesを割り当てることはできない
  // https://meetup-jp.toast.com/1843
  @Prop({required: true})
  datas!: Languages[]

  // created() {
  // }
}
</script>

<style lang="scss" scoped>
  $deleteBtnColor: tomato;

  .delete {
    background-color: $deleteBtnColor;
    outline: none;
    color: #fff;
    border: 1px solid #000;
  }

  .update {
    background-color: green;
    color: #fff;
    border: 1px solid #000;
  }
</style>
