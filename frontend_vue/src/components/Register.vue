<template>
  <div class="content">

    <!-- registration form -->
    <form class="text-center border border-light p-5" @submit.prevent="register" oninput='password2.setCustomValidity(password2.value != password.value ? "Passwords do not match." : "")'>
      <p class="h4 mb-4">User Info</p>
      <input v-model="email" type="email" class="form-control mb-4" placeholder="E-mail" required="required" />
      <input v-model="firstname" type="text" class="form-control mb-4" placeholder="Firstname" />
      <input v-model="lastname" type="text" class="form-control mb-4" placeholder="Lastname" required="required" />
      <input v-model="password" id="password" type="password" class="form-control mb-4" placeholder="Password" required="required" />
      <input v-model="password2" id="password2" type="password" class="form-control mb-4" placeholder="Confirm password" required="required" />
      <button class="btn btn-dark btn-block my-4" type="submit">Register</button>
      <p><a v-on:click="back" class="text-secondary" href="">Back to Login</a></p>
      <div class="text-danger">{{ errorMessage }}</div>
    </form>

  </div>
</template>

<script>
import { REGISTER } from '../store/actions/auth'

export default {
  data () {
    return {
      email: '',
      firstname: '',
      lastname: '',
      password: '',
      password2: '',
      errorMessage: ''
    }
  },
  methods: {
    register: function () {
      const { email, firstname, lastname, password } = this
      this.$store
        .dispatch(REGISTER, { email, firstname, lastname, password })
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
    back: function () {
      this.$router.push('/')
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
