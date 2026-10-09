<script setup lang="ts">
import { ref } from "vue";
import type { ItemValue } from "../models";
import { useAPI } from "../composables/useAPI";
import { usePost } from "../composables/usePost";

const props = defineProps<{
  value: ItemValue;
}>();

const emit = defineEmits<{
  delete: [id: ItemValue];
  create: [value: ItemValue];
}>();

const defaultColors = ["#11734b", "#d4edbc", "#ffe6a0", "#ffcfc9", "#b10202"];

const itemValue = ref(props.value);
const draftColor = ref(itemValue.value.color);
const draftValue = ref(itemValue.value.value);

const { call, data, error, loading } = useAPI();
const {
  post,
  data: postData,
  error: postError,
  loading: postLoading,
} = usePost<ItemValue>();

const API_BASE = `/api/items/${itemValue.value.item_id}/values`;
const handleDelete = async () => {
  const ok = await call(`${API_BASE}/${itemValue.value.id}`, {
    method: "DELETE",
  });
  if (ok) emit("delete", itemValue.value);
  else {
    console.error("error", error.value);
    console.error("data", data.value);
  }
};

const handleCreate = async () => {
  const ok = await post(`${API_BASE}`, {
    value: draftValue.value,
    color: draftColor.value,
  });
  if (ok && postData.value) {
    itemValue.value = postData.value;
    emit("create", itemValue.value);
  }
};
</script>

<template>
  <div :inert="loading || postLoading">
    <span :inert="!!itemValue.id">
      <input type="text" v-model="draftValue" />
      <input type="color" v-model="draftColor" />

      <button
        class="defaultColor"
        v-for="color in defaultColors"
        :style="{ backgroundColor: color }"
        @click="draftColor = color"
      ></button>
    </span>

    <button v-if="itemValue.id" @click="handleDelete">Delete</button>
    <button v-if="!itemValue.id" @click="handleCreate">Save</button>

    <p>{{ error }}</p>
    <p>{{ postError }}</p>
  </div>
</template>

<style scoped>
div {
  border: 1px solid var(--border);
  border-radius: 10px;
}
.defaultColor {
  margin: 3px;
  padding: 0;
  width: 25px;
  height: 25px;
  border: none;
  border-radius: 20px;
}
</style>
