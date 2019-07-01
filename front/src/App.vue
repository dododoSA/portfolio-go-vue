<template>
  <div id="app">
    <nav>
      <ul style="list-style: none;">
        <li v-if="currntuserId != null"><a href="#" @click="user">User</a></li> 
        <li v-if="currntuserId == null"><a href="#" @click="signup">SignUp</a></li>
        <li v-if="currntuserId == null"><a href="#" @click="login">Login</a></li>
        <li v-if="currntuserId != null"><a href="#" @click="logout">ログアウト</a></li>
      </ul>
    </nav>
    <img src="./assets/logo.png">
    <router-view/>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'App',
  data(){
    return {
      currentuserId: ''
    }
  },
  methods: {
      logout: function (){
          axios.post("/logout")
          .then(function(response){
              console.log(response)
          })
          .catch(function(error){
              console.log(error)
          })
          this.$router.push('/')
      },
      //なぜかrouter-linkできないので
      signup: function() {
        this.$router.push('/signup')
      },
      user: function() {
        this.$router.push('/users/' + this.currentuserId)
      },
      login: function() {
        this.$router.push('/login')
      }
  },
  created: function(){
    let _this = this
    axios.get('/whoami')
    .then(function(res){
      _this.currentuserId = res.data.user_id
    })
    .catch(function(error){
      console.log(error)
        _this.currentuserId = null
    })
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
