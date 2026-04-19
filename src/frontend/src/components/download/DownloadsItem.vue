<script setup>
// vue
import { storeToRefs } from 'pinia'
// local
import { useDownloadsStore } from '@/stores/download/downloads'

// local
import DownloadItem from './downloads/DownloadItem.vue'

const { downloads } = storeToRefs(useDownloadsStore())
</script>

<template>
  <DownloadItem
    v-for="download in downloads"
    v-bind:key="download"
    :url="download.url"
    :task="download.task"
    :errors="download.errors"
  />
</template>
