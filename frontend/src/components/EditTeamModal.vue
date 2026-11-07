<script setup lang="ts">
import { ref, watchEffect } from "vue";
import type { Team } from "../models";
import { usePost } from "../composables/usePost";

const props = defineProps<{
  team: Team;
}>();

const emit = defineEmits<{
  update: [name: string, url: string];
}>();

const open = ref(false);
const { post, data, error } = usePost();
const name = ref();
const url = ref();

watchEffect(() => {
  name.value = props.team.name;
  url.value = props.team.url;
});

const handleSubmit = async () => {
  const ok = await post(`/api/teams/${props.team.id}`, {
    name: name.value,
    url: url.value,
  });

  if (ok) emit("update", name.value, url.value);
  else {
    console.log({ error: error.value, data: data.value });
    alert("There was an error updating the team. Try again later!");
  }

  open.value = false;
};
</script>

<template>
  <UModal v-model:open="open" title="Edit Team">
    <UButton @click="open = true">Edit</UButton>

    <template #body>
      <form @submit.prevent="handleSubmit">
        <UFormField label="Name" size="xl" required>
          <UInput v-model="name" type="text" class="w-full" />
        </UFormField>

        <UFormField label="Source" size="xl" class="mb-3">
          <UInput v-model="url" type="text" class="w-full" />
        </UFormField>

        <div class="flex justify-between">
          <UButton type="button" @click="open = false" color="neutral"
            >Cancel</UButton
          >
          <UButton type="submit">Save</UButton>
        </div>
      </form>
    </template>
  </UModal>
</template>
