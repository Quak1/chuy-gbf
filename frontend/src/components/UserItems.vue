<script setup lang="ts">
import { computed, ref } from "vue";
import EditItemModal from "./EditItemModal.vue";
import type { UserItem, User } from "../models";
import { useFetch } from "../composables/useFetch";
import { useRoute } from "vue-router";
import ElementSelector from "./ElementSelector.vue";
import ItemImage from "./ItemImage.vue";
import { useModal } from "../composables/useModal";
import AddCommentModal from "./AddCommentModal.vue";
import HoverTooltip from "./HoverTooltip.vue";
import CommentIcon from "./icons/CommentIcon.vue";

const elementFilter = ref("dark");
const { selectedData, openModal, closeModal } = useModal<UserItem>();
const {
  selectedData: modalComment,
  openModal: openCommentModal,
  closeModal: closeCommentModal,
} = useModal<string>();

const route = useRoute();
const userID = computed(() => {
  const param = route.params.userID;
  return Array.isArray(param) ? param[0] : param;
});

const { data, loading } = useFetch<UserItem[]>(
  `/api/users/${userID.value}/items`,
);
const { data: userData } = useFetch<User>(`/api/users/${userID.value}`);

function handleItemUpdate(updatedItem: UserItem) {
  if (!data.value) return;

  const index = data.value.findIndex((i) => i.id === updatedItem.id);
  if (index !== -1) data.value[index] = updatedItem;

  closeModal();
}

function handleComment(comment: string) {
  closeCommentModal();
  if (!userData.value) return;
  userData.value.comment = comment;
}

const filteredItems = computed(() => {
  if (!data.value) return [];
  return data.value.filter((item) => item.element == elementFilter.value);
});
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-if="data">
    <ElementSelector
      v-if="data"
      :items="data"
      @update="(e) => (elementFilter = e)"
    />

    <div v-if="userData" class="comment-container">
      <button @click="openCommentModal('commenting')">
        {{ userData.comment ? "Edit" : "Add" }} user comment
      </button>
      <HoverTooltip v-if="userData.comment" class="tooltop">
        <template #outer>
          <CommentIcon />
        </template>
        {{ userData.comment }}
      </HoverTooltip>
    </div>

    <div class="container">
      <div v-for="item in filteredItems" :key="item.id">
        <ItemImage :item="item" />
        <h3>{{ item.name }}</h3>
        <p :style="{ backgroundColor: item.color }">
          {{ item.value || "no value" }}
        </p>
        <button @click="openModal(item)">Edit</button>
      </div>
    </div>

    <EditItemModal
      v-if="selectedData"
      :item="selectedData"
      @close="closeModal"
      @update="handleItemUpdate"
    />

    <AddCommentModal
      v-if="modalComment"
      :userID="userID"
      @close="closeCommentModal"
      @update="handleComment"
    />
  </div>
</template>

<style scoped>
.comment-container {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
  margin: 10px;

  .tooltop {
    width: 30px;
    height: 30px;
  }
}
button {
  font-size: inherit;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.container > div {
  width: 200px;
  padding: 10px;
  border: 2px solid var(--border);
  border-radius: 10px;
  display: flex;
  flex-direction: column;
  gap: 8px;

  img {
    align-self: center;
  }

  h3 {
    margin: 0;
    text-transform: capitalize;
  }

  p {
    color: var(--black-clear);
    background-color: rgba(255, 255, 255, 0.2);
    font-size: 20px;
    text-align: center;
    border-radius: 20px;
    padding: 5px 0;
    margin-top: auto;
  }
}
</style>
