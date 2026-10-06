import { ref, watchEffect, toValue } from "vue";

export function useFetch<T>(url: string) {
  const data = ref<T | null>(null);
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const fetchData = () => {
    data.value = null;
    error.value = null;
    loading.value = true;

    fetch(toValue(url))
      .then((res) => res.json())
      .then((json) => (data.value = json))
      .catch((err) => (error.value = err))
      .finally(() => (loading.value = false));
  };

  watchEffect(() => {
    fetchData();
  });

  return { data, error, loading };
}
