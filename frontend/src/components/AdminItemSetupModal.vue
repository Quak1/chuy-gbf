<script setup lang="ts">
import { ref } from "vue";
import { useFetch } from "../composables/useFetch";
import type { Item, ItemValue } from "../models";
import Modal from "./Modal.vue";
import ValuePicker from "./ValuePicker.vue";

const props = defineProps<{
  item: Item;
}>();

const emit = defineEmits(["close", "updated"]);

const isCreating = ref(false);
const { data, loading, error } = useFetch<ItemValue[]>(
  `/api/items/${props.item.id}/values`,
);

const handleAdd = () => {
  isCreating.value = true;
  if (!data.value) data.value = [];

  data.value.push({
    id: 0,
    item_id: props.item.id,
    value: "0",
    color: "#000000",
  });
};

const handleDelete = (value: ItemValue) => {
  data.value = data.value?.filter((v) => v.id !== value.id) || [];
};

const handleCreate = (value: ItemValue) => {
  isCreating.value = false;
  if (!data.value) return;
  const index = data.value.findIndex((v) => v.id === 0);
  if (index !== -1) data.value[index] = value;
};
</script>

<template>
  <Modal @close="$emit('close')">
    <div>
      <h3>Edit values {{ item.name }}</h3>

      <ValuePicker
        v-if="data"
        v-for="value in data"
        :value="value"
        :key="value.id"
        @delete="handleDelete"
        @create="handleCreate"
      />

      <div class="buttons">
        <button @click="handleAdd" :disabled="isCreating">Add Value</button>
        <button @click="emit('close')" :disabled="loading">Close</button>
      </div>

      <p v-if="loading">Loading...</p>
      <p v-if="error">{{ error }}</p>
    </div>
  </Modal>
</template>
