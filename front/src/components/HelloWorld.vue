<template>
  <div class='hello'>
    <h1>{{ msg }}</h1>
    <h2>Essential Links</h2>
    <button @click="apiHello">hello</button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'HelloWorld',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  },
  methods: {
    apiHello: function () {
      axios.get('https://portfolio-go-vue.herokuapp.com/hello')
        .then(response => {
          if (response.status === 200){
            this.msg = response.data
          }
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
