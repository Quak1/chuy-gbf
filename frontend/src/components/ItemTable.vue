<script setup lang="ts">
import { computed, ref } from "vue";
import { useFetch } from "../composables/useFetch";
import type { TableData } from "../models";

const { data, error, loading } = useFetch<TableData>("/api/users/items");
const elementFilter = ref("dark");

const elements = computed(() => {
  if (!data.value) return [];
  return [...new Set(data.value.items.map((item) => item.element))];
});

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
    <div>
      <button v-for="element in elements" @click="elementFilter = element">
        {{ element }}
      </button>
    </div>
    <div v-if="data">
      <table>
        <thead>
          <tr>
            <th>Username</th>
            <th v-for="item in filteredItems" :key="item.id">
              {{ item.name.toUpperCase() }}
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
