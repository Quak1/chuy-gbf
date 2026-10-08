import { createRouter, createWebHistory } from "vue-router";
import ItemTable from "./components/ItemTable.vue";
import NotFound from "./components/NotFound.vue";
import Login from "./components/Login.vue";
import UserItems from "./components/UserItems.vue";
import AdminItems from "./components/AdminItems.vue";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "Home", component: ItemTable },
    { path: "/:pathMatch(.*)*", name: "NotFound", component: NotFound },
    { path: "/login", name: "Login", component: Login },
    { path: "/items", name: "Items", component: AdminItems },
    {
      path: "/users/:userID(\\d+)/items",
      name: "UserItems",
      component: UserItems,
    },
  ],
});
