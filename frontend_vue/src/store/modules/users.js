import axios from 'axios'
import {
  USERS_GET,
  USERS_GET_ERROR,
  USERS_GET_SUCCESS,
  USER_UPDATE,
  USER_DELETE
} from '../actions/users'

const state = {
  users: []
}

const getters = {
  users: state => state.users
}

const actions = {
  [USERS_GET]: ({commit}) => {
    return new Promise((resolve, reject) => {
      axios.get(process.env.BASE_URL + '/users', {headers: {Authorization: `Bearer ${localStorage.getItem('token')}`}})
        .then(resp => {
          commit(USERS_GET_SUCCESS, resp.data)
          resolve(resp)
        })
        .catch(err => {
          commit(USERS_GET_ERROR)
          reject(err)
        })
    })
  },
  [USER_UPDATE]: ({commit}, user) => {
    return new Promise((resolve, reject) => {
      axios.put(process.env.BASE_URL + '/user', user, {headers: {Authorization: `Bearer ${localStorage.getItem('token')}`}})
        .then(resp => {
          resolve(resp)
        })
        .catch(err => {
          reject(err)
        })
    })
  },
  [USER_DELETE]: ({commit}, id) => {
    return new Promise((resolve, reject) => {
      axios.delete(process.env.BASE_URL + '/user', {headers: {Authorization: `Bearer ${localStorage.getItem('token')}`}, params: {id: id}})
        .then(resp => {
          resolve(resp)
        })
        .catch(err => {
          reject(err)
        })
    })
  }
}

const mutations = {
  [USERS_GET_SUCCESS]: (state, payload) => {
    state.users = payload
  },
  [USERS_GET_ERROR]: state => {
    state.status = 'Error when fetching users.'
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
