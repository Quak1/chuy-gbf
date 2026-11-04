<script setup lang="ts">
import { computed, ref, watch } from "vue";
import type { SearchItemModalData } from "../models";
import Modal from "./Modal.vue";
import { useItems } from "../composables/useItems";
import ItemImage from "./ItemImage.vue";

const props = defineProps<{
  data: SearchItemModalData;
}>();
const emit = defineEmits<{
  close: [];
  set: [itemId: string];
}>();

const { items } = useItems();

const page = ref(0);
const search = ref("");
const itemsPerPage = ref(7);

const filtered = computed(() => {
  if (!items.value) {
    return [];
  }

  const entries = items.value[props.data.category];
  return Object.keys(entries)
    .filter((key) => entries[key].name.includes(search.value))
    .map((key) => entries[key]);
});

watch(search, () => (page.value = 0));

const totalPages = computed(() =>
  Math.ceil(filtered.value.length / itemsPerPage.value),
);

const pages = computed(() => {
  const sidePages = 3;
  const maxLength = sidePages * 2 + 1;

  if (totalPages.value <= maxLength)
    return Array.from({ length: totalPages.value }, (_, i) => i);

  let start = page.value - 3;
  if (start < 0) {
    start = 0;
  }
  if (start + maxLength > totalPages.value) {
    start = totalPages.value - maxLength;
  }

  return Array.from({ length: maxLength }, (_, i) => i + start);
});

const paginated = computed(() => {
  const start = page.value * itemsPerPage.value;
  return filtered.value.slice(start, start + itemsPerPage.value);
});
</script>

<template>
  <Modal @close="emit('close')">
    <div class="container">
      <h3>Edit {{ data.category }}</h3>
      <input type="text" v-model="search" />
      <button @click="emit('set', 'clear')">Remove item</button>

      <div class="item-container">
        <div
          v-for="item in paginated"
          @click="emit('set', item.id)"
          class="item"
          :key="item.id"
        >
          <ItemImage :item="item" disabled />
          <div>{{ item.name }}</div>
        </div>
      </div>

      <div class="paginator">
        <div v-if="pages[0] != 0" class="side">
          <button @click="page = 0" :class="{ active: page === 0 }">1</button>
          <span>...</span>
        </div>

        <button
          v-for="n in pages"
          @click="page = n"
          :class="{ active: page === n }"
          :key="n"
        >
          {{ n + 1 }}
        </button>

        <div v-if="pages[pages.length - 1] != totalPages - 1" class="side">
          <span>...</span>
          <button
            @click="page = totalPages - 1"
            :class="{ active: page === totalPages - 1 }"
          >
            {{ totalPages }}
          </button>
        </div>
      </div>
    </div>
  </Modal>
</template>

<style scoped>
dialog {
  background-color: var(--accent);
  width: fit-content;
}
.container {
  max-height: 70vh;
  text-align: center;
  text-wrap: nowrap;
}

input {
  padding: 1px 5px;
  font-size: 1.1rem;
}

.item-container {
  width: 100%;

  .item {
    margin: 5px 0;
    min-height: 80px;
    display: flex;
    align-items: center;
    gap: 10px;
    padding-right: 10px;

    font-size: 1.1rem;
    background-color: var(--accent-light);
    border-radius: 10px;
    text-transform: capitalize;
    cursor: pointer;
    overflow: hidden;
  }

  a {
    flex-shrink: 0;
  }
}
.paginator {
  display: flex;
  justify-content: center;
  gap: 2px;

  button {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 40px;
    height: 40px;
  }

  .active {
    border: 2px solid var(--accent-light);
  }

  .side {
    display: flex;
    align-items: center;
    span {
      font-weight: bold;
      margin: 0 5px;
    }
  }
}
</style>
