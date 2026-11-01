<script setup lang="ts">
import { computed, ref } from "vue";
import { useFetch } from "../composables/useFetch";
import type { DBUpdate, Item } from "../models";
import { useUser } from "../composables/useUser";
import { usePost } from "../composables/usePost";
import { useRouter } from "vue-router";
import { useModal } from "../composables/useModal";
import AdminItemSetupModal from "../components/AdminItemSetupModal.vue";
import ItemImage from "../components/ItemImage.vue";

type UpdateDBResponse = {
  message: string;
  errors?: string[];
};

const router = useRouter();
const { user } = useUser();
const { data, loading, error } = useFetch<Item[]>("/api/items");
const { data: updateFetchData } = useFetch<DBUpdate>("/api/update-data/last");
const { post, data: postData, error: postError } = usePost();
const {
  post: postUpdate,
  data: dataUpdate,
  error: errorUpdate,
} = usePost<UpdateDBResponse>();
const elementFilter = ref("");
const typeFilter = ref("");
const seriesFilter = ref("");
const searchFilter = ref("");
const { selectedData, openModal, closeModal } = useModal<Item>();

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
  const ok = await post(url, {});
  if (ok) {
    item.enabled = !item.enabled;
  } else {
    alert("Failed to update DB");
    console.error(postData.value);
    console.error(postError.value);
  }
};

const handlePostUpdateData = async (forceUpdate = false) => {
  const update = confirm(
    `Are you sure you want to${forceUpdate ? " FORCE" : ""} update DB data?`,
  );
  if (!update) return;

  const payload = forceUpdate ? { force: true } : {};
  const ok = await postUpdate("/api/update-data", payload);
  if (!ok) {
    console.error(errorUpdate.value);
    alert("There was a server error updating DB data.");
  } else if (dataUpdate.value?.errors) {
    console.error(dataUpdate.value.errors);
    alert("Check logs for DB error messages.");
  } else if (dataUpdate.value?.message) {
    const force = confirm(`${dataUpdate.value.message}. Try force update?`);
    if (force) handlePostUpdateData(true);
  } else {
    alert("Successful database data update.");
  }
};
</script>

<template>
  <AdminItemSetupModal
    v-if="selectedData"
    :item="selectedData"
    @close="closeModal"
  />

  <div class="container">
    <div v-if="loading">Loading</div>
    <div v-if="error">{{ error }}</div>

    <div class="top">
      <label
        >Search
        <input type="text" v-model="searchFilter" />
      </label>

      <div>
        <span v-if="updateFetchData"
          >Last update: {{ updateFetchData.created_at }} |
        </span>
        <button @click="handlePostUpdateData(false)">Update DB Data</button>
        | <a href="/api/database/backup">Download DB backup</a>
      </div>
    </div>

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
          <td>Setup</td>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in filtered" :key="item.id">
          <td>
            <ItemImage :item="item" lazyLoading />
            {{ item.id }}
          </td>
          <td>
            {{ item.name }}
          </td>
          <td>{{ item.element }}</td>
          <td>{{ item.type }}</td>
          <td>{{ item.series }}</td>
          <td>
            <button @click="toggleEnabled(item)">
              {{ item.enabled ? "Disable" : "Enable" }}
            </button>
            <button @click="openModal(item)">Edit values</button>
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
table {
  width: 100%;
}
td {
  word-break: break-word;
  overflow-wrap: break-word;
}
tr td:first-child {
  width: 100px;
  text-align: center;
}

thead button {
  display: block;
}

.top {
  display: flex;
  justify-content: space-between;
}
</style>
