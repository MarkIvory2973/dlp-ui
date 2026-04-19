<script setup>
// vue
import { watch, onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { RouterView } from 'vue-router'
// varlet
import { StyleProvider } from '@varlet/ui'
// local
import { useThemeStore } from './stores/theme'

// local
import TitleItem from './components/TitleItem.vue'
import NavItem from './components/NavItem.vue'

const { theme } = storeToRefs(useThemeStore())
watch(theme, () => {
  StyleProvider(theme.value)
})

onMounted(() => {
  StyleProvider(theme.value)
})
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
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin: 1rem;
  margin-bottom: 0.5rem;
}

main {
  flex: 1;
  padding: 0 1rem;
  padding-top: 0.5rem;
  overflow-y: scroll;
}
</style>
