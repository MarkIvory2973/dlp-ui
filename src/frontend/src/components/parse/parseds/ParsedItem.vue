<script setup>
// vue
import { ref } from 'vue'

// local
import ErrorsItem from './parsed/ErrorsItem.vue'
import DeleteButton from './parsed/DeleteButton.vue'
import PreviewButton from './parsed/PreviewButton.vue'
import DownloadButton from './parsed/DownloadButton.vue'
import PreviewItem from './parsed/PreviewItem.vue'

const { url, task, errors } = defineProps(['url', 'task', 'errors'])

const floating = ref(false)
</script>

<template>
  <VarCard
    :title="
      task.entries?.length
        ? `${task.entries[0].title} (${task.entries.length})`
        : task.done
          ? '无解析结果'
          : '正在解析 URL'
    "
    :subtitle="url"
    v-model:floating="floating"
  >
    <template #description v-if="errors?.length">
      <ErrorsItem :errors="errors" />
    </template>
    <template #extra>
      <div class="extra">
        <DeleteButton :url="url" />
        <PreviewButton v-model="floating" />
        <DownloadButton :url="url" :task="task" />
      </div>
    </template>
    <template #floating-content>
      <VarDivider dashed />
      <PreviewItem :task="task" />
    </template>
  </VarCard>
</template>

<style scoped>
div.extra {
  display: flex;
  gap: 0.5rem;
}
</style>
