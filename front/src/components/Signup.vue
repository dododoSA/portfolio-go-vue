<template>
    <div class="signup">
        <h2>Sign up</h2>
        <input type="text" placeholder="Username" v-model="username">
        <input type="password" placeholder="Password" v-model="password">
        <button @click="signUp">登録</button>
        <p>アカウントを持っている方は
            <router-link to="/login">こちら</router-link>
        </p>
    </div>
</template>

<script>
export default {
    name: 'Signup',
    data () {
        return {
            newuser:{
                username: '',
                password: ''
            }
        }
    },
    methods: {
        signUp: function() {
            if (newuser.username != '' && newuser.password != ''){
                this.$http.post('https://portfolio-go-vue.herokuapp.com/signup', newuser, function(data, status, request){
                    console.log("usercreate success")
                    console.log(status)
                    this.$router.push({name: 'HelloWorld'})
                }).error(function (data, status, request){
                    console.log("usercreate failed")
                })
            }
        }
    }
}
</script>

<style>

</style>
