// vue
import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useDownloadsStore = defineStore('downloads', () => {
  const downloads = ref([])

  return { downloads }
})
