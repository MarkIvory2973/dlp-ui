<script setup>
const floating = defineModel()
</script>

<template>
  <VarButton @click="floating = !floating" round text>
    <VarIcon name="view" />
  </VarButton>
</template>
