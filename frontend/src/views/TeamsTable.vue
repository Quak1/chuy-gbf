<script setup lang="ts">
import { RouterLink, useRouter } from "vue-router";
import { useFetch } from "../composables/useFetch";
import { usePost } from "../composables/usePost";
import type { Team } from "../models";
import {
  columnFilteringFeature,
  createColumnHelper,
  createFilteredRowModel,
  filterFn_includesString,
  FlexRender,
  tableFeatures,
  useTable,
} from "@tanstack/vue-table";
import { computed, h } from "vue";
import ItemsImageList from "../components/ItemsImageList.vue";

const router = useRouter();
const { data, error, loading } = useFetch<Team[]>("/api/teams");
const {
  post,
  data: postData,
  error: postError,
} = usePost<{ team_id: number }>();

const handleCreateBtn = async () => {
  const ok = await post("/api/teams", {});
  if (ok) {
    router.push(`/teams/builder/${postData.value?.team_id}`);
  } else {
    console.error(postError.value);
  }
};

const features = tableFeatures({
  columnFilteringFeature,
  filteredRowModel: createFilteredRowModel(),
  filterFns: {
    includesString: filterFn_includesString,
  },
});

const columnHelper = createColumnHelper<typeof features, Team>();
const columns = columnHelper.columns([
  columnHelper.accessor("name", {
    header: "Name",
    cell: (p) =>
      h(
        RouterLink,
        { to: `/teams/builder/${p.row.original.id}` },
        p.getValue(),
      ),
  }),
  columnHelper.accessor("character", {
    header: "Character Images",
    cell: (props) =>
      h(ItemsImageList, { items: props.getValue(), category: "character" }),
  }),
  columnHelper.accessor("summon", {
    header: "Summon Images",
    cell: (props) =>
      h(ItemsImageList, { items: props.getValue(), category: "summon" }),
  }),
  columnHelper.accessor("weapon", {
    header: "Weapon Images",
    cell: (props) =>
      h(ItemsImageList, { items: props.getValue(), category: "weapon" }),
  }),
]);

const tableData = computed(() => data.value ?? []);
const table = useTable({ features, columns, data: tableData });
</script>

<template>
  <button @click="handleCreateBtn">Create new team</button>

  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error }}</div>

  <table v-else class="data-table">
    <thead>
      <tr v-for="headerGroup in table.getHeaderGroups()" :key="headerGroup.id">
        <th v-for="header in headerGroup.headers" :key="header.id">
          <FlexRender v-if="!header.isPlaceholder" :header="header" />
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-if="table.getRowModel().rows.length === 0">
        <td :colspan="columns.length">No records found.</td>
      </tr>
      <tr v-else v-for="row in table.getRowModel().rows" :key="row.id">
        <td v-for="cell in row.getAllCells()" :key="cell.id">
          <FlexRender :cell="cell" />
        </td>
      </tr>
    </tbody>
  </table>
</template>

<style scoped></style>
