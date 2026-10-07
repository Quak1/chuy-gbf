<script setup lang="ts">
import { computed } from "vue";
import { useUser } from "./composables/useUser";

const { user } = useUser(false);
const itemsURL = computed(() => {
  return `/users/${user.value?.id}/items`;
});
</script>

<template>
  <nav>
    <RouterLink to="/">Home</RouterLink>
    <RouterLink to="/hello">Hello</RouterLink>
    <RouterLink to="/login">Login</RouterLink>
    <RouterLink :to="itemsURL" v-if="user">Items</RouterLink>
  </nav>
  <RouterView />
</template>
