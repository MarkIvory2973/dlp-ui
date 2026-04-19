<script setup>
// vue
import { ref } from 'vue'
import { storeToRefs } from 'pinia'
// varlet
import z from 'zod'
import { Snackbar } from '@varlet/ui'
// local
import { useEnvStore } from '@/stores/env'

const urlsForm = ref(null)
const urls = ref('')
const { baseUrl } = storeToRefs(useEnvStore())
function parse() {
  urls.value.split('\n').forEach(async (url, index) => {
    const response = await fetch(`${baseUrl.value}/api/parse`, {
      method: 'POST',
      body: JSON.stringify({
        url: url,
      }),
    })
    if (!response.ok) {
      switch (response.status) {
        case 409:
          setTimeout(() => Snackbar.warning('无法请求解析 URL: 重复的解析请求'), index * 3000)
          break

        default: {
          const error = (await response.text()) || '未知错误'
          setTimeout(() => Snackbar.error(`无法请求解析 URL: ${error}`), index * 3000)
          break
        }
      }
      return
    }

    setTimeout(() => Snackbar.success('成功请求解析 URL'), index * 3000)
  })

  urlsForm.value.reset()
}
</script>

<template>
  <VarInput
    ref="urlsForm"
    v-model.trim="urls"
    :rules="z.url('无效的 URL').or(z.literal(''))"
    :rows="urls.split('\n').length"
    placeholder="URL"
    variant="outlined"
    textarea
    clearable
  />

  <VarButton @click="parse" :disabled="!urls" type="primary">
    <VarIcon name="xml" />
    解析
  </VarButton>
</template>
