import { ref } from "vue";

export function useModal<T>() {
  const selectedData = ref<T | null>(null);

  function openModal(data: T) {
    selectedData.value = data;
  }

  function closeModal() {
    selectedData.value = null;
  }

  return { selectedData, openModal, closeModal };
}
