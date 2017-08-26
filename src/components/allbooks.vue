<template>
  <div id="allbooks">
    <h1>{{ Info }}</h1>
      <div class="bar">
        <el-input type="text" v-model="searchString" placeholder="Enter your search terms" icon="search" />
      </div>
      <ul>
        <li v-for="book in filteredBooks">
          <img v-bind:src="book.Pic">
          <p>{{book.Title}}</p>
        </li>
      </ul>
    <el-button type="primary" @click="push">ホームへ</el-button>
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
#allbooks h1{
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  margin-top: 60px;
  color: #6495ED;
},
ul {
  
}
</style>
