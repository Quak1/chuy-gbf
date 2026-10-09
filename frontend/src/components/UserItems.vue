<script setup lang="ts">
import { computed, ref } from "vue";
import EditItemModal from "./EditItemModal.vue";
import type { UserItem } from "../models";
import { useFetch } from "../composables/useFetch";
import { useRoute } from "vue-router";
import ElementSelector from "./ElementSelector.vue";
import ItemImage from "./ItemImage.vue";
import { useModal } from "../composables/useModal";

const elementFilter = ref("dark");
const { selectedData, openModal, closeModal } = useModal<UserItem>();

const route = useRoute();
const { data, loading } = useFetch<UserItem[]>(
  `/api/users/${route.params.userID}/items`,
);

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
        <p :style="{ backgroundColor: item.color }">
          {{ item.value || "no value" }}
        </p>
        <button @click="openModal(item)">Edit</button>
      </div>
    </div>

    <EditItemModal
      v-if="selectedData"
      :item="selectedData"
      @close="closeModal"
      @update="handleItemUpdate"
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
  gap: 8px;

  img {
    align-self: center;
  }

  h3 {
    margin: 0;
    text-transform: capitalize;
  }

  p {
    color: var(--black-clear);
    background-color: rgba(255, 255, 255, 0.2);
    font-size: 20px;
    text-align: center;
    border-radius: 20px;
    padding: 5px 0;
    margin-top: auto;
  }
}
</style>
