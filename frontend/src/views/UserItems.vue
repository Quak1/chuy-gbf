<script setup lang="ts">
import { computed, ref } from "vue";
import EditItemModal from "../components/EditItemModal.vue";
import type { UserItem, User } from "../models";
import { useFetch } from "../composables/useFetch";
import { useRoute, useRouter } from "vue-router";
import ElementSelector from "../components/ElementSelector.vue";
import ItemImage from "../components/ItemImage.vue";
import AddCommentModal from "../components/AddCommentModal.vue";
import HoverTooltip from "../components/HoverTooltip.vue";
import { useAPI } from "../composables/useAPI";
import { useUser } from "../composables/useUser";

const elementFilter = ref("dark");

const route = useRoute();
const userID = computed(() => {
  const param = route.params.userID;
  return Array.isArray(param) ? param[0] : param;
});

const { data, loading } = useFetch<UserItem[]>(
  `/api/users/${userID.value}/items`,
);
const { data: userData } = useFetch<User>(`/api/users/${userID.value}`);
const { call, error: apiError } = useAPI();
const { user } = useUser();
const router = useRouter();

function handleItemUpdate(updatedItem: UserItem) {
  if (!data.value) return;

  const index = data.value.findIndex((i) => i.id === updatedItem.id);
  if (index !== -1) data.value[index] = updatedItem;
}

function handleComment(comment: string) {
  if (!userData.value) return;
  userData.value.comment = comment;
}

const filteredItems = computed(() => {
  if (!data.value) return [];
  return data.value.filter(
    (item) => elementFilter.value === "" || item.element == elementFilter.value,
  );
});

const handleDeleteUser = async () => {
  const username = userData.value?.username;
  const id = userData.value?.id;
  if (!confirm(`Delete user '${username}' with id:${id}`)) {
    return;
  }

  const ok = await call(`/api/users/${id}`, { method: "DELETE" });
  if (!ok) alert(apiError);
  else router.push({ name: "Home" });
};
</script>

<template>
  <div v-if="loading" class="text-info">Loading...</div>
  <div v-if="data">
    <div
      v-if="userData && user?.role !== 'admin'"
      class="flex justify-center items-center gap-1 my-2"
    >
      <AddCommentModal
        :userID="userID"
        @update="handleComment"
        :comment="userData.comment"
      />
      <HoverTooltip v-if="userData.comment" :text="userData.comment" />
    </div>

    <ElementSelector
      v-if="data"
      :items="data"
      @update="(e) => (elementFilter = e)"
    />

    <div v-if="user?.role === 'admin'" class="center">
      <UButton @click="handleDeleteUser" color="error">Delete user</UButton>
    </div>

    <div class="flex flex-wrap gap-2">
      <UCard
        v-for="item in filteredItems"
        :key="item.id"
        :title="item.name"
        variant="subtle"
        :ui="{
          header: 'capitalize p-2 sm:p-2',
          body: 'p-1 sm:p-1 flex flex-col items-center',
          root: 'w-[200px]',
        }"
      >
        <ItemImage :item="item" />
        <EditItemModal :item="item" @update="handleItemUpdate" />
      </UCard>
    </div>
  </div>
</template>
