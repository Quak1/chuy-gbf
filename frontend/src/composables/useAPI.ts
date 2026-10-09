import { ref, toValue } from "vue";

export function useAPI<T>() {
  const data = ref<T | null>(null);
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const call = async (url: string, options: any) => {
    data.value = null;
    error.value = null;
    loading.value = true;

    try {
      const res = await fetch(toValue(url), options);

      if (res.status !== 204) data.value = await res.json();

      if (!res.ok) {
        throw new Error(data.value.error);
      }
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      } else {
        console.error("An unexpected error occurred: ", err);
        error.value = new Error("Unexpected error occurred");
      }
    }

    loading.value = false;
    return error.value ? false : true;
  };

  return { call, data, error, loading };
}
