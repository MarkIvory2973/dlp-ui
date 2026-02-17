<script setup>
// vue
import { storeToRefs } from 'pinia'
import { onMounted, onUnmounted } from 'vue'
// lodash
import isEqual from 'lodash/isEqual'
// local
import { useEnvStore } from '@/stores/env'
import { useParsedsStore } from '@/stores/parse/parseds'

// varlet
import { Snackbar } from '@varlet/ui'
// local
import ParseItem from '@/components/parse/ParseItem.vue'
import ParsedItem from '@/components/parse/ParsedItem.vue'

const { baseUrl } = storeToRefs(useEnvStore())
const { parseds } = storeToRefs(useParsedsStore())

async function refresh() {
  const response = await fetch(`${baseUrl}/api/parse`)
  if (!response.ok) {
    switch (response.status) {
      default: {
        const error = (await response.text()) || '未知错误'
        Snackbar.error(`无法刷新解析结果: ${error}`)
        break
      }
    }
    return
  }

  const newParseds = await response.json()
  if (!isEqual(parseds.value, newParseds)) {
    parseds.value = newParseds
  }
}

let timer = null
onMounted(async () => {
  await refresh()
  timer = setInterval(refresh, 5000)
})
onUnmounted(() => {
  clearInterval(timer)
  timer = null
})
</script>

<template>
  <VarSpace direction="column" size="large">
    <ParseItem />

    <span>解析结果 ({{ Object.entries(parseds).length }})</span>

    <ParsedItem
      v-for="url in Object.keys(parseds)"
      v-bind:key="url"
      @refresh="refresh"
      :url="url"
      :parsed="parseds[url]"
    />
  </VarSpace>
</template>

<style scoped>
span {
  cursor: default;
  user-select: none;
}
</style>
