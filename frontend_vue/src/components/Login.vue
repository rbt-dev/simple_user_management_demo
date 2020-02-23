<template>
  <div class="content">

    <!-- login form -->
    <form class="text-center border border-light p-5" @submit.prevent="login">
      <p class="h4 mb-4">Simple User Management</p>
      <input v-model="email" v-on:input="onChange" type="email" class="form-control mb-4" placeholder="E-mail" required="required" />
      <input v-model="password" v-on:input="onChange" type="password" class="form-control mb-4" placeholder="Password" required="required" />
      <button class="btn btn-dark btn-block my-4" type="submit">Login</button>
      <p>Not a member? <a v-on:click="register" class="text-secondary" href="">Register</a></p>
      <div class="text-danger">{{ errorMessage }}</div>
    </form>

  </div>
</template>

<script>
import { LOGIN } from '../store/actions/auth'

export default {
  data () {
    return {
      email: '',
      password: '',
      errorMessage: ''
    }
  },
  methods: {
    login: function () {
      const { email, password } = this
      this.$store
        .dispatch(LOGIN, { 'username': email, password })
        .then((resp) => {
          if (resp.status === 200) {
            this.$router.push('/users')
          } else {
            this.errorMessage = this.$store.getters.status
          }
        })
        .catch(err => {
          console.log(err)
          this.errorMessage = this.$store.getters.status
        })
    },
    onChange: function () {
      if (this.errorMessage.length > 0) {
        this.errorMessage = ''
      }
    },
    register: function () {
      this.$router.push('/register')
    }
  }
}
</script>

<style scoped>
.content {
  display:table-cell;
  vertical-align:middle;
}
form {
  display:table;
  margin:auto;
  box-shadow:
    0 2px 5px 0 rgba(0,0,0,0.16),
    0 2px 10px 0 rgba(0,0,0,0.12);
  padding-bottom: 2rem !important;
}
.h4 {
  width: 300px;
}
.text-danger {
  height: 1rem;
}
</style>
