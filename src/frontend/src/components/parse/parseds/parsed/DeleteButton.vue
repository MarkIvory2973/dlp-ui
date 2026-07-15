<script setup>
// vue
import { storeToRefs } from 'pinia'
// local
import { useEnvStore } from '@/stores/env'

// varlet
import { Snackbar } from '@varlet/ui'

const { baseUrl } = storeToRefs(useEnvStore())
const { url } = defineProps(['url'])
async function remove() {
  const response = await fetch(`${baseUrl.value}/api/parse`, {
    method: 'DELETE',
    body: JSON.stringify({
      url: url,
    }),
  })
  if (!response.ok) {
    switch (response.status) {
      case 404: {
        Snackbar.error('无法删除解析结果: 解析结果不存在')
        break
      }
      case 423: {
        Snackbar.warning(`无法删除解析结果: 解析任务未完成`)
        break
      }
      default: {
        const error = (await response.text()) || '未知错误'
        Snackbar.error(`无法删除解析结果: ${error}`)
        break
      }
    }
    return
  }

  Snackbar.success('成功删除解析结果')
}
</script>

<template>
  <VarButton @click="remove" round text>
    <VarIcon name="delete" />
  </VarButton>
</template>
