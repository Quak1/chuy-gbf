import { createRouter, createWebHistory } from "vue-router";
import ItemTable from "./views/ItemTable.vue";
import NotFound from "./views/NotFound.vue";
import Login from "./views/Login.vue";
import UserItems from "./views/UserItems.vue";
import AdminItems from "./views/AdminItems.vue";
import TeamBuilder from "./views/TeamBuilder.vue";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "Home", component: ItemTable },
    { path: "/:pathMatch(.*)*", name: "NotFound", component: NotFound },
    { path: "/login", name: "Login", component: Login },
    { path: "/items", name: "Items", component: AdminItems },
    { path: "/teams/builder", name: "TeamBuilder", component: TeamBuilder },
    {
      path: "/users/:userID(\\d+)/items",
      name: "UserItems",
      component: UserItems,
    },
  ],
});
