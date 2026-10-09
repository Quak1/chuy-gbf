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
dialog {
  padding: 20px;
  width: 400px;
  background-color: var(--accent-light);
  border: 10px solid var(--border);
  border-radius: 25px;
  color: var(--bg);
}

h3 {
  margin: 5px 0;
  font-size: 25px;
  text-transform: capitalize;
}

label {
  margin: 10px 0;
  display: flex;
  align-items: center;
  gap: 5px;
  font-size: 18px;
}

input {
  width: 100%;
  font-size: 18px;
  background-color: var(--accent-light);
  color: var(--bg);
  border: 2px solid var(--bg);
  border-radius: 10px;
}

.buttons {
  display: flex;
  justify-content: end;
  gap: 10px;
}

button {
  font-size: 16px;
  padding: 10px;
  background-color: var(--bg);
  border: none;
  border-radius: 10px;
  color: var(--text);
  font-weight: bold;
}
button:hover {
  filter: brightness(1.3);
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
