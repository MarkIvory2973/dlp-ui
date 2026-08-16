// vue
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useEnvStore = defineStore('env', () => {
  const baseUrl = ref(import.meta.env.PROD ? '' : 'http://localhost:5000')
  const baseWsUrl = ref(import.meta.env.PROD ? `ws://${location.host}` : 'ws://localhost:5000')

  return { baseUrl, baseWsUrl }
})
