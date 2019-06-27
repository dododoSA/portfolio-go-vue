<template>
<div class="user">
    <h2>ユーザー紹介</h2>
    <p>{{userdata.username}}</p>
    <br/>
    <p>{{userdata.profile}}</p>
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
