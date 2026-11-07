<script setup lang="ts">
import type { ItemValue } from "../models";
import { useAPI } from "../composables/useAPI";

const props = defineProps<{
  value: ItemValue;
}>();

const emit = defineEmits<{
  delete: [id: ItemValue];
}>();

const { call, data, error, loading } = useAPI();

const API_BASE = `/api/items/${props.value.item_id}/values`;

const handleDelete = async () => {
  const ok = await call(`${API_BASE}/${props.value.id}`, {
    method: "DELETE",
  });
  if (ok) emit("delete", props.value);
  else {
    console.error("error", error.value);
    console.error("data", data.value);
  }
};
</script>

<template>
  <div class="flex items-center bg-muted p-1 rounded-sm">
    <span
      class="text-ink text-lg font-bold py-2 px-3 rounded-sm mr-1"
      :style="{ backgroundColor: value.color }"
      >{{ value.value }}</span
    >
    <UButton
      @click="handleDelete"
      color="error"
      icon="i-lucide-trash"
      :loading="loading"
      >Delete</UButton
    >
  </div>
</template>
