<script setup>
// vue
import { storeToRefs } from 'pinia'
import { useRoute } from 'vue-router'
// lodash
import { isEqual } from 'lodash'
// local
import { useThemeStore } from '@/stores/theme'
import { themes } from '@/assets/js/varlet'

const route = useRoute()

const { theme } = storeToRefs(useThemeStore())
function changeTheme() {
  if (isEqual(theme.value, themes.dark)) {
    theme.value = themes.light
  } else if (isEqual(theme.value, themes.light)) {
    theme.value = themes.dark
  }
}
</script>

<template>
  <p>{{ route.meta.title }}</p>

  <VarButton @click="changeTheme" round text>
    <VarIcon v-if="isEqual(theme, themes.dark)" name="white-balance-sunny" />
    <VarIcon v-else-if="isEqual(theme, themes.light)" name="weather-night" />
  </VarButton>
</template>

<style scoped>
p {
  margin: 0;
  font-size: large;
}
</style>
