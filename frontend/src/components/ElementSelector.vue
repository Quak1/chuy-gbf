<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import type { Item } from "../models";

const { items, defaultElement = "" } = defineProps<{
  items: Item[];
  defaultElement?: string;
}>();

const emit = defineEmits<{
  update: [element: string];
}>();

const activeElement = ref("");

const elements = computed(() => [
  "clear",
  ...new Set(items?.map((item) => item.element)),
]);

onMounted(() => {
  emit("update", defaultElement);
  activeElement.value = defaultElement;
});

const handleClick = (element: string) => {
  if (element === "clear") element = "";
  emit("update", element);
  activeElement.value = element;
};
</script>

<template>
  <div class="flex justify-center gap-1 m-2">
    <UButton
      v-for="e in elements"
      @click="handleClick(e)"
      :variant="activeElement === e ? 'solid' : 'outline'"
    >
      {{ e[0].toUpperCase() + e.slice(1) }}
    </UButton>
  </div>
</template>
