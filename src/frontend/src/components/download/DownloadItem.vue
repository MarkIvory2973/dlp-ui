<script setup>
const { url, download } = defineProps(['url', 'download'])
</script>

<template>
  <VarCard
    :title="url"
    :subtitle="
      !download.progress.done
        ? download.title == 'downloading'
          ? '正在下载'
          : `正在下载: ${download.title}`
        : '下载终了'
    "
  >
    <template v-if="download.errors" #description>
      <VarSpace class="alert" direction="column" size="large">
        <VarAlert
          v-for="error in download.errors"
          v-bind:key="error"
          :message="error"
          type="danger"
        />
      </VarSpace>
    </template>
    <template v-if="!download.progress.done" #extra>
      <VarProgress :value="(download.progress.current / download.progress.total) * 100" label>
        {{ ((download.progress.current / download.progress.total) * 100).toFixed(0) }}%
        {{ (download.progress.speed / 1024 / 1024).toFixed(2) }} MB/s
      </VarProgress>
    </template>
  </VarCard>
</template>

<style scoped>
.alerts {
  margin: 1rem;
  margin-bottom: 0;
}
</style>
