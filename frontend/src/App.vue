<template>
  <div>
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

import type {AxiosInstance, AxiosResponse, AxiosError} from 'axios'
declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $axios: AxiosInstance
    catTags: string[]
  }
}

export default defineComponent({
  data() {
    return {
    }
  },
})
</script>



<style scoped>

</style>
