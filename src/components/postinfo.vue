<template>
  <div id="postinfo">
    <h1>{{ Info }}</h1>
    <div class="postinfo">
      <div class="edit">
        <h2>{{Info2}}</h2>
        <div class="form">
          <label for="title"></label>
          <el-input id="title" type="text" placeholder="Title" v-model="title" name="title"></el-input>
        </div>
        <div class="marginTop"></div>
        <div class="form-group">
          <label for="author"></label>
          <el-input id="author" type="text" placeholder="Author" v-model="author" name="author"></el-input>
        </div>
        <div class="marginTop"></div>

        <div class="form-group">

          <label for="isbn13"></label>
          <el-input id="isbn13" type="text" placeholder="Isbn13" v-model="isbn13" name="isbn13"></el-input>
        </div>

        <div class="marginTop"></div>
        <h3>{{StateInfo}}</h3>
        <div class="form-group">
          <el-switch on-text="可" off-text="不可" v-model="state"></el-switch>
        </div>
      </div>
    </div>
    <div class="marginTop"></div>
    <el-popover
    ref="popover5"
    placement="top"
    width="160"
    v-model="visible2">
    <p>保存しますか?</p>
    <div style="text-align: right; margin: 0">
      <el-button size="mini" type="text" @click="visible2 = false">キャンセル</el-button>
      <el-button type="primary" size="mini" @click="confirm()">確認</el-button>
    </div>
  </el-popover>

  <el-button v-popover:popover5>保存</el-button>
  <el-button @click="clear">リセット</el-button>
  <el-button @click="back">キャンセル</el-button>
</div>
</template>

<script>
import _ from 'lodash'
import request from 'superagent'
export default {
  data () {
    return {
      title: '',
      author: '',
      isbn13: '',
      state: false,
      Info: 'Book Edit',
      Info2: '本の情報を入力してください',
      StateInfo: '貸し出し状況',
      visible2: false
    }
  },
  methods: {
    back () {
      this.title = ''
      this.author = ''
      this.isbn13 = ''
      this.state = false
      this.$router.push({ name: 'home' })
    },
    clear () {
      this.title = ''
      this.author = ''
      this.isbn13 = ''
      this.state = false
    },
    post () {
      request
      .post('http://localhost:9090/book/postbookinfo')
      .set('Content-Type', 'application/json')
      .send(this.set())
      .end((err, res) => {
        if (err) console.log('err')
        if (res) console.log('success')
      })
    },
    set () {
      let obj = {}
      obj.Title = this.title
      obj.Author = this.author
      obj.Isbn13 = this.isbn13
      obj.State = this.state
      console.log(JSON.stringify(obj))
      return obj
    },
    formCheck () {
      let returnValue = false
      if (this.title === '') {
        returnValue = true
      } else if (this.author === '') {
        returnValue = true
      } else if (this.isbn13 === '') {
        returnValue = true
      } else if (this.isbn13.length !== 13) {
        returnValue = true
      }
      return returnValue
    },
    alertMessage () {
      this.$message({
        message: '入力ミスがあります。',
        type: 'warning'
      })
    },
    formAlert () {
      if (this.formCheck()) {
        this.alertMessage()
        this.visible2 = false
      }
    },
    confirm () {
      if (this.formCheck()) {
        this.alertMessage()
        this.visible2 = false
        return
      }
      this.storage()
      _.debounce(() => {
        this.$router.push({ name: 'home' })
      }, 1250)
      this.$message({
        message: '保存しました',
        type: 'success'
      })
      this.visible2 = false
      this.clear()
    },
    storage () {
      this.post()
    }
  }
}
</script>

<style scoped>
#postinfo h1, h2{
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #6495ED;
  margin-top: 30px;
}
.postinfo {
  background: #F5F5F5;
  padding: 3%;
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
.postinfo h3{
  color: #6495ED;
}
.form-group el-date-picker {
  width: 100%;
  text-align: left;
}
.form-group el-input {
  border: 0px;
}
.edit {
  margin: 0 auto;
  width: 70%;
}
.marginTop {
  margin-top: 2%;
}
.el-popover p{
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
