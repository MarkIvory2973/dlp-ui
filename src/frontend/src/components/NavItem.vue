<script setup>
// vue
import { ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
watch(route, (route) => {
  active.value = route.name
})

const active = ref('parse')
const router = useRouter()
watch(active, (name) => {
  router.push({ name: name })
})
</script>

<template>
  <VarBottomNavigation v-model:active="active" variant>
    <VarBottomNavigationItem name="parse" label="解析" icon="xml" />
    <VarBottomNavigationItem name="download" label="下载" icon="download" />
  </VarBottomNavigation>
</template>
