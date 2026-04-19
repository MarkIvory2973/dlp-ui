<script setup>
// vue
import { ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
watch(route, () => {
  active.value = route.name
})

const active = ref(route.name)
const router = useRouter()
watch(active, () => {
  router.push({ name: active.value })
})
</script>

<template>
  <VarBottomNavigation v-model:active="active" variant>
    <VarBottomNavigationItem name="parse" label="解析" icon="xml" />
    <VarBottomNavigationItem name="download" label="下载" icon="download" />
  </VarBottomNavigation>
</template>
