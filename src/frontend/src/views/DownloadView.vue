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
import { useDownloadsStore } from '@/stores/download/downloads'

// local
import DownloadsItem from '@/components/download/DownloadsItem.vue'

const { baseUrl } = storeToRefs(useEnvStore())
const { downloads } = storeToRefs(useDownloadsStore())
async function refresh() {
  const response = await fetch(`${baseUrl.value}/api/download`)
  if (!response.ok) {
    switch (response.status) {
      default: {
        const error = (await response.text()) || '未知错误'
        Snackbar.error(`无法刷新下载结果: ${error}`)
        break
      }
    }
    return
  }

  const newDownloads = await response.json()
  if (newDownloads && !isEqual(downloads.value, newDownloads)) {
    downloads.value = newDownloads
  }
}

const { baseWsUrl } = storeToRefs(useEnvStore())
let backend = null
onMounted(async () => {
  await refresh()

  if (backend) {
    return
  }

  backend = new WebSocket(`${baseWsUrl.value}/backend/download`)
  backend.onmessage = async function (event) {
    if (event.data != 'refresh') {
      return
    }

    await refresh()
  }
})
onUnmounted(() => {
  if (!backend) {
    return
  }

  backend.close(1000)
  backend = null
})
</script>

<template>
  <VarAlert title="下载文件: ./down" />

  <p>下载队列 ({{ downloads.length }})</p>

  <div class="stack">
    <DownloadsItem />
  </div>
</template>

<style scoped>
div.stack {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}
</style>
