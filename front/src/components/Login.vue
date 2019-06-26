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
import axios from 'axios'
export default {
    name: 'Login',
    data: ()=> {
        return {
            username: '',
            password: ''
        }
    },
    methods: {
        logIn: function(){
            if (this.username != '' && this.password != ''){
                axios.post('/login', {
                    username: this.username,
                    password: this.password
                })
                .then(function(){
                    console.log(response)
                })
                .catch(function (error){
                    console.log(error)
                })
            }
        }
    }
}
</script>
