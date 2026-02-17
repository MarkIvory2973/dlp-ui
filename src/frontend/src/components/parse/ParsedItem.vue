<script setup>
// vue
import { ref } from 'vue'

// local
import DeleteItem from './parsed/DeleteButton.vue'
import PreviewItem from './parsed/PreviewItem.vue'
import DownloadButton from './parsed/DownloadButton.vue'

const emit = defineEmits(['refresh'])
const { url, parsed } = defineProps(['url', 'parsed'])

const floating = ref(false)
</script>

<template>
  <VarCard
    v-model:floating="floating"
    :title="
      parsed.entries[0].title == 'parsing'
        ? '正在解析 URL'
        : parsed.entries[0].title || '无解析结果'
    "
    :subtitle="url"
  >
    <template v-if="floating" #title><div></div></template>
    <template v-if="floating" #subtitle><div></div></template>
    <template v-if="!floating && parsed.errors" #description>
      <VarSpace class="alerts" direction="column" size="large">
        <VarAlert
          v-for="error in parsed.errors"
          v-bind:key="error"
          type="danger"
          :message="error"
        />
      </VarSpace>
    </template>
    <template v-if="!floating" #extra>
      <VarSpace align="center">
        <DeleteItem @refresh="emit('refresh')" :url="url" />

        <VarButton @click="floating = !floating" round text>
          <VarIcon name="view" />
        </VarButton>

        <DownloadButton :url="url" :parsed="parsed" />
      </VarSpace>
    </template>

    <template #floating-content>
      <PreviewItem :parsed="parsed" />
    </template>
  </VarCard>
</template>

<style>
.alerts {
  margin: 1rem;
  margin-bottom: 0;
}
</style>
