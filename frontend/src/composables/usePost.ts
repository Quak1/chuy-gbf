import { useAPI } from "./useAPI";

export function usePost<T>() {
  const { call, data, error, loading } = useAPI<T>();

  const post = async (url: string, payload: any) => {
    const options = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    };

    return call(url, options);
  };

  return { post, data, error, loading };
}
