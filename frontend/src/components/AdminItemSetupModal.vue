<script setup lang="ts">
import { computed, watchEffect } from "vue";
import { useFetch } from "../composables/useFetch";
import type { Item, ItemValue } from "../models";
import ValuePicker from "./ValuePicker.vue";

const props = defineProps<{
  item: Item;
}>();

const open = defineModel<boolean>("open");

const apiUrl = computed(() => `/api/items/${props.item.id}/values`);
const { data, loading, error, fetchData } = useFetch<ItemValue[]>(apiUrl.value);

watchEffect(() => fetchData(apiUrl.value));

const handleDelete = (value: ItemValue) => {
  if (!data.value) return;
  data.value = data.value.filter((v) => v.id !== value.id);
};

const handleCreate = (value: ItemValue) => {
  if (!data.value) return;

  const lastIndex = data.value.length - 1;
  if (data.value[lastIndex].id < 0) {
    data.value[lastIndex] = value;
  }
};
</script>

<template>
  <UModal v-model:open="open" :title="`Edit ${item.name} values`">
    <template #body>
      <div class="flex flex-wrap gap-2" v-if="data">
        <ValueDisplay
          v-for="value in data"
          :value="value"
          @delete="handleDelete"
        />
      </div>

      <USeparator class="my-3" />

      <ValuePicker v-if="item" :item="item" @create="handleCreate" />

      <p v-if="loading" class="text-info">Loading...</p>
      <p v-if="error" class="text-error">{{ error }}</p>
    </template>
  </UModal>
</template>
