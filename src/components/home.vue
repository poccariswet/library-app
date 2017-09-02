<template>
  <div id="home">
    <h1>{{ title }}</h1>
    <div class="button">
      <el-button type="primary" @click="push_AllBooks">全ての本の情報へ</el-button>
      <div class="marginTop"></div>
      <el-button type="primary" @click="push_PostInfo">本の情報を追加へ</el-button>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      title: 'Library-API'
    }
  },
  methods: {
    push_AllBooks () {
      this.$router.push({ name: 'allbooks' })
    },
    push_PostInfo () {
      this.$router.push({ name: 'postinfo' })
    }
  }
}
</script>

<style scoped>

#home {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

#home button {
  text-align: center;
  height: 100px;
  line-height: 50px;
  font-size: 2em;
  border: 1px solid #ccc;
  border-radius: 40px;
  margin: 5px;
  cursor: pointer;
}
.marginTop {
  margin-top: 1%;
}
</style>
