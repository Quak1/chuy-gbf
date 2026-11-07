<script setup lang="ts">
import { computed, ref, useTemplateRef } from "vue";
import { useFetch } from "../composables/useFetch";
import type { DBUpdate, Item } from "../models";
import { useUser } from "../composables/useUser";
import { usePost } from "../composables/usePost";
import { useRouter } from "vue-router";
import AdminItemSetupModal from "../components/AdminItemSetupModal.vue";
import { getPaginationRowModel, type Row } from "@tanstack/vue-table";
import ItemImage from "../components/ItemImage.vue";
import type { DropdownMenuItem, TableColumn } from "@nuxt/ui";

type UpdateDBResponse = {
  message: string;
  errors?: string[];
};

const table = useTemplateRef("table");
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
const globalFilter = ref("");
const selectedItem = ref<Item>();

const editItemModalOpen = ref(false);

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
    const matchElement =
      !elementFilter.value || item.element === elementFilter.value;
    const matchType = !typeFilter.value || item.type === typeFilter.value;
    const matchSeries =
      !seriesFilter.value || item.series === seriesFilter.value;
    return matchElement && matchType && matchSeries;
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

const columns: TableColumn<Item>[] = [
  { accessorKey: "id" },
  { accessorKey: "name" },
  { accessorKey: "element" },
  { accessorKey: "type" },
  { accessorKey: "series" },
  { id: "actions", meta: { class: { td: "text-right" } } },
];

const pagination = ref({
  pageIndex: 0,
  pageSize: 30,
});

function getRowActions(row: Row<Item>): DropdownMenuItem[] {
  return [
    {
      label: `${row.original.enabled ? "Disable" : "Enable"} item`,
      icon: `${row.original.enabled ? "i-lucide-eye-closed" : "i-lucide-eye"}`,
      onSelect: () => toggleEnabled(row.original),
    },
    {
      label: "Edit values",
      icon: "i-lucide-pencil",
      onSelect: () => {
        selectedItem.value = row.original;
        editItemModalOpen.value = true;
      },
    },
  ];
}
</script>

<template>
  <AdminItemSetupModal
    v-model:open="editItemModalOpen"
    v-if="selectedItem"
    :item="selectedItem"
  />

  <div class="container">
    <div v-if="loading">Loading</div>
    <div v-if="error">{{ error }}</div>

    <div class="flex justify-between py-3 border-b border-accented">
      <UInput v-model="globalFilter" class="max-w-sm" placeholder="Filter..." />

      <div>
        <span v-if="updateFetchData" class="text-sm text-muted"
          >Last update: {{ updateFetchData.created_at }}
        </span>
        | <a href="/api/database/backup">Download DB backup</a> |
        <UButton @click="handlePostUpdateData(false)" variant="outline"
          >Update DB Data</UButton
        >
      </div>
    </div>

    <UTable
      ref="table"
      :data="filtered"
      :columns="columns"
      v-model:global-filter="globalFilter"
      v-model:pagination="pagination"
      :pagination-options="{ getPaginationRowModel: getPaginationRowModel() }"
      @select="() => {}"
      :ui="{ td: 'p-0' }"
      :meta="{
        class: { tr: (row) => (row.original.enabled ? 'bg-primary/10' : '') },
      }"
    >
      <template #id-cell="{ row }">
        <div class="w-[100px] text-center">
          <ItemImage :item="row.original" />
          <p>{{ row.original.id }}</p>
        </div>
      </template>

      <template #actions-cell="{ row }">
        <UDropdownMenu :items="getRowActions(row)" :content="{ align: 'end' }">
          <UButton
            icon="i-lucide-ellipsis-vertical"
            color="neutral"
            variant="ghost"
          />
        </UDropdownMenu>
      </template>
    </UTable>

    <div class="flex justify-end border-t border-default py-4">
      <UPagination
        :page="(table?.tableApi?.getState().pagination.pageIndex || 0) + 1"
        :items-per-page="table?.tableApi?.getState().pagination.pageSize"
        :total="table?.tableApi?.getFilteredRowModel().rows.length"
        @update:page="(p: number) => table?.tableApi?.setPageIndex(p - 1)"
      />
    </div>
  </div>
</template>
