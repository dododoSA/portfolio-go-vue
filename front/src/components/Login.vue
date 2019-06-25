<template>
    <div class="login">
        <h2>Log in</h2>
        <input type="text" placeholder="Username" v-model="username">
        <input type="password" placeholder="Password" v-model="password">
        <button @click="logIn">ログイン</button>
        <p>アカウントを持ってない方は
            <router-link to="/signup">こちら</router-link>
        </p>
    </div>
</template>

<script>
export default {
    name: 'Login',
    data () {
        return {
            username: '',
            password: ''
        }
    },
    methods: {
        logIn: function(){
            if (this.username != '' && this.password != ''){
                this.$http.post('https://portfolio-go-vue.herokuapp.com/login', newuser, function(data, status, request){
                    console.log("login success")
                    console.log(status)
                    this.$router.push({name: 'HelloWorld'})
                }).error(function (data, status, request){
                    console.log("login failed")
                })
            }
        }
    }
}
</script>
