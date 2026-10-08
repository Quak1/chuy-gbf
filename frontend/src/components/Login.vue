<script setup lang="ts">
import { ref, watch } from "vue";
import type { User } from "../models";
import { useUser } from "../composables/useUser";
import { usePost } from "../composables/usePost";
import { useRouter } from "vue-router";

const router = useRouter();
const username = ref("");
const { login } = useUser(false);
const { post, data, error, loading } = usePost<User>();

const handleSubmit = async () => {
  post("/api/users", { username: username.value });
};

watch(data, () => {
  if (data.value) {
    login(data.value);
    router.push({ name: "Home" });
  }
});
</script>

<template>
  <form @submit.prevent="handleSubmit">
    <label for="username"> Username: </label>
    <input type="text" id="username" v-model="username" required />
    <p v-if="error">Error: {{ error.message }}</p>
    <p v-else-if="data">Data: {{ data }}</p>
    <button type="submit">Login</button>
    <p v-if="loading">Loading...</p>
  </form>
</template>

<style scoped>
form {
  width: 300px;
  display: flex;
  flex-direction: column;
  margin: 0 auto;
}

label {
  display: block;
  font-size: 24px;
  font-weight: bold;
  color: var(--text);
}

input {
  max-width: 100%;
  padding: 10px 14px;
  border: 1px solid var(--border);
  border-radius: 10px;
  font-size: 16px;
  outline: none;
  margin: 10px 0;
}

input:focus {
  border-color: var(--accent-light);
}

button {
  width: 100%;
  font-size: 16px;
  padding: 5px 0;
}
</style>
