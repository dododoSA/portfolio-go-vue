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
        logIn:function(){
            let _this = this
            let flag = false
            if (this.username != '' && this.password != ''){
                axios.post('/login', {
                    username: this.username,
                    password: this.password
                })
                .then(function(response){
                    console.log(response)
                    if(response.status == 200){
                        flag = true
                    }
                    console.log(flag)
                    _this.$store.dispatch('GET_ME')
                })
                .catch(function (error){
                    console.log(error)
                })
                console.log(flag)
                if(flag){
                    _this.$router.push('/users/' + _this.$store.getters.currentuserId)
                }
            }
        }
    },
    created() {
        this.$store.dispatch('GET_ME')
    }
}
</script>
