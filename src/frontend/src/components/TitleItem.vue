<script setup>
// vue
import { useRoute } from 'vue-router'

const route = useRoute()
</script>

<template>
  <span>{{ route.meta.title }}</span>
</template>

<style scoped>
span {
  font-size: large;
  cursor: default;
  user-select: none;
}
</style>
