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

export interface Item {
  id: string;
  name: string;
  element: string;
  type: string;
  series: string;
  enabled: boolean;
  category: string;
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
