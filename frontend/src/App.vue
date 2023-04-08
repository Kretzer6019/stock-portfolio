<template>
  <div>
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { globalVariables, GlobalVariables } from './globals/variables';

export default defineComponent({
  data() {
    return {
    }
  },
  provide() {
    return {
      globalVariables: globalVariables as GlobalVariables
    };
  }
})
</script>



<style scoped>

</style>
