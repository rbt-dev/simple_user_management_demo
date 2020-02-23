import axios from 'axios'
import {
  LOGIN,
  LOGOUT,
  REGISTER,
  LOGIN_ERROR,
  LOGIN_SUCCESS,
  REGISTRATION_ERROR,
  CLEAR_STATUS
} from '../actions/auth'

const state = {
  status: '',
  user: '',
  admin: 0,
  token: localStorage.getItem('token') || ''
}

const getters = {
  isAuthenticated: state => !!state.token,
  user: state => state.user,
  admin: state => state.admin,
  status: state => state.status
}

const actions = {
  [LOGIN]: ({commit}, user) => {
    return new Promise((resolve, reject) => {
      axios.post(process.env.BASE_URL + '/authenticate', user)
        .then(resp => {
          localStorage.setItem('token', resp.data.token)
          commit(LOGIN_SUCCESS, resp.data)
          resolve(resp)
        })
        .catch(err => {
          commit(LOGIN_ERROR)
          localStorage.removeItem('token')
          reject(err)
        })
    })
  },
  [LOGOUT]: ({commit}) => {
    return new Promise((resolve, reject) => {
      localStorage.clear()
      resolve(null)
    })
  },
  [REGISTER]: ({commit}, user) => {
    return new Promise((resolve, reject) => {
      axios.post(process.env.BASE_URL + '/register', user)
        .then(resp => {
          localStorage.setItem('token', resp.data.token)
          commit(LOGIN_SUCCESS, resp.data)
          resolve(resp)
        })
        .catch(err => {
          commit(REGISTRATION_ERROR)
          localStorage.removeItem('token')
          reject(err)
        })
    })
  }
}

const mutations = {
  [LOGIN_ERROR]: state => {
    state.status = 'Login failed'
  },
  [LOGIN_SUCCESS]: (state, payload) => {
    state.user = payload.username
    state.admin = payload.admin
    state.token = payload.token
  },
  [REGISTRATION_ERROR]: state => {
    state.status = 'Registration failed'
  },
  [CLEAR_STATUS]: state => {
    state.status = ''
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
