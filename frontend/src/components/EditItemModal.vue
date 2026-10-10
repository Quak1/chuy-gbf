<script setup lang="ts">
import { ref } from "vue";
import type { ItemValue, UserItem } from "../models";
import { usePost } from "../composables/usePost";
import { useUser } from "../composables/useUser";
import Modal from "./Modal.vue";
import { useFetch } from "../composables/useFetch";

const props = defineProps<{
  item: UserItem;
}>();

const emit = defineEmits<{
  close: [];
  update: [item: UserItem];
}>();

const valueChoice = ref<ItemValue | null>(null);
const { user } = useUser(false);
const { post, error, loading } = usePost();
const {
  data,
  error: errorFetch,
  loading: loadingFetch,
} = useFetch<ItemValue[]>(`/api/items/${props.item.id}/values`);

async function submitValue(value: ItemValue) {
  if (!value) {
    emit("close");
    return;
  }

  const url = `/api/users/${user.value?.id}/items/${props.item.id}`;
  const ok = await post(url, {
    valueID: value.id,
  });

  if (ok && valueChoice) {
    emit("update", {
      ...props.item,
      color: value.color,
      value: value.value,
    });
  }
}
</script>

<template>
  <Modal @close="$emit('close')">
    <div v-if="loading">Loading...</div>
    <div v-if="!loadingFetch">
      <h3>Edit {{ item.name }}</h3>

      <div class="values">
        <button
          v-for="value in data"
          :key="value.id"
          :style="{ backgroundColor: value.color }"
          @click="submitValue(value)"
        >
          {{ value.value }}
        </button>
      </div>

      <div class="buttons">
        <button @click="emit('close')" :disabled="loading">Cancel</button>
      </div>

      <p v-if="error">POST | {{ error }}</p>
      <p v-if="errorFetch">GET | {{ errorFetch }}</p>
    </div>
  </Modal>
</template>

<style scoped>
.buttons {
  display: flex;
  justify-content: end;
  gap: 10px;
}

.values {
  display: flex;
  gap: 10px;

  button {
    font-size: 20px;
    color: var(--black-clear);
    border: 2px solid var(--black-clear);
  }
}
</style>
