// vue
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useEnvStore = defineStore('env', () => {
  const baseUrl = ref(import.meta.env.PROD ? '' : 'http://localhost:5000')

  return { baseUrl }
})
