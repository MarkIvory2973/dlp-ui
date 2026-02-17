<script setup>
const { parsed } = defineProps(['parsed'])
</script>

<template>
  <VarCell :description="`解析结果 (${parsed.entries.length})`" />

  <VarCell
    v-for="entry in parsed.entries"
    v-bind:key="entry"
    :title="entry.title || '无解析结果'"
    :description="entry.creators?.join(' / ') || entry.uploader || '不支持该 URL 的作者显示'"
    ripple
  />
</template>
