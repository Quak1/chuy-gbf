<script setup lang="ts">
import { computed, ref } from "vue";
import EditItemModal from "./EditItemModal.vue";
import type { UserItem } from "../models";
import { useFetch } from "../composables/useFetch";
import { useRoute } from "vue-router";
import ElementSelector from "./ElementSelector.vue";
import ItemImage from "./ItemImage.vue";

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

const filteredItems = computed(() => {
  if (!data.value) return [];
  return data.value.filter((item) => item.element == elementFilter.value);
});
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-if="data">
    <ElementSelector
      v-if="data"
      :items="data"
      @update="(e) => (elementFilter = e)"
    />

    <div class="container">
      <div v-for="item in filteredItems" :key="item.id">
        <ItemImage :item="item" />
        <h3>{{ item.name }}</h3>
        <p>Value: {{ item.value }}</p>
        <button @click="openEditModal(item)">Edit</button>
      </div>
    </div>

    <EditItemModal
      v-if="isModalOpen && selectedItem"
      :item="selectedItem"
      @close="closeModal"
      @updated="handleItemUpdate"
    />
  </div>
</template>

<style scoped>
.container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.container > div {
  width: 200px;
  padding: 10px;
  border: 2px solid var(--border);
  border-radius: 10px;
  display: flex;
  flex-direction: column;
  gap: 10px;

  img {
    align-self: center;
  }

  h3 {
    margin: 0;
    text-transform: capitalize;
  }

  button {
    margin-top: auto;
  }
}
</style>
