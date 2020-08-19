<template>
  <table>
    <thead>
      <tr>
        <th>No.</th>
        <th>日本語</th>
        <th>English</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="data in datas" :key="data.id">
        <td>{{data.id}}</td>
        <td>{{data.japanese}}</td>
        <td>{{data.english}}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  data() {
    return {
      datas: [{id: 1, japanese: 'テスト', english: 'test'}]
    }
  }
}
</script>
