<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  data() {
    return {
    }
  },
  methods: {
    getData() {
      axios
        .get("http://localhost:8080/")
        .then((response) => {
          console.log(response.data)
        })
        .catch((error) => {
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
