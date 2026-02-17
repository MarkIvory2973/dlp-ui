// vue
import { defineStore } from 'pinia'
import { ref } from 'vue'
// varlet
import { themes } from '../assets/js/varlet'
import { StyleProvider } from '@varlet/ui'

export const useThemeStore = defineStore('theme', () => {
  const theme = ref(themes.dark)

  function changeTheme() {
    if (theme.value == themes.dark) {
      theme.value = themes.light
    } else {
      theme.value = themes.dark
    }

    StyleProvider(theme.value)
  }

  return { theme, changeTheme }
})
