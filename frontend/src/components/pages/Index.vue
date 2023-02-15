<template>
    <div class="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div class="w-full max-w-md space-y-8">
        <div>
          <img class="mx-auto h-12 w-auto" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600" alt="Your Company" />
          <h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">Sign in to your account</h2>
        </div>
        <form class="mt-8 space-y-6">
          <input type="hidden" name="remember" value="true" />
          <div class="-space-y-px rounded-md shadow-sm">
            <div>
              <label for="email-address" class="sr-only">Email address</label>
              <input id="email-address" v-model="email" name="email" type="email" autocomplete="email" required class="relative block w-full appearance-none rounded-none rounded-t-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm" placeholder="Email address" />
            </div>
            <div>
              <label for="password" class="sr-only">Password</label>
              <input id="password" name="password"  v-model="password" type="password" autocomplete="current-password" required class="relative block w-full appearance-none rounded-none rounded-b-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm" placeholder="Password" />
            </div>
          </div>
  
          <div class="text-center">
            <button @click.prevent="create = !create" class="text-indigo-700 underline">
              <span v-if="!create">Crie sua conta</span>
              <span v-else>Fazer login</span>  
            </button>
          </div>
  
          <div>
            <button @click.prevent="submit(create)" class="group relative flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
              <span v-if="!create">Fazer Login</span> 
              <span v-else>Criar Conta</span> 
            </button>
          </div>
        </form>
      </div>
    </div>
</template>
  
<script lang="ts">
  import axios, { AxiosError, AxiosResponse } from 'axios'
  import { defineComponent } from 'vue'

  interface request {
    email: string,
    password: string,
  }

  export default defineComponent({
      data() {
          return {
            create: false,
            password: "",
            email: "",
          }
      },
      methods:{
        submit(create: Boolean){
          let request: request = {
            email: this.email,
            password: this.password,
          }

          if (create) {
            axios
              .post("http://localhost:8080/criar-conta", request)
              .then((response: AxiosResponse) => {
                if (response.status == 200) console.log(response.data);
              })
              .catch((error: AxiosError) => {
                console.log(error);
              })
          } else {
            axios
              .post("http://localhost:8080/login", request)
              .then((response: AxiosResponse) => {
                if (response.status == 200) window.location.pathname = "/usuario";
              })
              .catch((error: AxiosError) => {
                console.log(error);
              })
          }
        } 
      },
  })
</script>