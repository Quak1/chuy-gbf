import { createRouter, createWebHistory } from "vue-router";
import ItemTable from "./components/ItemTable.vue";
import HelloWorld from "./components/HelloWorld.vue";
import NotFound from "./components/NotFound.vue";
import Login from "./components/Login.vue";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "Home", component: ItemTable },
    { path: "/hello", name: "Hello", component: HelloWorld },
    { path: "/:pathMatch(.*)*", name: "NotFound", component: NotFound },
    { path: "/login", name: "Login", component: Login },
  ],
});
