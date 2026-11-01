import type { ItemsByCategory } from "../models";
import { useFetch } from "./useFetch";

const { data, error, loading } = useFetch<ItemsByCategory>(
  "/api/items/categorized",
);

export function useItems() {
  return { items: data, error, loading };
}
