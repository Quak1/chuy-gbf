<script setup lang="ts">
import { ref, watch } from "vue";
import type { User } from "../models";
import { useUser } from "../composables/useUser";
import { usePost } from "../composables/usePost";
import { router } from "../router";

const username = ref("");
const { login } = useUser(false);
const { post, data, error, loading } = usePost<User>("/api/users");

const handleSubmit = async () => {
  post({ username: username.value });
};

watch(data, () => {
  if (data.value) {
    login(data.value.username, data.value.id);
    router.push({ name: "Home" });
  }
});
</script>

<template>
  <form @submit.prevent="handleSubmit">
    <label>
      Username:
      <input type="text" v-model="username" required />
      <p v-if="error">Error: {{ error.message }}</p>
      <p v-else-if="data">Data: {{ data }}</p>
    </label>
    <button type="submit">Login</button>
    <p v-if="loading">Loading...</p>
  </form>
</template>
