<script setup lang="ts">
import { watch } from "vue";
import { useFetch } from "../composables/useFetch";
import type { Item, ItemValue } from "../models";
import Modal from "./Modal.vue";
import ValuePicker from "./ValuePicker.vue";

const props = defineProps<{
  item: Item;
}>();

const emit = defineEmits(["close", "updated"]);

const { data, loading, error } = useFetch<ItemValue[]>(
  `/api/items/${props.item.id}/values`,
);

let tempId = -1;

const addEmptyValue = () => {
  if (!data.value) data.value = [];

  const lastItem = data.value.at(-1);
  if (lastItem && lastItem.id < 0) return;

  data.value.push({
    id: tempId--,
    item_id: props.item.id,
    value: "0",
    color: "#FFFFFF",
  });
};

watch(
  data,
  (newData) => {
    if (newData) addEmptyValue();
  },
  { immediate: true },
);

const handleDelete = (value: ItemValue) => {
  if (!data.value) return;
  data.value = data.value.filter((v) => v.id !== value.id);
};

const handleCreate = (value: ItemValue) => {
  if (!data.value) return;

  const lastIndex = data.value.length - 1;
  if (data.value[lastIndex].id < 0) {
    data.value[lastIndex] = value;
    addEmptyValue();
  }

  console.log(data.value);
};
</script>

<template>
  <Modal @close="emit('close')">
    <div>
      <h3>Edit values {{ item.name }}</h3>

      <ValuePicker
        v-for="value in data"
        :value="value"
        :key="value.id"
        @delete="handleDelete"
        @create="handleCreate"
      />

      <div class="buttons">
        <button @click="emit('close')" :disabled="loading">Close</button>
      </div>

      <p v-if="loading">Loading...</p>
      <p v-if="error">{{ error }}</p>
    </div>
  </Modal>
</template>

<style>
dialog {
  background-color: var(--accent);
  width: fit-content;
}
.buttons {
  display: flex;
  align-items: center;
  justify-content: end;
  gap: 5px;
  margin-top: 10px;
}
</style>
