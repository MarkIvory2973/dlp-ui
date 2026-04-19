<script setup>
// local
import ErrorsItem from './download/ErrorsItem.vue'
import ProgressItem from './download/ProgressItem.vue'

const { url, task, errors } = defineProps(['url', 'task', 'errors'])
</script>

<template>
  <VarCard :title="url" :subtitle="task.done ? '下载完成' : `正在下载: ${task.title}`">
    <template #description v-if="errors?.length">
      <ErrorsItem :errors="errors" />
    </template>
    <template #extra v-if="!task.done">
      <ProgressItem :task="task" />
    </template>
  </VarCard>
</template>
