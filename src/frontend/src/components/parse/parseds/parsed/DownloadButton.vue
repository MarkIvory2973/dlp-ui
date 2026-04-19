<script setup>
// vue
import { storeToRefs } from 'pinia'
import { ref } from 'vue'
// varlet
import { Snackbar } from '@varlet/ui'
// local
import { useEnvStore } from '@/stores/env'

const { url, task } = defineProps(['url', 'task'])

const { baseUrl } = storeToRefs(useEnvStore())
const formats = ref([])
async function download() {
  const response = await fetch(`${baseUrl.value}/api/download`, {
    method: 'POST',
    body: JSON.stringify({
      url: url,
      format: formats.value.join('+'),
    }),
  })
  if (!response.ok) {
    switch (response.status) {
      case 409:
        Snackbar.warning('无法请求下载: 重复的下载请求')
        break

      default: {
        const error = (await response.text()) || '未知错误'
        Snackbar.error(`无法请求下载: ${error}`)
        break
      }
    }
    return
  }

  formats.value = []

  Snackbar.success('成功请求下载')
}
</script>

<template>
  <VarMenuSelect v-model="formats" placement="bottom-end" multiple scrollable>
    <VarButtonGroup type="primary">
      <VarButton @click.stop="download" :disabled="!formats.length">
        <VarIcon name="download" />
      </VarButton>
      <VarButton
        :disabled="!task.entries?.length || !task.entries[0].formats?.length"
        style="padding: 0 6px"
      >
        <VarIcon name="menu-down" />
      </VarButton>
    </VarButtonGroup>

    <template v-if="task.entries?.length && task.entries[0].formats?.length" #options>
      <VarMenuOption
        v-for="format in task.entries[0].formats"
        v-bind:key="format"
        :label="`${format.format} (.${format.ext})`"
        :value="format.format_id"
        ripple
      />
    </template>
  </VarMenuSelect>
</template>
