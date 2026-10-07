<script setup lang="ts">
import { onMounted, ref } from "vue";
import type { UserItem } from "../models";
import { usePost } from "../composables/usePost";
import { useUser } from "../composables/useUser";

const props = defineProps<{
  item: UserItem;
}>();

const emit = defineEmits(["close", "updated"]);

const dialogRef = ref<HTMLDialogElement | null>(null);
const draftValue = ref(props.item.value);
const { user } = useUser(false);
const { post, error, loading } = usePost();

onMounted(() => {
  dialogRef.value?.showModal();
});

function onCancel(e: Event) {
  e.preventDefault();
  emit("close");
}

async function submitValue() {
  if (draftValue.value === props.item.value) {
    emit("close");
    return;
  }

  const url = `/api/users/${user.value?.id}/items/${props.item.id}`;
  const ok = await post(url, {
    value: draftValue.value,
    username: user.value?.username,
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
  <dialog ref="dialogRef" @cancel="onCancel">
    <div>
      <h3>Edit {{ item.name }}</h3>

      <label
        >Value
        <input
          v-model="draftValue"
          :disabled="loading"
          @keyup.enter="submitValue"
          autofocus
        />
      </label>

      <div>
        <button @click="emit('close')" :disabled="loading">Cancel</button>
        <button @click="submitValue" :disabled="loading">
          {{ loading ? "Saving..." : "Save" }}
        </button>
      </div>

      <p v-if="error">{{ error }}</p>
    </div>
  </dialog>
</template>
