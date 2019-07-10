<template>
    <div class="signup">
        <h2>Sign up</h2>
        <input type="text" placeholder="Username" v-model="username">
        <input type="password" placeholder="Password" v-model="password">
        <button class="btn btn-primary"  @click="signUp">登録</button>
        <p>アカウントを持っている方は
            <router-link to="/login">こちら</router-link>
        </p>
    </div>
</template>

<script>
import axios from 'axios'
import { mapState } from 'vuex'

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
            let _this = this
            if (this.username != '' && this.password != ''){
                axios.post('/signup',{
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
                .catch(function(error){
                    console.log(error)
                })
            }
        }
    },
    computed: {
        ...mapState(['currentuserId'])
    },
    created(){
        this.$store.dispatch('GET_ME')
    }
}
</script>

<style>

</style>
