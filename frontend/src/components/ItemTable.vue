<script setup lang="ts">
import { computed, ref } from "vue";
import { useFetch } from "../composables/useFetch";
import type { TableData } from "../models";
import ElementSelector from "./ElementSelector.vue";
import ItemImage from "./ItemImage.vue";
import CommentIcon from "./icons/CommentIcon.vue";
import HoverTooltip from "./HoverTooltip.vue";

const { data, error, loading } = useFetch<TableData>("/api/users/items");
const elementFilter = ref("dark");
const searchString = ref("");

const filteredItems = computed(() => {
  if (!data.value) return [];

  return data.value.items.filter((item) => {
    const matchElement = elementFilter.value === item.element;
    const matchItemName = item.name.includes(searchString.value);
    return matchElement && matchItemName;
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

    <div class="search-container">
      <label for="search">Item search: </label>
      <input
        type="text"
        id="search"
        v-model="searchString"
        autocomplete="off"
      />
    </div>

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
            <td>
              <div>
                <RouterLink :to="`/users/${user.id}/items`"
                  >{{ user.username }}
                </RouterLink>
                <HoverTooltip v-if="user.comment" class="tooltip">
                  <template #outer>
                    <CommentIcon />
                  </template>
                  {{ user.comment }}
                </HoverTooltip>
              </div>
            </td>
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
.search-container {
  margin: 10px 0;

  input {
    padding: 5px;
    background-color: var(--accent-light);
    color: var(--bg);
    border: 2px solid var(--bg);
    border-radius: 30px;
  }
}

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
  min-width: 100px;
}
tr th:not(:first-child) {
  width: 100px;
  min-width: 100px;
}

tr td:first-child {
  > div {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  padding: 5px;
  a {
    text-decoration: none;
    color: inherit;
  }
  a:hover {
    filter: brightness(1.3);
  }
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

.tooltip {
  margin-left: 10px;
  height: 20px;
  width: 20px;
}
</style>
