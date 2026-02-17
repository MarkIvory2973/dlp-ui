<script setup>
// vue
import { RouterView } from 'vue-router'
import { storeToRefs } from 'pinia'
// varlet
import { StyleProvider } from '@varlet/ui'
// local
import { useThemeStore } from './stores/theme'

// local
import TitleItem from './components/TitleItem.vue'
import NavItem from './components/NavItem.vue'

const { theme } = storeToRefs(useThemeStore())
StyleProvider(theme.value)
</script>

<template>
  <header>
    <TitleItem />
  </header>

  <main>
    <RouterView />
  </main>

  <footer>
    <NavItem />
  </footer>
</template>

<style scoped>
header {
  margin: 1rem;
}

main {
  flex: 1;
  padding: 0 1rem;
  padding-top: 4px;
  overflow-y: scroll;
}
</style>
