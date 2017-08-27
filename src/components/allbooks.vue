<template>
  <div id="allbooks">
    <h1>{{ Info }}</h1>
      <div class="bar">
        <el-input type="text" v-model="searchString" placeholder="Enter your search book" icon="search" />
      </div>
      <div class="block">
        <ul v-for="book in filteredBooks">
          <img v-bind:src="book.Pic">
          <p>{{book.Title}}</p>
      </ul>
      <!-- <#pagenation>-->
    </div>
    <el-button type="primary" @click="push">ホームへ</el-button>
    <el-button type="primary" @click="push_PostInfo">本の追加へ</el-button>
  </div>
</template>

<script>
import Utils from '../request'
import Q from 'q'
export default {
  data () {
    return {
      Info: 'AllBooksInfo',
      searchString: '',
      // page: 0,
      // ItemSize: 100,
      Books: [
        {
          Id: '',
          Title: '',
          Author: '',
          Isbn13: '',
          State: false,
          Pic: ''
        }
      ]
    }
  },
  mounted () {
    this.set()
    this.call()
  },
  methods: {
    push () {
      this.$router.push({ name: 'home' })
    },
    push_PostInfo () {
      this.$router.push({ name: 'postinfo' })
    },
    set () {
    },
    call () {
      Q.fcall(() => {
        return Utils.request2('http://localhost:9090/book/pullbooksinfo/all')
      })
      .then(d => {
        this.Books = d.body
      })
    }
  },
  computed: {
    filteredBooks: function () {
      var Bookarray = this.Books
      var searchString = this.searchString

      if (!searchString) {
        return Bookarray
      }

      searchString = searchString.trim().toLowerCase()

      Bookarray = Bookarray.filter(function (item) {
        if (item.Title.toLowerCase().indexOf(searchString) !== -1) {
          return item
        }
      })
      return Bookarray
    }
  }
}
</script>

<style scoped>
#allbooks{
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  margin-top: 10px;
}
/*topの色*/
h1 {
  color: #6495ED;
}

ul{
  text-align: left;
}

ul img {
  width:50px;
  height:70px;
  float:left;
  border:none;
}

ul p {
  margin-left: 50px;
  margin: 30px;
  font-weight: bold;
  padding-top: 30px;
  padding: 20px;
  color:#6e7a7f;
}

p:hover{
  color: #FF6347;
  cursor: pointer;
}
</style>

<!--
<el-pagination
  layout="prev, pager, next"
  :total="this.ItemSize">
</el-pagination>
 -->
