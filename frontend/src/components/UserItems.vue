<script setup lang="ts">
import { computed, ref } from "vue";
import EditItemModal from "./EditItemModal.vue";
import type { UserItem } from "../models";
import { useFetch } from "../composables/useFetch";
import { useRoute } from "vue-router";

const isModalOpen = ref(false);
const selectedItem = ref<UserItem | null>(null);
const elementFilter = ref("dark");

const route = useRoute();
const { data, loading } = useFetch<UserItem[]>(
  `/api/users/${route.params.userID}/items`,
);

function openEditModal(item: UserItem) {
  selectedItem.value = item;
  isModalOpen.value = true;
}

function closeModal() {
  isModalOpen.value = false;
  selectedItem.value = null;
}

function handleItemUpdate(updatedItem: UserItem) {
  if (!data.value) return;

  const index = data.value.findIndex((i) => i.id === updatedItem.id);
  if (index !== -1) data.value[index] = updatedItem;

  closeModal();
}

const elements = computed(() => {
  if (!data.value) return [];
  return [...new Set(data.value.map((item) => item.element))];
});
const filteredItems = computed(() => {
  if (!data.value) return [];
  return data.value.filter((item) => item.element == elementFilter.value);
});
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-if="data">
    <div>
      <button v-for="e in elements" @click="elementFilter = e">{{ e }}</button>
    </div>

    <div v-for="item in filteredItems" :key="item.id">
      <span>{{ item.name }}</span>
      <span>Value: {{ item.value }}</span>
      <button @click="openEditModal(item)">Edit</button>
    </div>

    <EditItemModal
      v-if="isModalOpen && selectedItem"
      :item="selectedItem"
      @close="closeModal"
      @updated="handleItemUpdate"
    />
  </div>
</template>
