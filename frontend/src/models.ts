export interface TableData {
  items: Item[];
  users: Record<number, UserItemValues>;
}

export interface UserItemValues {
  id: number;
  username: string;
  comment: string;
  values: Record<string, ItemValue>;
}

export type ItemCategory = "character" | "summon" | "weapon";

export interface Item {
  id: string;
  name: string;
  element: string;
  type: string;
  series: string;
  enabled: boolean;
  category: ItemCategory;
  wiki_path: string;
}

export interface ItemsByCategory<T = Item> {
  character: Record<string, T>;
  summon: Record<string, T>;
  weapon: Record<string, T>;
}

export interface User {
  id: number;
  username: string;
  role: string;
  comment: string;
}

export interface UserItem extends Item {
  value: string;
  color: string;
  enabled: true;
}

export interface ItemValue {
  id: number;
  item_id: string;
  value: string;
  color: string;
}

export interface DBUpdate {
  id: number;
  is_ok: boolean;
  comments: string;
  created_at: Date;
}

export interface Slot extends Item {
  position: number;
}

export type SlotItem = Slot | null;

export interface Team {
  name: string;
  url: string;
  id: number;
  character: SlotItem[];
  summon: SlotItem[];
  weapon: SlotItem[];
}

export interface SearchItemModalData {
  category: ItemCategory;
  idx: number;
}
