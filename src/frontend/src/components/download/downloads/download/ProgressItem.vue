<script setup>
const { task } = defineProps(['task'])
</script>

<template>
  <VarProgress :value="(task.current / task.total) * 100" label>
    {{ ((task.current / task.total) * 100).toFixed(0) || 0 }} %
    {{ (task.speed / 1048576).toFixed(2) }} M/s
  </VarProgress>
</template>
