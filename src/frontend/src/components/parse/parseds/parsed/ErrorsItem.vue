<script setup>
const { errors } = defineProps(['errors'])
</script>

<template>
  <div class="alerts">
    <VarAlert
      v-for="error in errors"
      v-bind:key="error"
      :message="error"
      type="danger"
      variant="tonal"
    />
  </div>
</template>

<style scoped>
div.alerts {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin: 1rem;
  margin-bottom: 0;
}
</style>
