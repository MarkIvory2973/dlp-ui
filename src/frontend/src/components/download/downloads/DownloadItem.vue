<script setup>
// local
import ErrorsItem from './download/ErrorsItem.vue'
import ProgressItem from './download/ProgressItem.vue'

const { url, job, errors } = defineProps(['url', 'job', 'errors'])
</script>

<template>
  <VarCard :title="url" :subtitle="job.done ? '下载完成' : `正在下载: ${job.title}`">
    <template #description v-if="errors?.length">
      <ErrorsItem :errors="errors" />
    </template>
    <template #extra v-if="!job.done">
      <ProgressItem :job="job" />
    </template>
  </VarCard>
</template>
