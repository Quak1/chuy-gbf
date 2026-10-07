import { ref } from "vue";
import type { User } from "../models";
import { router } from "../router";

const userKey = "user";
const user = ref<User | null>(null);

const login = (u: User) => {
  user.value = u;
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
      user.value = { id, username: val.username, role: val.role };
    } catch (e) {
      console.error("Error parsing localStorage");
    }
  } else if (redirect) {
    router.push({ name: "Login" });
  }

  return { user, login, logout };
}
