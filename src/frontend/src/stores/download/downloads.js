// vue
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useDownloadsStore = defineStore('downloads', () => {
  const downloads = ref({})

  return { downloads }
})
