<template>
    <div class="products">
			<router-link :to="{ path: '/users/'+ id + '/products/new'}" v-if="currentuserId == id">新規投稿</router-link>
        <h2>作品一覧</h2>
        <div v-for="product in products" v-bind:key="product.id">
            <img src="../assets/aiueo.jpg" width="100" />
            <h3>
                {{product.productname}}
            </h3>
            <p>
                {{product.intro}}
            </p>
            <br/>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
import { mapState } from 'vuex';
export default {
    name: 'User',
    props: {
        id: String
    },//ユーザーのid 複数おｋにするつもり
    data: () => {
        return {
			products: {}
        }
    },
    created: function() {
        let _this = this
        axios.get("/users/" + _this.id +"/products")
        .then(function(response){
            _this.products = response.data
		})
        .catch(function(error){
            console.log(error)
		})
    },
    computed: {
        ...mapState(['currentuserId'])
    },
    methods: {
        image_src: function(productname){
            return "../assets/" + productname + ".jpg"
        }
    }
}
</script>

<style>

</style>
