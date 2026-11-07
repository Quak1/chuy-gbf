<script setup lang="ts">
import { ref } from "vue";
import type { Item, ItemValue } from "../models";
import { usePost } from "../composables/usePost";

const props = defineProps<{
  item: Item;
}>();

const emit = defineEmits<{
  create: [value: ItemValue];
}>();

const defaultColors = ["#11734b", "#d4edbc", "#ffe6a0", "#ffcfc9", "#b10202"];

const draftColor = ref("#FFFFFF");
const draftValue = ref("0");

const {
  post,
  data: postData,
  error: postError,
  loading: postLoading,
} = usePost<ItemValue>();

const API_BASE = `/api/items/${props.item.id}/values`;

const handleCreate = async () => {
  const ok = await post(`${API_BASE}`, {
    value: draftValue.value,
    color: draftColor.value,
  });

  if (ok && postData.value) {
    emit("create", postData.value);

    draftColor.value = "#FFFFFF";
    draftValue.value = "0";
  }
};
</script>

<template>
  <div>
    <div class="flex items-center justify-center gap-5">
      <div>
        <UInput
          type="text"
          v-model="draftValue"
          autofocus
          :style="{ backgroundColor: draftColor }"
          :ui="{ base: 'text-ink font-bold text-center' }"
        />
        <div class="flex justify-center gap-1 mt-1">
          <UButton
            v-for="color in defaultColors"
            @click="draftColor = color"
            class="size-[30px]"
            :style="{ backgroundColor: color }"
          ></UButton>
        </div>
      </div>

      <UColorPicker type="color" v-model="draftColor" />
    </div>

    <p class="text-error">{{ postError }}</p>
    <UButton @click="handleCreate" :loading="postLoading">Save</UButton>
  </div>
</template>
