<script setup lang="ts">
import { ref } from "vue";
import type { UserItem } from "../models";
import { usePost } from "../composables/usePost";
import { useUser } from "../composables/useUser";
import Modal from "./Modal.vue";

const props = defineProps<{
  item: UserItem;
}>();

const emit = defineEmits(["close", "updated"]);

const draftValue = ref(props.item.value);
const { user } = useUser(false);
const { post, error, loading } = usePost();

async function submitValue() {
  if (draftValue.value === props.item.value) {
    emit("close");
    return;
  }

  const url = `/api/users/${user.value?.id}/items/${props.item.id}`;
  const ok = await post(url, {
    value: draftValue.value,
  });

  if (ok) {
    emit("updated", {
      ...props.item,
      value: draftValue.value,
    });
  }
}
</script>

<template>
  <Modal @close="$emit('close')">
    <div>
      <h3>Edit {{ item.name }}</h3>

      <label
        >Value:
        <input
          v-model="draftValue"
          :disabled="loading"
          @keyup.enter="submitValue"
          autofocus
        />
      </label>

      <div class="buttons">
        <button @click="emit('close')" :disabled="loading">Cancel</button>
        <button @click="submitValue" :disabled="loading">
          {{ loading ? "Saving..." : "Save" }}
        </button>
      </div>

      <p v-if="error">{{ error }}</p>
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
</style>
