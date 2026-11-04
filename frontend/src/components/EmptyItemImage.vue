<script setup lang="ts">
import { computed } from "vue";
import { getEmptyImagePath } from "../utils";

const props = defineProps<{
  category: string;
  portrait?: boolean;
  lazyLoading?: boolean;
}>();

const src = computed(() => getEmptyImagePath(props));
</script>

<template>
  <img :src alt="empty" :loading="lazyLoading ? 'lazy' : undefined" />
</template>

<style scoped>
img {
  width: 100%;
  max-height: 100%;
  max-width: 100%;
  min-width: 50px;
  display: block;
}
</style>
