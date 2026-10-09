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
              <div :style="{ backgroundColor: user.values[item.id]?.color }">
                {{ user.values[item.id]?.value }}
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.table-container {
  overflow-x: auto;
}

table {
  table-layout: fixed;
  border-collapse: collapse;
  user-select: none;
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
  border: 1px solid var(--black-clear);
}

tr th:first-child {
  width: 120px;
  min-width: 120px;
}
tr th:not(:first-child) {
  width: 100px;
  min-width: 100px;
}

tr td:first-child {
  padding: 5px;
}

tr td:not(:first-child) {
  padding: 0px;
  height: 100%;
  height: 28px;

  div {
    display: flex;
    align-items: center;
    justify-content: center;

    height: 100%;
    padding: 5px;
    margin: 1px;
    box-sizing: border-box;
    border-radius: 50px;

    font-weight: bold;
    line-height: 0em;
    color: var(--black-clear);
  }
}
</style>
