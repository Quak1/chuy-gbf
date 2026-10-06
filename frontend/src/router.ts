import { createRouter, createWebHistory } from "vue-router";
import ItemTable from "./components/ItemTable.vue";
import HelloWorld from "./components/HelloWorld.vue";
import NotFound from "./components/NotFound.vue";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", component: ItemTable },
    { path: "/hello", component: HelloWorld },
    { path: "/:pathMatch(.*)*", name: "NotFound", component: NotFound },
  ],
});
