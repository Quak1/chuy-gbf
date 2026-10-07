<script setup lang="ts">
import type { Item } from "../models";

const props = defineProps<{
  items: Item[];
}>();

const emit = defineEmits<{
  update: [element: string];
}>();

const elements = [...new Set(props.items?.map((item) => item.element))];
</script>

<template>
  <div>
    <button v-for="e in elements" @click="$emit('update', e)">
      {{ e }}
    </button>
  </div>
</template>

<style scoped>
div {
  display: flex;
  justify-content: center;
  gap: 10px;
  margin: 10px;
}

button {
  padding: 10px;
  font-size: 16px;
  background-color: var(--border);
  border-radius: 10px;
  border: none;
  text-transform: capitalize;
}
button:hover {
  filter: brightness(0.8);
}
</style>
