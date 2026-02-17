<script setup>
// vue
import { storeToRefs } from 'pinia'
// local
import { useEnvStore } from '@/stores/env'

// varlet
import { Snackbar } from '@varlet/ui'

const emit = defineEmits(['refresh'])
const { url } = defineProps(['url'])

const { baseUrl } = storeToRefs(useEnvStore())

async function remove() {
  const response = await fetch(`${baseUrl.value}/api/parse`, {
    method: 'DELETE',
    body: JSON.stringify({
      url: url,
    }),
  })
  if (!response.ok) {
    switch (response.status) {
      default: {
        const error = (await response.text()) || '未知错误'
        Snackbar.error(`无法删除解析结果: ${error}`)
        break
      }
    }
    return
  }

  emit('refresh')

  Snackbar.success('成功删除解析结果')
}
</script>

<template>
  <VarButton @click="remove" round text>
    <VarIcon name="delete" />
  </VarButton>
</template>
