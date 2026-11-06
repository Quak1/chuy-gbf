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
  <UContainer>
    <form @submit.prevent="handleSubmit" class="mx-auto w-fit">
      <UFormField label="Username" size="xl" required class="my-2">
        <UInput v-model="username" type="text" />
      </UFormField>

      <p v-if="error" class="text-red-500 text-sm">Error:{{ error.message }}</p>
      <p v-else-if="data" class="text-yellow-500 text-sm">Data: {{ data }}</p>
      <p v-if="loading" class="text-blue-500 text-sm">Loading...</p>

      <UButton type="submit" size="xl" class="block ml-auto">Login</UButton>
    </form>
  </UContainer>
</template>
