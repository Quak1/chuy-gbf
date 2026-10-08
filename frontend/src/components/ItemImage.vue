<script setup lang="ts">
import { computed } from "vue";
import type { Item } from "../models";
import { getItemImageURL } from "../utils";

const props = defineProps<{
  item: Item;
  highRes?: boolean;
}>();

const src = computed(() => getItemImageURL(props.item, props.highRes));
</script>

<template>
  <img :src :alt="item.name" :class="{ max150: !highRes }" />
</template>

<style scoped>
img {
  width: 100%;
  height: auto;
  min-width: 150px;
  display: block;
}

.max150 {
  max-width: 150px;
}
</style>
