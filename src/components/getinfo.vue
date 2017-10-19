<template>
  <div id="getinfo">
    <div class="detail">
      <div class="block">
        <img v-bind:src="data.Pic">
      </div>
      <h3>タイトル</h3>
      <p>{{ data.Title }}</p>
      <h3>著者/作者</h3>
      <p>{{ data.Author }}</p>
      <h3>Isbn13</h3>
      <p>{{ data.Isbn13 }}</p>
      <h3>貸し出し状況</h3>
      <!-- <p>{{ data.State }}</p> -->
      <p>{{ Info }}</p>
    </div>
    <div class="marginTop"></div>
    <el-button type="danger" @click="push_AllBooks">戻る</el-button>
    <el-button type="success" @click="dialogEdit = true">編集</el-button>
    <el-button @click="push_home()">ホームへ</el-button>

    <!-- MODAL -->
    <el-dialog
    title="EDIT"
    :visible.sync="dialogEdit"
    size="tiny">
    <div class="marginTop"></div>
    <label>Title</label>
    <el-input
    placeholder="Title"
    v-model="title">
  </el-input>
  <div class="marginTop"></div>
  <label>Author</label>
  <el-input
  placeholder="Author"
  v-model="author">
</el-input>
<div class="marginTop"></div>
<label>Isbn13</label>
<el-input
placeholder="Isbn13"
v-model="isbn13">
</el-input>
<div class="marginTop"></div>
<label>貸し出し状況</label>
<el-switch on-text="可" off-text="不可" v-model="state"></el-switch>
<div class="marginTop"></div>
<label>Pic address</label>
<el-input
placeholder="picture address"
v-model="pic">
</el-input>
<span slot="footer" class="dialog-footer">
  <el-button @click="dialogEdit = false">Cancel</el-button>
  <el-button type="success" @click="edit">Confirm</el-button>
</span>
</el-dialog>
</div>
</template>

<script>
import request from 'superagent'
import Q from 'q'
import Utils from '../request.js'
export default {
  data () {
    return {
      id: '',
      data: '',
      title: '',
      author: '',
      isbn13: '',
      state: '',
      pic: '',
      Info: '',
      dialogEdit: false
    }
  },
  mounted () {
    this.call()
  },
  methods: {
    set () {
      this.id = this.data.Id
      this.title = this.data.Title
      this.author = this.data.Author
      this.isbn13 = this.data.Isbn13
      this.state = this.data.State
      this.pic = this.data.Pic
    },
    push_AllBooks () {
      this.$router.push({ name: 'allbooks' })
    },
    push_home () {
      this.$router.push({ name: 'home' })
    },
    call () {
      Q.fcall(() => {
        return Utils.request2(`http://localhost:9090/book/api/get/${this.$route.params.title}`)
      })
      .then(d => {
        console.log(JSON.stringify(d.body))
        this.data = d.body
        this.Info = this.getstate(this.data.State)
        this.set()
      })
    },
    put (_obj) {
      request
      .put('http://localhost:9090/book/api/update')
      .set('Content-Type', 'application/json')
      .send(_obj)
      .end((err, res) => {
        if (err) console.log('err')
        if (res) console.log('success')
      })
    },
    getstate (state) {
      if (state) {
        return '貸出可'
      } else {
        return '貸出不可'
      }
    },
    edit () {
      let obj = {}
      obj.Id = this.id
      obj.Title = this.title
      obj.Author = this.author
      obj.Isbn13 = this.isbn13
      obj.State = this.state
      obj.Pic = this.pic
      this.data = obj
      this.put(obj)
      this.message()
      this.dialogEdit = false
      console.log(JSON.stringify(obj))
      this.Info = this.getstate(this.state)
    },
    message () {
      this.$message({
        message: '変更が完了しました。',
        type: 'success',
        duration: 2500
      })
    }
  }
}
</script>

<style scoped>
.marginTop {
  margin-top: 7%;
}

.detail {
  background: #F0F8FF;
}
.detail h3{
  padding: auto;
  text-align:left;
  margin-left: 34%;
}

.detail p{
  padding: 1%;
  text-align:left;
  margin-left: 34%;
}

.block img{
  margin-left: 5%;
  margin-top: 1.5%;
  text-align: left;
  float:left;
  width:300px;
  height:400px;
}
</style>
