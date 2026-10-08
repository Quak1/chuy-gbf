<script setup lang="ts">
import { computed, ref } from "vue";
import { useFetch } from "../composables/useFetch";
import type { Item } from "../models";
import { useUser } from "../composables/useUser";
import { usePost } from "../composables/usePost";
import { useRouter } from "vue-router";

const router = useRouter();
const { user } = useUser();
const { data, loading, error } = useFetch<Item[]>("/api/items");
const { post, data: postData, error: postError } = usePost();
const elementFilter = ref("");
const typeFilter = ref("");
const seriesFilter = ref("");
const searchFilter = ref("");

if (user.value?.role !== "admin") {
  router.push({ name: "Home" });
}

const elements = computed(() => {
  if (!data.value) return [];
  return [...new Set(data.value.map((item) => item.element))];
});

const types = computed(() => {
  if (!data.value) return [];
  return [...new Set(data.value.map((item) => item.type))];
});

const series = computed(() => {
  if (!data.value) return [];
  return [...new Set(data.value.map((item) => item.series))];
});

const filtered = computed(() => {
  if (!data.value) return [];

  return data.value.filter((item) => {
    const matchSearch =
      !searchFilter.value ||
      item.name.includes(searchFilter.value) ||
      item.id.includes(searchFilter.value);
    const matchElement =
      !elementFilter.value || item.element === elementFilter.value;
    const matchType = !typeFilter.value || item.type === typeFilter.value;
    const matchSeries =
      !seriesFilter.value || item.series === seriesFilter.value;
    return matchSearch && matchElement && matchType && matchSeries;
  });
});

const toggleEnabled = async (item: Item) => {
  const url = `/api/items/${item.id}/${item.enabled ? "disable" : "enable"}`;
  const ok = await post(url, { username: user.value?.username });
  if (ok) {
    item.enabled = !item.enabled;
  } else {
    alert("Failed to update DB");
    console.error(postData.value);
    console.error(postError.value);
  }
};
</script>

<template>
  <div class="container">
    <div v-if="loading">Loading</div>
    <div v-if="error">{{ error }}</div>
    <label
      >Search
      <input type="text" v-model="searchFilter" />
    </label>
    <table v-if="filtered">
      <thead>
        <tr>
          <td>ID</td>
          <td>Name</td>
          <td>
            <button @click="elementFilter = ''">Element</button>
            <select v-model="elementFilter">
              <option v-for="element in elements">{{ element }}</option>
            </select>
          </td>
          <td>
            <button @click="typeFilter = ''">Type</button>
            <select v-model="typeFilter">
              <option v-for="type in types">{{ type }}</option>
            </select>
          </td>
          <td>
            <button @click="seriesFilter = ''">Series</button>
            <select v-model="seriesFilter">
              <option v-for="s in series">{{ s }}</option>
            </select>
          </td>
          <td>Enabled</td>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in filtered" :key="item.id">
          <td>{{ item.id }}</td>
          <td>{{ item.name }}</td>
          <td>{{ item.element }}</td>
          <td>{{ item.type }}</td>
          <td>{{ item.series }}</td>
          <td>
            <button @click="toggleEnabled(item)">
              {{ item.enabled ? "Disable" : "Enable" }}
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.container {
  max-width: 1200px;
  margin: 0 auto;
}
td {
  word-break: break-word;
  overflow-wrap: break-word;
}

thead button {
  display: block;
}
</style>
