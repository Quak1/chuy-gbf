export interface TableData {
  items: Item[];
  users: Record<number, UserItemValues>;
}

export interface UserItemValues {
  id: number;
  username: string;
  values: Record<string, string>;
}

export interface Item {
  id: string;
  name: string;
  element: string;
  type: string;
  series: string;
  enabled: boolean;
}

export interface User {
  id: number;
  username: string;
  role: string;
}
