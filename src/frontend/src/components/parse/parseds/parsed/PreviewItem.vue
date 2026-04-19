<script setup>
const { task } = defineProps(['task'])
</script>

<template>
  <VarCell :description="`解析结果 (${task.entries?.length || 0})`" />

  <VarCell
    v-if="!task.entries?.length"
    :title="task.done ? '无解析结果' : '正在解析 URL'"
    description="不支持该 URL 的作者显示"
    ripple
  />

  <VarCell
    v-for="entry in task.entries"
    v-bind:key="entry"
    :title="entry.title"
    :description="entry.creators?.join(' / ') || entry.uploader || '不支持该 URL 的作者显示'"
    ripple
  />
</template>
