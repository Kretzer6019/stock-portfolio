<script lang="ts">
import { defineComponent } from 'vue'
import axios from 'axios';

export default defineComponent({
  data() {
    return {
    }
  },
  methods: {
    async getData() {
      try {
          const response = await axios.get('http://localhost:8080/');
          console.log(response.data);
      } catch (error) {
          console.error(error);
      }
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
