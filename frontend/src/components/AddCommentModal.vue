<script setup lang="ts">
import { ref } from "vue";
import { usePost } from "../composables/usePost";

const props = defineProps<{
  userID: string;
  comment: string;
}>();

const emit = defineEmits<{
  update: [comment: string];
}>();

const open = ref(false);
const commentDraft = ref(props.comment);
const { post, error, loading } = usePost();

const handleSubmit = async () => {
  const ok = await post(`/api/users/${props.userID}/comments`, {
    comment: commentDraft.value,
  });

  if (ok) {
    emit("update", commentDraft.value);
    open.value = false;
  }
};
</script>

<template>
  <UModal
    v-model:open="open"
    title="Edit user comment"
    :ui="{ footer: 'justify-end' }"
  >
    <UButton>Edit user comment</UButton>

    <template #body>
      <UForm @submit.prevent="handleSubmit" id="comment-form">
        <UTextarea
          v-model="commentDraft"
          name="comment"
          id="comment"
          class="w-full"
          :rows="8"
          :loading="loading"
          :disabled="loading"
          autofocus
        ></UTextarea>

        <p v-if="error" class="text-error">{{ error }}</p>
      </UForm>
    </template>

    <template #footer>
      <UButton @click="open = false" color="neutral" :disabled="loading"
        >Cancel</UButton
      >
      <UButton type="submit" form="comment-form" :disabled="loading"
        >Submit</UButton
      >
    </template>
  </UModal>
</template>

<style scoped></style>
