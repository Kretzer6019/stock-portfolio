<script lang="ts">
import { defineComponent } from 'vue'
import axios, { AxiosError, AxiosResponse } from "axios";

export default defineComponent({
  data() {
    return {
    }
  },
  methods: {
    getData() {
      axios
        .get("http://localhost:8080/")
        .then((response: AxiosResponse) => {
          console.log(response.data)
        })
        .catch((error: AxiosError) => {
          console.log(error)
        })
    }
  },
  created() {
    this.getData();
  } 
})
</script>

<template>
  <div>
    <h1>Hello World!</h1>
  </div>
</template>


<style scoped>

</style>
