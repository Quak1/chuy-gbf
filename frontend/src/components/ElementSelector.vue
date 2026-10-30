<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import type { Item } from "../models";

const props = defineProps<{
  items: Item[];
}>();

const emit = defineEmits<{
  update: [element: string];
}>();

const activeElement = ref("");

const elements = computed(() => [
  ...new Set(props.items?.map((item) => item.element)),
]);
onMounted(() => {
  const defaultElement = "dark";
  const element = elements.value.includes(defaultElement)
    ? defaultElement
    : elements.value[0];

  emit("update", element);
  activeElement.value = element;
});

const handleClick = (element: string) => {
  emit("update", element);
  activeElement.value = element;
};
</script>

<template>
  <div>
    <button
      v-for="e in elements"
      @click="handleClick(e)"
      :class="{ active: activeElement === e }"
    >
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

.active {
  border: 2px dashed var(--accent-light);
}
</style>
