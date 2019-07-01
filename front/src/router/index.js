import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Signup from '@/components/Signup'
import Login from '@/components/Login'
import User from '@/components/User'
import NewProductForm from '@/components/NewProductForm'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '*',
      redirect: 'signin'
    },
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/signup',
      name: 'Signup',
      component: Signup
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/users/:id',
      component: User,
      props: route => ({
        id: String(route.params.id)
      })
    },
    {
      path: '/users/:id/products/new',
      name: 'products-new',
      component: NewProductForm,
      props: route => ({
        id: String(route.params.id)
      })
    }
  ]
})
