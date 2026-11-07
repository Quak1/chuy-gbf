<script setup lang="ts">
import { computed, ref, watch } from "vue";
import type { SearchItemModalData } from "../models";
import { useItems } from "../composables/useItems";
import ItemImage from "./ItemImage.vue";

const props = defineProps<{
  data: SearchItemModalData;
}>();

const emit = defineEmits<{
  set: [itemId: string];
}>();

const { items } = useItems();

const open = defineModel<boolean>("open");

const page = ref(1);
const itemsPerPage = 10;
const search = ref("");
const elementFilter = ref("");

const categoryItems = computed(() => {
  if (!items.value) return [];
  return Object.values(items.value[props.data.category]);
});

const filtered = computed(() => {
  return categoryItems.value.filter((item) => {
    const matchSearch = item.name.includes(search.value);
    const matchElement =
      elementFilter.value === "" || item.element === elementFilter.value;
    return matchSearch && matchElement;
  });
});

watch(search, () => (page.value = 1));

const paginated = computed(() => {
  const start = (page.value - 1) * itemsPerPage;
  return filtered.value.slice(start, start + itemsPerPage);
});

function handleElementUpdate(element: string) {
  page.value = 1;
  elementFilter.value = element;
}
</script>

<template>
  <UModal v-model:open="open" :title="`Edit ${data.category}`">
    <template #body>
      <div class="flex gap-2 items-center justify-center">
        <UInput
          type="text"
          v-model="search"
          placeholder="Sarch..."
          icon="i-lucide-search"
        />
        <UButton
          @click="emit('set', 'clear')"
          color="error"
          icon="i-lucide-trash"
          >Remove item</UButton
        >
      </div>

      <ElementSelector :items="categoryItems" @update="handleElementUpdate" />

      <div class="flex flex-col gap-1 my-2">
        <div
          v-for="item in paginated"
          @click="emit('set', item.id)"
          class="flex items-center gap-2 bg-muted rounded-sm capitalize cursor-pointer min-h-[80px]"
          :key="item.id"
        >
          <ItemImage :item="item" disabled />
          <div>{{ item.name }}</div>
        </div>
      </div>

      <UPagination
        v-model:page="page"
        :total="filtered.length"
        :itemsPerPage="itemsPerPage"
        :ui="{ list: 'justify-center' }"
      />
    </template>
  </UModal>
</template>
