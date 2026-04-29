<script setup lang="ts">
import { ref } from "vue"
const username = ref("")
const password = ref("")
const error = ref("")
const api = useRuntimeConfig().public.apiBase

async function login() {
  error.value = ""
  try {
    const res = await $fetch<{ token: string }>(`${api}/login`, {
      method: "POST",
      body: { username: username.value, password: password.value },
    })
    const token = useCookie("token")
    token.value = res.token
    await navigateTo("/customers")
  } catch (e: any) {
    error.value = "invalid credentials"
  }
}
</script>

<template>
  <div class="min-h-screen bg-gray-50 flex items-center justify-center">
    <div class="bg-white border border-gray-200 rounded-xl p-8 w-full max-w-sm">
      <h2 class="text-lg font-medium mb-6">Sign in</h2>
      <div class="mb-4">
        <label class="block text-xs text-gray-500 mb-1">Username</label>
        <input v-model="username" type="text" placeholder="admin"
          class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg bg-gray-50 focus:outline-none focus:border-gray-400" />
      </div>
      <div class="mb-6">
        <label class="block text-xs text-gray-500 mb-1">Password</label>
        <input v-model="password" type="password" placeholder="••••••"
          class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg bg-gray-50 focus:outline-none focus:border-gray-400" />
      </div>
      <button @click="login"
        class="w-full py-2 text-sm font-medium bg-gray-900 text-white rounded-lg hover:bg-gray-700 transition-colors">
        Login
      </button>
      <p v-if="error" class="mt-3 text-xs text-red-500">{{ error }}</p>
    </div>
  </div>
</template>