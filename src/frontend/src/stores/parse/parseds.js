// vue
import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useParsedsStore = defineStore('parseds', () => {
  const parseds = ref([])

  return { parseds }
})
