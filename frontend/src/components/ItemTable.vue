<script setup lang="ts">
import { computed, ref } from "vue";
import { useFetch } from "../composables/useFetch";
import type { TableData } from "../models";
import ElementSelector from "./ElementSelector.vue";
import ItemImage from "./ItemImage.vue";

const { data, error, loading } = useFetch<TableData>("/api/users/items");
const elementFilter = ref("dark");

const filteredItems = computed(() => {
  return data.value?.items.filter((item) => {
    return item.element === elementFilter.value;
  });
});
</script>

<template>
  <div>
    <div v-if="loading">Loading...</div>
    <div v-if="error">{{ error }}</div>
    <ElementSelector
      v-if="data"
      :items="data.items"
      @update="(e) => (elementFilter = e)"
    />

    <div v-if="data" class="table-container">
      <table>
        <thead>
          <tr>
            <th></th>
            <th v-for="item in filteredItems" :key="item.id">
              <ItemImage :item="item" />
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in data.users" :key="user.id">
            <td>{{ user.username }}</td>
            <td v-for="item in filteredItems" :key="item.id">
              {{ user.values[item.id] }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.table-container {
  width: 100%;
  overflow-x: auto;
}

th {
  padding: 0;
}

tr th:first-child,
tr td:first-child {
  position: sticky;
  left: 0;
  z-index: 1;
  background-color: var(--border);
}
</style>
