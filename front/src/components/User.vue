<template>
<div class="user">
    <h2>ユーザー紹介</h2>
    <p>{{userdata.username}}</p>
    <br/>
    <p v-if="userdata.profile != ''">{{userdata.profile}}</p>
        <a href="#" @click="logout" v-if="currentuserId == id">ログアウト</a>
    <product :id="id"></product>
</div>
</template>

<script>
import axios from 'axios'
import Product from './Product.vue'
export default {
    name: 'User',
    props: {
        id: String
    },
    data: () => {
        return {
            userdata: {}
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
            this.$route.push('/')
        }
    },
    created: function() {
        let _this = this
        axios.get("/users/" + _this.id)
        .then(function(response) {
            _this.userdata = response.data
        })
        .catch(function(error) {
            console.log(error)
        })
    },
    components: {
        Product
    }
}
</script>

<style>

</style>
