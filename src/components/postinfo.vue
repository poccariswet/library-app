<template>
  <div id="postinfo">
    <h1>{{Info}}</h1>
    <el-form :model="bookData" :rules="rules" ref="bookData" label-width="100px" class="postinfo">
      <h2>{{Info2}}</h2>
      <el-form-item label="Title" prop="title">
        <el-input v-model="bookData.title" placeholder="book title"></el-input>
      </el-form-item>
      <el-form-item label="Author" prop="author">
        <el-input v-model="bookData.author" placeholder="book author"></el-input>
      </el-form-item>
      <el-form-item label="Isbn13" prop="isbn13">
        <el-input v-model="bookData.isbn13" placeholder="book isbn 13 numbers"></el-input>
      </el-form-item>
      <el-form-item label="貸し出し状況" prop="state">
        <el-switch on-text="可" off-text="不可" v-model="bookData.state"></el-switch>
      </el-form-item>
      <el-form-item label="Address" prop="pic" >
        <el-input v-model="bookData.pic" placeholder="picture address"></el-input>
      </el-form-item>
    </el-form>
    <div class="button">
      <el-button type="primary" @click="confirm('bookData')">保存</el-button>
      <el-button @click="resetForm()">リセット</el-button>
      <el-button @click="back()">戻る</el-button>
    </div>
  </div>
</template>

<script>
import _ from 'lodash'
import request from 'superagent'
export default {
  data () {
    return {
      bookData: {
        title: '',
        author: '',
        isbn13: '',
        state: false,
        pic: ''
      },
      rules: {
        title: [
          { required: true, message: 'Please input book name', trigger: 'blur' }
        ],
        author: [
          { required: true, message: 'Please input book author name', trigger: 'blur' }
        ],
        isbn13: [
          { required: true, message: 'Please input book isbn only 13 numbers', trigger: 'blur' },
          { min: 13, max: 13, message: 'only 13 numbers!', trigger: 'blur' }
        ]
      },
      Info: 'Book Post',
      Info2: '本の情報を入力してください'
    }
  },
  methods: {
    back () {
      this.title = ''
      this.author = ''
      this.isbn13 = ''
      this.state = false
      this.pic = ''
      this.push_home()
    },
    resetForm () {
      this.bookData.title = ''
      this.bookData.author = ''
      this.bookData.isbn13 = ''
      this.bookData.state = false
      this.bookData.pic = ''
    },
    post () {
      request
      .post('http://localhost:9090/book/api/post')
      .set('Content-Type', 'application/json')
      .send(this.set())
      .end((err, res) => {
        if (err) console.log('err')
        if (res) console.log('success')
      })
    },
    set () {
      let obj = {}
      obj.Title = this.bookData.title
      obj.Author = this.bookData.author
      obj.Isbn13 = this.bookData.isbn13
      obj.State = this.bookData.state
      obj.Pic = this.bookData.pic
      console.log(JSON.stringify(obj))
      return obj
    },
    push_home () {
      this.$router.push({name: 'home'})
    },
    confirm (formData) {
      if (
        this.bookData.title !== '' && this.bookData.author !== '' && this.bookData.isbn13.length === 13
      ) {
        if (this.bookData.pic === '') {
          this.bookData.pic = 'http://cdn.tsutaya.tsite.jp/static/tsite_common/images/noimage.png'
        }
        this.store()
        _.debounce(() => {
          this.push_home()
        }, 1250)
        this.$message({
          message: '保存しました',
          type: 'success'
        })
        this.resetForm()
      }
    },
    store () {
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

.button {
  margin-top: 5%;
}
</style>
