<script setup>
const { job } = defineProps(['job'])
</script>

<template>
  <VarProgress :value="(job.current / job.total) * 100" label>
    {{ ((job.current / job.total) * 100).toFixed(0) || 0 }} %
    {{ (job.speed / 1048576).toFixed(2) }} M/s
  </VarProgress>
</template>
