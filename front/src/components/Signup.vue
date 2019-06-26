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
import axios from 'axios'
export default {
    name: 'Signup',
    data () {
        return {
            username: '',
            password: ''
        }
    },
    methods: {
        signUp: function() {
            if (this.username != '' && this.password != ''){
                axios.post('/signup',{
                    username: this.username,
                    password: this.password
                })
                .then(function(response){
                    console.log(response)
                })
                .catch(function(error){
                    console.log(error)
                })
                this.router.push({name : 'HelloWorld'})
            }
        }
    }
}
</script>

<style>

</style>
