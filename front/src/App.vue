<template>
  <div id="app">
    <nav>
      <ul style="list-style: none;">
        <li><router-link to="/">Home</router-link></li>
        <li v-if="this.$store.getters.isLoggedIn"><a href="#" @click="user">User</a></li> 
        <li v-if="!this.$store.getters.isLoggedIn"><a href="#" @click="signup">SignUp</a></li>
        <li v-if="!this.$store.getters.isLoggedIn"><a href="#" @click="login">Login</a></li>
        <li v-if="this.$store.getters.isLoggedIn"><a href="#" @click="logout">ログアウト</a></li>
      </ul>
    </nav>
    <img src="./assets/logo.png">
    <router-view/>
  </div>
</template>

<script>
import axios from 'axios'
import {mapState, mapMutations} from 'vuex'

export default {
  name: 'App',
  methods: {
      ...mapMutations({LOGOUT}),
      logout: function (){
        let _this = this
          axios.post("/logout")
          .then(function(response){
              console.log(response)
              _this.LOGOUT()
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
        this.$router.push('/users/' + this.$store.getters.currentuserId)
      },
      login: function() {
        this.$router.push('/login')
      }
  },
  created: function(){
    this.$store.dispatch('GET_ME')
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
