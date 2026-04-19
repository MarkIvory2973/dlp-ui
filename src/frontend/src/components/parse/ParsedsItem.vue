<script setup>
// vue
import { storeToRefs } from 'pinia'
// local
import { useParsedsStore } from '@/stores/parse/parseds'

// local
import ParsedItem from './parseds/ParsedItem.vue'

const { parseds } = storeToRefs(useParsedsStore())
</script>

<template>
  <ParsedItem
    v-for="parsed in parseds"
    v-bind:key="parsed"
    :url="parsed.url"
    :task="parsed.task"
    :errors="parsed.errors"
  />
</template>
