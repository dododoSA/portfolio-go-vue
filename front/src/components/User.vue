<template>
<div class="user">
    <h2>ユーザー紹介</h2>
    <p>{{userdata.username}}</p>
    <br/>
    <p v-if="userdata.profile != ''">{{userdata.profile}}</p>
    
    <product :id="id"></product>
</div>
</template>

<script>
import axios from 'axios'
import Product from './Product.vue'
import { mapState } from 'vuex';
export default {
    name: 'User',
    props: {
        id: String
    },
    data: () => {
        return {
            userdata: {},
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
        
        this.$store.dispatch('GET_ME')
    },
    components: {
        Product
    }
}
</script>

<style>

</style>
