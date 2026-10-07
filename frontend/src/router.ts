import { createRouter, createWebHistory } from "vue-router";
import ItemTable from "./components/ItemTable.vue";
import NotFound from "./components/NotFound.vue";
import Login from "./components/Login.vue";
import ItemList from "./components/ItemList.vue";
import UserItems from "./components/UserItems.vue";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "Home", component: ItemTable },
    { path: "/:pathMatch(.*)*", name: "NotFound", component: NotFound },
    { path: "/login", name: "Login", component: Login },
    { path: "/items", name: "Items", component: ItemList },
    {
      path: "/users/:userID(\\d+)/items",
      name: "UserItems",
      component: UserItems,
    },
  ],
});
