<template>
    <div class="products">
			<router-link :to="{ path: '/users/'+ id + '/products/new'}" v-if="currentuserId == id">新規投稿</router-link>
        <h2>作品一覧</h2>
        <ul v-for="product in products" v-bind:key="product.id">
            <li>
                {{product.productname}}
            </li>
            <li>
                {{product.intro}}
            </li>
        </ul>
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
    }
}
</script>

<style>

</style>
