<script setup>
// vue
import { storeToRefs } from 'pinia'
import { onMounted, onUnmounted } from 'vue'
// lodash
import isEqual from 'lodash/isEqual'
// local
import { useEnvStore } from '@/stores/env'
import { useDownloadsStore } from '@/stores/download/downloads'

// varlet
import { Snackbar } from '@varlet/ui'
// local
import DownloadItem from '@/components/download/DownloadItem.vue'

const { baseUrl } = storeToRefs(useEnvStore())
const { downloads } = storeToRefs(useDownloadsStore())

async function refresh() {
  const response = await fetch(`${baseUrl.value}/api/download`)
  if (!response.ok) {
    switch (response.status) {
      default: {
        const error = (await response.text()) || '未知错误'
        Snackbar.error(`无法刷新下载队列: ${error}`)
        break
      }
    }
    return
  }

  const newDownloads = await response.json()
  if (!isEqual(downloads.value, newDownloads)) {
    downloads.value = newDownloads
  }
}

let timer = null
onMounted(async () => {
  await refresh()
  timer = setInterval(refresh, 1000)
})
onUnmounted(() => {
  clearInterval(timer)
  timer = null
})
</script>

<template>
  <VarSpace direction="column" size="large">
    <DownloadItem
      v-for="url in Object.keys(downloads)"
      v-bind:key="url"
      :url="url"
      :download="downloads[url]"
    />
  </VarSpace>
</template>

<style scoped></style>
