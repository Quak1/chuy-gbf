import { ref } from "vue";
import type { User } from "../models";
import { router } from "../router";

const userKey = "user";
const user = ref<User | null>(null);

const login = (username: string, id: number) => {
  user.value = { username, id };
  localStorage.setItem(userKey, JSON.stringify(user.value));
};

const logout = () => {
  user.value = null;
  localStorage.removeItem(userKey);
};

export function useUser(redirect: boolean = true) {
  const stored = localStorage.getItem(userKey);

  if (stored != null) {
    try {
      const val = JSON.parse(stored);
      const id = parseInt(val.id);
      const username = String(val.username);
      user.value = { username, id };
    } catch (e) {
      console.error("Error parsing localStorage");
    }
  } else if (redirect) {
    router.push({ name: "Login" });
  }

  return { user, login, logout };
}
