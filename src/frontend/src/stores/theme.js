// vue
import { ref } from 'vue'
import { defineStore } from 'pinia'
// varlet
import { themes } from '../assets/js/varlet'

export const useThemeStore = defineStore('theme', () => {
  const theme = ref(themes.dark)

  return { theme }
})
