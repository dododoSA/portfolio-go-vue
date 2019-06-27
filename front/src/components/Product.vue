<template>
    <div class="products">
        <h2>作品一覧</h2>
        <ul v-for="product in products" v-bind:key="product.id">
            <li>
                {{product.name}}
            </li>
            <li>
                {{product.intro}}
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    name: 'User',
    props: ['id'],//ユーザーのid 複数おｋにするつもり
    data: () => {
        return {
            products: {}
        }
    },
    created: function() {
        let _this = this
        axios.get("/users/" + String(_this.id) +"/products")
        .then(function(response){
            _this.products = response.data
        })
        .catch(function(error){
            console.log(error)
        })
    }
}
</script>

<style>

</style>
