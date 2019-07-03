import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

import {
  GET_ME, GET_FAILURE, GET_SUCCESS, LOGOUT
} from './mutation-types.js'

Vue.use(Vuex)

Vue.config.devtools = true

const state = {
  currentuserId: ''
}

const mutations = {
  [GET_SUCCESS] (state, data) {
    console.log('GET_ME_SUCCESS')
    state.currentuserId = data.user_id
  },
  [GET_FAILURE] (state, error) {
    console.log(error)
    state.currentuserId = null
  },
  [LOGOUT] (state) {
    state.currentuserId = ''
  }
}

const actions = {
  [GET_ME] () {
    let _this = this
    axios.get('/whoami')
      .then(function (res) {
        _this.commit(GET_SUCCESS, res.data)
      })
      .catch(function (error) {
        _this.commit(GET_FAILURE, error)
      })
  }
}

const getters = {
  currentuserId: state => {
    return state.currentuserId
  },
  isLoggedIn: state => {
    return state.currentuserId != null
  }
}

export default new Vuex.Store({
  state: state,
  mutations: mutations,
  actions: actions,
  getters: getters
})
