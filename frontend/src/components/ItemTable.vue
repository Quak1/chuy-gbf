<script setup lang="ts">
import { useFetch } from "../composables/useFetch";
import type { TableData } from "../models";

const { data, error, loading } = useFetch<TableData>("/api/users/items");
</script>

<template>
  <div>
    <div v-if="loading">Loading...</div>
    <div v-if="error">{{ error }}</div>
    <div v-if="data">
      <table>
        <thead>
          <tr>
            <th>Username</th>
            <th v-for="item in data.items" :key="item.id">
              {{ item.name.toUpperCase() }}
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in data.users" :key="user.id">
            <td>{{ user.username }}</td>
            <td v-for="item in data.items" :key="item.id">
              {{ user.values[item.id] }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
