<script setup lang="ts">
import { computed } from "vue";
import type { Item } from "../models";
import { getItemImageURL } from "../utils";

const props = defineProps<{
  item: Item;
  highRes?: boolean;
  portrait?: boolean;
  lazyLoading?: boolean;
  disabled?: boolean;
  party?: boolean;
}>();

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
      :loading="lazyLoading ? 'lazy' : undefined"
    />
  </a>
</template>

<style scoped>
.disabled {
  pointer-events: none;
}
img {
  width: 100%;
  height: auto;
  min-width: 50px;
  display: block;
}

.max150 {
  max-width: 150px;
}
</style>
