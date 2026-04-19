<script setup>
// vue
import { storeToRefs } from 'pinia'
import { onMounted, onUnmounted } from 'vue'
// varlet
import { Snackbar } from '@varlet/ui'
// lodash
import { isEqual } from 'lodash'
// local
import { useEnvStore } from '@/stores/env'
import { useParsedsStore } from '@/stores/parse/parseds'

// local
import ParseItem from '@/components/parse/ParseItem.vue'
import ParsedsItem from '@/components/parse/ParsedsItem.vue'

const { baseUrl } = storeToRefs(useEnvStore())
const { parseds } = storeToRefs(useParsedsStore())
async function refresh() {
  const response = await fetch(`${baseUrl.value}/api/parse`)
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
  if (newParseds && !isEqual(parseds.value, newParseds)) {
    parseds.value = newParseds
  }
}

let timer = null
onMounted(async () => {
  await refresh()
  timer = setInterval(refresh, 200)
})
onUnmounted(() => {
  clearInterval(timer)
  timer = null
})
</script>

<template>
  <div class="stack">
    <ParseItem />
  </div>

  <p>解析结果 ({{ parseds.length }})</p>

  <div class="stack">
    <ParsedsItem />
  </div>
</template>

<style scoped>
div.stack {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}
</style>
