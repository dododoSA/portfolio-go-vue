<template>
    <div class="login">
        <h2>Log in</h2>
        <input type="text" placeholder="Username" v-model="username">
        <input type="password" placeholder="Password" v-model="password">
        <button class="btn btn-primary" @click="logIn">ログイン</button>
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
        logIn:function(){
            let _this = this
            if (this.username != '' && this.password != ''){
                axios.post('/login', {
                    username: this.username,
                    password: this.password
                })
                .then(function(response){
                    console.log(response)
                    axios.get('/whoami')
                    .then(function (res) {
                        _this.$router.push('/users/' + res.data.user_id)
                    })
                    .catch(function (error) {
                        console.log(error)
                    })
                })
                .catch(function (error){
                    console.log(error)
                })
            }
        }
    },
    created() {
        this.$store.dispatch('GET_ME')
    }
}
</script>
