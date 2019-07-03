import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

import {
  GET_ME, GET_FAILURE, GET_SUCCESS
} from './mutation-types.js'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    currentuserId: ''
  },
  mutations: {
    [GET_SUCCESS] (state, data) {
      console.log('GET_ME_SUCCESS')
      state.currentuserId = data.user_id
      if (state.currentuserId === '') {
        state.currentuserId = null
      }
    },
    [GET_FAILURE] (state, error) {
      console.log(error)
    }
  },
  actions: {
    [GET_ME] () {
      axios.get('/whoami')
        .then(function (res) {
          this.commit(GET_SUCCESS, res.data)
        })
        .catch(function (error) {
          this.commit(GET_FAILURE, error)
        })
    }
  },
  getters: {
    currentuserId (state) {
      return state.currentuserId
    }
  }
})
