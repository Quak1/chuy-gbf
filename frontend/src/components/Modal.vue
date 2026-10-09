<script setup lang="ts">
import { onMounted, ref } from "vue";

const emit = defineEmits(["close"]);

const dialogRef = ref<HTMLDialogElement | null>(null);

onMounted(() => {
  dialogRef.value?.showModal();
});

function onCancel(e: Event) {
  e.preventDefault();
  emit("close");
}
</script>

<template>
  <dialog ref="dialogRef" @cancel="onCancel">
    <slot />
  </dialog>
</template>
