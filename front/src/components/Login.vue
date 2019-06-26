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
    data () {
        return {
            username: '',
            password: ''
        }
    },
    methods: {
        logIn: function(){
            if (this.username != '' && this.password != ''){
                let formData = new FormData()
                formData.append("username", this.username)
                formData.append("password", this.password)
                
                axios.post('https://portfolio-go-vue.herokuapp.com/login', {
                    formData
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
