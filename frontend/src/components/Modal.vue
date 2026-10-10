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

<style>
dialog {
  padding: 20px;
  width: 400px;
  background-color: var(--accent-light);
  border: 5px solid var(--border);
  border-radius: 25px;
  color: var(--bg);

  h3 {
    margin: 5px 0;
    font-size: 25px;
    text-transform: capitalize;
  }

  button {
    font-size: inherit;
    padding: 10px;
    background-color: var(--bg);
    border: none;
    border-radius: 10px;
    color: var(--text);
    font-weight: bold;
  }
  button:hover {
    filter: brightness(1.3);
  }
}
</style>
