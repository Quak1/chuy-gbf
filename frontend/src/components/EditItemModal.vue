<script setup lang="ts">
import { ref } from "vue";
import type { ItemValue, UserItem } from "../models";
import { usePost } from "../composables/usePost";
import { useUser } from "../composables/useUser";
import { useFetch } from "../composables/useFetch";

const props = defineProps<{
  item: UserItem;
}>();

const emit = defineEmits<{
  update: [item: UserItem];
}>();

const open = ref(false);
const valueChoice = ref<ItemValue | null>(null);
const { user } = useUser(false);
const { post, error, loading } = usePost();
const {
  data,
  error: errorFetch,
  loading: loadingFetch,
} = useFetch<ItemValue[]>(`/api/items/${props.item.id}/values`);

async function submitValue(value: ItemValue) {
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
    open.value = false;
  }
}
</script>

<template>
  <UModal :title="`Edit ${item.name}`" v-model:open="open">
    <UButton
      class="w-full flex justify-center mt-1 font-bold text-base text-ink"
      :style="{ backgroundColor: item.color }"
      >{{ item.value || " - " }}</UButton
    >

    <template #body>
      <div v-if="loading">Loading...</div>

      <div v-if="!loadingFetch">
        <div class="flex gap-2">
          <UButton
            v-for="value in data"
            :key="value.id"
            class="text-ink"
            size="xl"
            :style="{ backgroundColor: value.color }"
            @click="submitValue(value)"
          >
            {{ value.value }}
          </UButton>
        </div>

        <p v-if="error" class="text-error">POST | {{ error }}</p>
        <p v-if="errorFetch" class="text-error">GET | {{ errorFetch }}</p>
      </div>
    </template>

    <template #footer>
      <UButton @click="open = false" :disabled="loading" color="neutral"
        >Cancel</UButton
      >
    </template>
  </UModal>
</template>
