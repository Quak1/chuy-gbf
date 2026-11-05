<script setup lang="ts">
import { computed } from "vue";
import type { Item } from "../models";
import { getItemImageURL } from "../utils";

const {
  height = "auto",
  width = "100%",
  ...props
} = defineProps<{
  item: Item;
  highRes?: boolean;
  portrait?: boolean;
  lazyLoading?: boolean;
  disabled?: boolean;
  party?: boolean;
  square?: boolean;
  height?: number | string;
  width?: number | string;
}>();

const imageStyle = computed(() => ({
  height: typeof height === "number" ? `${height}px` : height,
  width: typeof width === "number" ? `${width}px` : width,
}));

const src = computed(() => getItemImageURL(props));
</script>

<template>
  <a
    :href="`https://gbf.wiki/${item.wiki_path}`"
    target="_blank"
    rel="noopener noreferrer"
    :class="{ disabled }"
  >
    <img
      :src
      :alt="item.name"
      :class="{ max150: !highRes }"
      :style="imageStyle"
      :loading="lazyLoading ? 'lazy' : undefined"
    />
  </a>
</template>

<style scoped>
.disabled {
  pointer-events: none;
}

img {
  min-width: 50px;
  display: block;
  max-width: 100%;
}

.max150 {
  max-width: 150px;
}
</style>
