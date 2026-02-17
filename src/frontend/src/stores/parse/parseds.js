// vue
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useParsedsStore = defineStore('parseds', () => {
  const parseds = ref({})

  return { parseds }
})
