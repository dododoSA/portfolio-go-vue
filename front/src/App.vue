<template>
  <div id="app">
    <nav>
      <ul>
        <li v-if="currntuserId != ''"><router-link :to="{ path: '/users/'+ currentuserId}">User</router-link></li> 
        <li v-if="currntuserId == ''"><router-link to="/signup">SignUp</router-link></li>
        <li v-if="currntuserId == ''"><router-link to="/login">Login</router-link></li>
        <li v-else><a href="#" @click="logout">ログアウト</a></li>
      </ul>
    </nav>
    <img src="./assets/logo.png">
    <router-view/>
  </div>
</template>

<script>
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
