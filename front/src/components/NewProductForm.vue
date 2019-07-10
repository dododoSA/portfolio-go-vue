<template>
    <div class="new-product-form">
        <h2>新規投稿</h2>
        <input type="text" placeholder="ProductName" v-model="productName">
        <br/>
        <textarea placeholder="explain" v-model="intro"/>
        <br/>
        <label>
            画像を選択
            <input type="file" @change="selectedFile"/>
        </label>
        <br/>
        <button class="btn btn-primary" @click="postProduct">投稿</button>
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
            img: null
        }
    },
    methods: {
        postProduct: function() {
            let _this = this
            if (this.productName != '' && this.intro != ''){
                let formData = new FormData()
                formData.append("product_name", this.productName)
                formData.append("intro", this.intro)
                formData.append('img', this.img)
                let config = {
                    headers: {
                        'content-type' : 'multipart/form-data'
                    }
                }
                axios.post('/users/' + _this.id + '/products/create',
                    formData,
                    config
                )
                .then(function(response){
                    console.log(response)
                    _this.$router.push('/users/'+_this.id)
                })
                .catch(function(error){
                    console.log(error)
                })
            }
        },
        selectedFile: function(e) {
            e.preventDefault()
            let files = e.target.files
            this.img = files[0]
        }
    }
}
</script>

<style>

</style>
