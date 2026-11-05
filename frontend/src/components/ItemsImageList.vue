<script setup lang="ts">
import type { ItemCategory, SlotItem } from "../models";
import ItemImage from "./ItemImage.vue";

const props = defineProps<{
  items: SlotItem[];
  category: ItemCategory;
  lazyLoading?: boolean;
}>();
</script>

<template>
  <div>
    <template v-for="item in items">
      <ItemImage v-if="item" :item square :height="50" />
    </template>
  </div>
</template>

<style scoped>
div {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
  align-items: center;
}

.empty {
  background-color: black;
}
</style>
