<template>
    <div class="new-product-form">
        <h2>新規投稿</h2>
        <input type="text" placeholder="ProductName" v-model="productName">
        <textarea placeholder="explain" v-model="intro"/>
        <button @click="postProduct">投稿</button>
    </div>
</template>

<script>
import axios from 'axios'
export default {
    name: 'NewProductForm',
    props: {
      id: String
    },
    data () {
        return {
            productName: '',
            intro: '',
        }
    },
    methods: {
        postProduct: function() {
            let _this = this
            if (this.productName != '' && this.intro != ''){
                axios.post('/users/' + _this.id + '/products/create',{
                    product_name: _this.productName,
                    intro: _this.intro
                })
                .then(function(response){
                    console.log(response)
                    _this.$router.push('/users/'+_this.id)
                })
                .catch(function(error){
                    console.log(error)
                })
            }
        }
    }
}
</script>

<style>

</style>
