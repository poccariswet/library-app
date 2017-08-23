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
      title: 'Home'
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






<!-- <template>
  <div id="P1sub2">
    <div class="detail">
      <h3>タイトル: {{ data.Title }}</h3>
      <p>名前: {{ data.Name }}</p>
      <p>スケジュール: {{ data.ScheduleDate }}</p>
      <p>内容: {{ data.Content }}</p>
    </div>
    <div class="marginTop"></div>
    <el-button type="danger" @click="sub1">戻る</el-button>
    <el-button type="success" @click="dialogEdit = true">編集</el-button>

    <el-dialog
      title="EDIT"
      :visible.sync="dialogEdit"
      size="tiny">
      <div class="marginTop"></div>
      <label>名前</label>
      <el-input
        placeholder="NAME"
        v-model="name">
      </el-input>
      <div class="marginTop"></div>
      <label>タイトル</label>
      <el-input
        type="textarea"
        autosize
        placeholder="TITLE"
        v-model="title">
      </el-input>
      <div class="marginTop"></div>
      <label>内容</label>
      <el-input
        type="textarea"
        autosize
        placeholder="CONTENT"
        v-model="content">
      </el-input>
      <div class="marginTop2"></div>
      <label>日付:</label>
      <el-date-picker
        v-model="date"
        type="date"
        placeholder="DATE">
      </el-date-picker>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogEdit = false">Cancel</el-button>
        <el-button type="success" @click="edit">Confirm</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
// import Store from '../../../store.js'
import request from 'superagent'
// import Q from 'q'
// import moment from 'moment'
// import Utils from '../../../request.js'
export default {
  data () {
    return {
      id: '',
      data: '',
      name: 'name',
      title: 'title',
      content: 'content',
      // date: moment().format('YYYY-MM-DD'),
      dialogEdit: false
    }
  },
  mounted () {
    this.call()
  },
  methods: {
    set () {
      this.name = this.data.Name
      this.title = this.data.Title
      this.content = this.data.Content
      this.date = this.data.ScheduleDate
    },
    sub1 () {
      this.$router.push({ name: 'page1.sub1' })
    },
    sub3 () {
      let id = this.$route.params.id
      this.$router.push({name: 'page1.sub3', params: {id: id}})
    },
    call () {
      // Q.fcall(() => {
        // return Utils.request2(`http://localhost:8888/schedule/getSchedules/${this.$route.params.id}`)
      // })
      .then(d => {
        console.log(JSON.stringify(d.body))
        this.data = d.body
        this.set()
      })
    },
    put (_obj) {
      request
      .put('http://localhost:8888/schedule/updateSchedule')
      .set('Content-Type', 'application/json')
      .send(_obj)
      .end((err, res) => {
        if (err) console.log('err')
        if (res) console.log('success')
      })
    },
    edit () {
      let obj = {}
      obj.Id = this.$route.params.id
      obj.Name = this.name
      obj.Title = this.title
      obj.Content = this.content
      // obj.ScheduleDate = moment(this.date).format('YYYY-MM-DD')
      this.data = obj
      this.put(obj)
      this.message()
      this.dialogEdit = false
      console.log(JSON.stringify(obj))
    },
    message () {
      this.$message({
        message: '変更しました',
        type: 'success',
        duration: 2500
      })
    }
  }
}
</script>

<style scoped>
.marginTop {
  margin-top: 5%;
}
.marginTop2 {
  margin-top: 7%;
}
.detail {
  padding: 5%;
  background: #F9FAFC;
}
.detail h3{
  text-align:left;
  margin-left: 10%;
}
.detail p{
  text-align:left;
  margin-left: 10%;
}
</style> -->
