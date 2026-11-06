<script setup lang="ts">
import { computed, h, ref } from "vue";
import { useFetch } from "../composables/useFetch";
import type { ItemValue, TableData, User } from "../models";
import ElementSelector from "../components/ElementSelector.vue";
import ItemImage from "../components/ItemImage.vue";
import HoverTooltip from "../components/HoverTooltip.vue";
import type { TableColumn, TableRow } from "@nuxt/ui";

const { data, error, loading } = useFetch<TableData>("/api/users/items");
const elementFilter = ref("");

const users = computed(() => {
  if (!data.value) return [];

  return Object.values(data.value.users).map((user) => ({
    user: { id: user.id, username: user.username, comment: user.comment },
    ...user.values,
  }));
});

type DynamicRow = Record<string, unknown>;
const columns = computed<TableColumn<DynamicRow>[]>(() => {
  if (!data.value) return [];

  const filtered = data.value.items.filter((item) => {
    const matchElement =
      elementFilter.value === "" || elementFilter.value === item.element;
    return matchElement;
  });

  const columns = filtered.map((item) => {
    const column: TableColumn<DynamicRow> = {
      accessorKey: item.id,
      header: () => h(ItemImage, { item: item }, ""),
      cell: ({ row }) => {
        const value: ItemValue | undefined = row.getValue(item.id);
        return value
          ? h(
              "div",
              {
                class:
                  "flex items-center justify-center rounded-full p-1 font-semibold text-lg leading-none",
                style: { backgroundColor: value.color, color: "white" },
              },
              value.value,
            )
          : "";
      },
      meta: {
        style: { th: { width: "100px", minWidth: "100px", maxWidth: "100px" } },
      },
    };
    return column;
  });

  columns.unshift({
    accessorKey: "user",
    header: "",
    meta: { style: { th: { minWidth: "75px", maxWidth: "fit-content" } } },
  });

  return columns;
});

const getUser = (row: TableRow<DynamicRow>) => row.original.user as User;
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

    <UTable
      :data="users"
      :columns="columns"
      :columnPinning="{ left: ['user'] }"
      @select="() => {}"
      :ui="{ th: 'p-0', td: 'p-0', base: 'min-w-0' }"
    >
      <template #user-cell="{ row }">
        <div class="p-1 flex justify-between items-center gap-1">
          <RouterLink :to="`/users/${getUser(row).id}/items`">
            {{ getUser(row).username }}
          </RouterLink>
          <HoverTooltip
            v-if="getUser(row).comment"
            :text="getUser(row).comment"
          />
        </div>
      </template>
    </UTable>
  </div>
</template>
