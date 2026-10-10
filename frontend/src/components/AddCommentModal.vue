<script setup lang="ts">
import { ref } from "vue";
import Modal from "./Modal.vue";
import { usePost } from "../composables/usePost";

const props = defineProps<{
  userID: string;
}>();

const emit = defineEmits<{
  close: [];
  update: [comment: string];
}>();

const commentDraft = ref("");
const { post, error, loading } = usePost();

const handleSubmit = async () => {
  const ok = await post(`/api/users/${props.userID}/comments`, {
    comment: commentDraft.value,
  });
  if (ok) emit("update", commentDraft.value);
};
</script>

<template>
  <Modal @close="$emit('close')">
    <form @submit.prevent="handleSubmit">
      <fieldset :disabled="loading">
        <h3>Add user comment</h3>

        <textarea
          name="comment"
          id="comment"
          v-model="commentDraft"
          rows="5"
        ></textarea>

        <p v-if="error">{{ error }}</p>

        <div class="buttons">
          <button type="submit">Submit</button>
          <button @click="$emit('close')">Cancel</button>
        </div>
      </fieldset>
    </form>
  </Modal>
</template>

<style scoped>
fieldset {
  border: none;
}
textarea {
  width: 100%;
  box-sizing: border-box;
  margin-bottom: 10px;
  font-size: inherit;
  resize: vertical;
}
.buttons {
  display: flex;
  justify-content: end;
  gap: 10px;
}
</style>
