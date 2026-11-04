<script setup lang="ts">
import { computed } from "vue";
import type { ItemCategory, SlotItem } from "../models";
import { getItemImageURL } from "../utils";

const props = defineProps<{
  items: SlotItem[];
  category: ItemCategory;
  lazyLoading?: boolean;
}>();

const srcs = computed(() =>
  props.items
    .filter((item) => item !== null)
    .map((item) => ({
      url: getItemImageURL({
        item,
        highRes: false,
        portrait: false,
        party: false,
        square: true,
      }),
      name: item.name,
    })),
);
</script>

<template>
  <div>
    <img
      v-for="src in srcs"
      :src="src.url"
      :alt="src.name"
      :loading="lazyLoading ? 'lazy' : undefined"
    />
  </div>
</template>

<style scoped>
div {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
}

img {
  height: 75px;
  width: auto;
  display: block;
}
</style>
