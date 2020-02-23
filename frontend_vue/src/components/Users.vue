<template>
  <div>
    <!-- header: show current user and logout button -->
    <nav class="navbar navbar-dark bg-dark">
      <p class="navbar-text mb-0">
        <b-icon icon="person-fill" class="h4 mb-0"></b-icon>
        {{ user }}
      </p>
      <div>
        <ul class="navbar-nav mr-auto logout">
          <li class="nav-item">
            <a class="nav-link" v-on:click="logout">
              <b-icon icon="x-square-fill" class="h4 mb-0"></b-icon>
              Logout
            </a>
          </li>
        </ul>
      </div>
    </nav>

    <!-- users table -->
    <b-table small :fields="fields" :items="items">
      <template v-slot:cell(actions)="data" v-if="admin">
        <a v-if="data.item.email != user" v-on:click="deleteUser(data.item.id)"><b-icon icon="trash" class="h5 mb-0 mr-2 icon"></b-icon></a>
        <!-- can't delete yourself -->
        <a v-if="data.item.email == user"><b-icon icon="trash" class="h5 mb-0 mr-2 icon" style="color:#cdcdcd;"></b-icon></a>
        <a v-on:click="editUser(data.item)"><b-icon icon="pencil" class="h5 mb-0 icon"></b-icon></a>
      </template>
    </b-table>

    <!-- modal for editing user -->
    <b-modal id="edit-user-modal" hide-footer>
      <template v-slot:modal-title>
        Edit user
      </template>
      <div class="d-block">
        <form class="text-center border border-light py-4 px-5" @submit.prevent="updateUser" oninput='password2.setCustomValidity(password2.value != password.value ? "Passwords do not match." : "")'>
          <input v-model="update.id" type="hidden" />
          <input v-model="update.email" type="email" class="form-control mb-4" placeholder="E-mail" required="required" />
          <input v-model="update.firstname" type="text" class="form-control mb-4" placeholder="Firstname" />
          <input v-model="update.lastname" type="text" class="form-control mb-4" placeholder="Lastname" required="required" />
          <input v-model="update.password" id="password" type="password" class="form-control mb-4" placeholder="Password" />
          <input id="password2" type="password" class="form-control mb-5" placeholder="Confirm password" />
          <button class="btn btn-dark btn-block" type="submit">Update</button>
        </form>
      </div>
    </b-modal>

  </div>
</template>

<script>
import {
  USERS_GET,
  USER_UPDATE,
  USER_DELETE
} from '../store/actions/users'
import { LOGOUT } from '../store/actions/auth'

export default {
  data () {
    return {
      user: this.$store.getters.user,
      admin: this.$store.getters.admin,
      items: this.$store.getters.users,
      fields: [
        'id',
        'email',
        'firstname',
        'lastname',
        'created',
        'modified',
        { key: 'actions', label: '' }
      ],
      update: {
        id: '',
        email: '',
        firstname: '',
        lastname: '',
        password: ''
      }
    }
  },
  created () {
    this.getUsers()
  },
  methods: {
    getUsers: function () {
      this.$store
        .dispatch(USERS_GET)
        .then(() => {
          this.items = this.$store.getters.users
        })
        .catch(err => {
          console.log(err)
          this.$router.push('/')
        })
    },
    deleteUser: function (id) {
      if (confirm('User will be deleted.')) {
        this.$store
          .dispatch(USER_DELETE, id)
          .then(() => this.getUsers())
      }
    },
    editUser: function (user) {
      this.update.id = user.id
      this.update.email = user.email
      this.update.firstname = user.firstname
      this.update.lastname = user.lastname
      this.$bvModal.show('edit-user-modal')
    },
    updateUser: function () {
      const { id, email, firstname, lastname, password } = this.update
      this.$store
        .dispatch(USER_UPDATE, { id, email, firstname, lastname, password })
        .then(() => {
          this.getUsers()
          this.$bvModal.hide('edit-user-modal')
        })
        .catch(err => {
          console.log(err)
          this.errorMessage = this.$store.getters.status
        })
    },
    logout: function () {
      this.$store
        .dispatch(LOGOUT)
        .then(() => this.$router.push('/'))
    }
  }
}
</script>

<style>
  .table {
    width: -webkit-calc(100% - 100px) !important;
    width: -moz-calc(100% - 100px) !important;
    width: calc(100% - 100px) !important;
    margin: 50px;
  }
  .modal-backdrop {
    opacity: 0.5;
  }
  .logout, .icon {
    cursor: pointer;
  }
</style>
