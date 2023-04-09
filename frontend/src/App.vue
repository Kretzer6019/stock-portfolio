<template>
  <div>
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

import { AxiosKey, injectStrict } from './http/symbols'

const http = injectStrict(AxiosKey)
export default defineComponent({
  data() {
    return {
    }
  },
})
</script>



<style scoped>

</style>
