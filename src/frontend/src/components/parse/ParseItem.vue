<script setup>
// vue
import { ref } from 'vue'
import { storeToRefs } from 'pinia'
// zod
import z from 'zod'
// local
import { useEnvStore } from '@/stores/env'

// varlet
import { Snackbar } from '@varlet/ui'

const { baseUrl } = storeToRefs(useEnvStore())

const form = ref(null)
const url = ref('')

async function parse() {
  url.value.split('\n').forEach(async (url, index) => {
    url = url.trim(' ')
    if (!url) {
      return
    }

    const response = await fetch(`${baseUrl}/api/parse`, {
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

  form.value.reset()
}
</script>

<template>
  <VarSpace direction="column" size="large">
    <VarInput
      ref="form"
      v-model.trim="url"
      :rules="z.url('无效的 URL').or(z.literal(''))"
      placeholder="URL"
      variant="outlined"
      :rows="url.split('\n').length"
      textarea
      clearable
    />

    <VarButton @click="parse" :disabled="!url" type="primary">
      <VarIcon name="xml" />
      解析
    </VarButton>
  </VarSpace>
</template>

<style scoped>
button {
  width: 100%;
}
</style>
