<script setup lang="ts">
import { ref } from "vue";
import type { Team } from "../models";
import Modal from "./Modal.vue";
import { usePost } from "../composables/usePost";

const props = defineProps<{
  data: Team;
}>();
const emit = defineEmits<{
  close: [];
  update: [name: string, url: string];
}>();

const { post, data, error } = usePost();
const name = ref(props.data.name);
const url = ref(props.data.url);

const handleSubmit = async () => {
  const ok = await post(`/api/teams/${props.data.id}`, {
    name: name.value,
    url: url.value,
  });

  if (ok) emit("update", name.value, url.value);
  else {
    console.log({ error: error.value, data: data.value });
    alert("There was an error updating the team. Try again later!");
  }

  emit("close");
};
</script>

<template>
  <Modal @close="emit('close')">
    <form @submit.prevent="handleSubmit">
      <h2>Edit Team:</h2>
      <div>
        <label for="name">Name: </label>
        <input id="name" type="text" v-model="name" />
      </div>
      <div>
        <label for="url">Source: </label>
        <input id="url" type="text" v-model="url" />
      </div>
      <div class="buttons">
        <button type="button" @click="emit('close')">Cancel</button>
        <button type="submit">Save</button>
      </div>
    </form>
  </Modal>
</template>

<style scoped>
h2 {
  margin: 5px 0;
}
input {
  padding: 5px;
}
form > div {
  display: flex;
  align-items: center;
  gap: 5px;
}
.buttons {
  justify-content: space-between;
}
</style>
