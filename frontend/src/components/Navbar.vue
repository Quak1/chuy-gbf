<script setup lang="ts">
import { computed } from "vue";
import { useUser } from "../composables/useUser";
import type { NavigationMenuItem } from "@nuxt/ui";

const { user, logout } = useUser(false);

const navItems = computed(() => {
  const items: NavigationMenuItem[] = [
    {
      label: "My Items",
      to: `/users/${user.value?.id}/items`,
      show: !!user.value && user.value.role !== "admin",
    },
    {
      label: "Teams",
      to: "/teams",
      show: !!user.value,
    },
    {
      label: "Login",
      to: "/login",
      show: !user.value,
    },
    {
      label: "Items Admin",
      to: "/items",
      show: user.value?.role === "admin",
    },
  ];

  return items.filter((item) => item.show);
});
</script>

<template>
  <UHeader title="GBF">
    <UNavigationMenu :items="navItems" />

    <template #right>
      <UButton
        v-if="user"
        @click="logout"
        color="error"
        variant="soft"
        label="Logout"
      />
    </template>
  </UHeader>
</template>
