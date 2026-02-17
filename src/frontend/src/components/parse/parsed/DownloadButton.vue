<script setup>
// vue
import { ref } from 'vue'
import { storeToRefs } from 'pinia'
// local
import { useEnvStore } from '@/stores/env'

// varlet
import { Snackbar } from '@varlet/ui'

const { url, parsed } = defineProps(['url', 'parsed'])

const { baseUrl } = storeToRefs(useEnvStore())

async function download() {
  const response = await fetch(`${baseUrl}/api/download`, {
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

const formats = ref([])
</script>

<template>
  <VarMenuSelect v-model="formats" placement="bottom-end" multiple>
    <VarButtonGroup type="primary">
      <VarButton @click.stop="download" :disabled="!formats.length">
        <VarIcon name="download" />
      </VarButton>

      <VarButton style="padding: 0 6px" :disabled="!parsed.entries[0].formats.length">
        <VarIcon name="menu-down" />
      </VarButton>
    </VarButtonGroup>

    <template #options>
      <VarMenuOption
        v-for="format in parsed.entries[0].formats"
        v-bind:key="format"
        :label="format.format + ' (.' + format.ext + ')'"
        :value="format.format_id"
      />
    </template>
  </VarMenuSelect>
</template>
