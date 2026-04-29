<script setup lang="ts">
definePageMeta({ middleware: "auth" })
import { ref, onMounted } from "vue"
const api = useRuntimeConfig().public.apiBase
const customers = ref<any[]>([])
const error = ref("")

// create
const name = ref("")
const comment = ref("")
const status = ref("")
const score = ref(0)

// top edit card
const editId = ref("")
const editName = ref("")
const editComment = ref("")
const editStatus = ref("")
const editScore = ref(0)

// inline editing
const inlineEditId = ref("")
const inlineName = ref("")
const inlineComment = ref("")
const inlineStatus = ref("")
const inlineScore = ref(0)

// single get
const getSingleId = ref("")
const singleCustomer = ref<any>(null)
const singleError = ref("")

const token = useCookie("token")

async function fetchCustomers() {
  try {
    customers.value = await $fetch(`${api}/customers`, {
      headers: { Authorization: `Bearer ${token.value}` }
    })
  } catch {
    await navigateTo("/")
  }
}

async function createCustomer() {
  error.value = ""
  try {
    await $fetch(`${api}/customers`, {
      method: "POST",
      headers: { Authorization: `Bearer ${token.value}` },
      body: { name: name.value, comment: comment.value, status: status.value, score: score.value}
    })
    name.value = ""
    comment.value = ""
    status.value = ""
    score.value = 0
    await fetchCustomers()
  } catch {
    error.value = "create failed"
  }
}

async function updateCustomer(id: string, n: string, c: string, status: string, score: number) {
  error.value = ""
  try {
    await $fetch(`${api}/customers/${id}`, {
      method: "PUT",
      headers: { Authorization: `Bearer ${token.value}` },
      body: { name: n, comment: c, status: status, score: score }
    })
    inlineEditId.value = ""
    editId.value = ""
    editName.value = ""
    editComment.value = ""
    await fetchCustomers()
  } catch {
    error.value = "update failed"
  }
}

function selectTopEdit(c: any) {
  editId.value = c.id
  editName.value = c.name
  editComment.value = c.comment
  editStatus.value = c.status
  editScore.value = c.score
}

function startInlineEdit(c: any) {
  inlineEditId.value = c.id
  inlineName.value = c.name
  inlineComment.value = c.comment
  inlineStatus.value = c.status
  inlineScore.value = c.score
}

function cancelInline() {
  inlineEditId.value = ""
}

async function deleteCustomer(id: string) {
  error.value = ""
  try {
    await $fetch(`${api}/customers/${id}`, {
      method: "DELETE",
      headers: { Authorization: `Bearer ${token.value}` }
    })
    if (inlineEditId.value === id) inlineEditId.value = ""
    if (editId.value === id) { editId.value = ""; editName.value = ""; editComment.value = "" }
    await fetchCustomers()
  } catch {
    error.value = "delete failed"
  }
}

async function getSingleCustomer() {
  singleError.value = ""
  singleCustomer.value = null
  try {
    singleCustomer.value = await $fetch(`${api}/customers/${getSingleId.value}`, {
      headers: { Authorization: `Bearer ${token.value}` }
    })
  } catch {
    singleError.value = "customer not found"
  }
}

onMounted(fetchCustomers)
</script>

<template>
  <div class="min-h-screen bg-gray-50 p-8">
    <div class="max-w-3xl mx-auto">
      <h1 class="text-xl font-medium mb-6">Customers</h1>

      <!-- create -->
      <div class="bg-white border border-gray-200 rounded-xl p-5 mb-4">
        <p class="text-xs text-gray-400 uppercase tracking-wide mb-3">Add customer</p>
        <div class="flex gap-2">
          <input v-model="name" placeholder="Name"
            class="flex-1 px-3 py-2 text-sm border border-gray-200 rounded-lg bg-gray-50 focus:outline-none focus:border-gray-400" />
          <input v-model="comment" placeholder="Comment"
            class="flex-1 px-3 py-2 text-sm border border-gray-200 rounded-lg bg-gray-50 focus:outline-none focus:border-gray-400" />
            <select v-model="status"
              class="w-28 px-2 py-2 text-sm border border-gray-200 rounded-lg bg-gray-50 focus:outline-none focus:border-gray-400">
              <option value="" disabled>Status</option>
              <option value="active">active</option>
              <option value="inactive">inactive</option>
            </select>
            <input v-model.number="score" type="number" placeholder="Score"
              class="w-16 px-2 py-2 text-sm border border-gray-200 rounded-lg bg-gray-50 focus:outline-none focus:border-gray-400" />
          <button @click="createCustomer"
            class="px-4 py-2 text-sm font-medium bg-gray-900 text-white rounded-lg hover:bg-gray-700 transition-colors">
            Create
          </button>
        </div>
      </div>

        <!-- top edit card -->
        <div class="bg-white border border-gray-200 rounded-xl p-5 mb-4">
          <p class="text-xs text-gray-400 uppercase tracking-wide mb-3">Edit customer</p>
          <div class="flex gap-2">
            <input v-model="editId" placeholder="ID"
              class="w-12 px-2 py-2 text-sm border border-gray-200 rounded-lg bg-gray-50 focus:outline-none focus:border-gray-400" />
            <input v-model="editName" placeholder="Name"
              class="w-28 px-2 py-2 text-sm border border-gray-200 rounded-lg bg-gray-50 focus:outline-none focus:border-gray-400" />
            <input v-model="editComment" placeholder="Comment"
              class="flex-1 px-2 py-2 text-sm border border-gray-200 rounded-lg bg-gray-50 focus:outline-none focus:border-gray-400" />
            <select v-model="editStatus"
              class="w-28 px-2 py-2 text-sm border border-gray-200 rounded-lg bg-gray-50 focus:outline-none focus:border-gray-400">
              <option value="" disabled>Status</option>
              <option value="active">active</option>
              <option value="inactive">inactive</option>
            </select>
            <input v-model.number="editScore" type="number" placeholder="Score"
              class="w-16 px-2 py-2 text-sm border border-gray-200 rounded-lg bg-gray-50 focus:outline-none focus:border-gray-400" />
            <button @click="updateCustomer(editId, editName, editComment, editStatus, editScore)"
              class="px-4 py-2 text-sm font-medium bg-gray-900 text-white rounded-lg hover:bg-gray-700 transition-colors">
              Update
            </button>
          </div>
        </div>

      <!-- get single -->
      <div class="bg-white border border-gray-200 rounded-xl p-5 mb-6">
        <p class="text-xs text-gray-400 uppercase tracking-wide mb-3">Get customer</p>
        <div class="flex gap-2 mb-3">
          <input v-model="getSingleId" placeholder="ID"
            class="w-16 px-3 py-2 text-sm border border-gray-200 rounded-lg bg-gray-50 focus:outline-none focus:border-gray-400" />
          <button @click="getSingleCustomer"
            class="px-4 py-2 text-sm font-medium border border-gray-200 rounded-lg hover:bg-gray-100 transition-colors">
            Fetch
          </button>
        </div>
        <div v-if="singleCustomer" class="text-sm grid grid-cols-2 gap-x-6 gap-y-1 text-gray-600">
          <div><span class="text-gray-400">ID</span> &nbsp;{{ singleCustomer.id }}</div>
          <div><span class="text-gray-400"></span> &nbsp;{{}}</div>
          <div><span class="text-gray-400">Name</span> &nbsp;{{ singleCustomer.name }}</div>
          <div><span class="text-gray-400">Comment</span> &nbsp;{{ singleCustomer.comment }}</div>
          <div><span class="text-gray-400">Status</span> &nbsp;{{ singleCustomer.status }}</div>
          <div><span class="text-gray-400">Score</span> &nbsp;{{ singleCustomer.score }}</div>
          <div><span class="text-gray-400">Created</span> &nbsp;{{ singleCustomer.createdAt }}</div>
          <div><span class="text-gray-400">Updated</span> &nbsp;{{ singleCustomer.updatedAt }}</div>
        </div>
        <p v-if="singleError" class="text-xs text-red-500">{{ singleError }}</p>
      </div>

      <!-- list -->
      <div class="bg-white border border-gray-200 rounded-xl p-5">
        <p class="text-xs text-gray-400 uppercase tracking-wide mb-3">All customers</p>
        <ul class="flex flex-col gap-2">
          <li v-for="c in customers" :key="c.id" class="border border-gray-100 rounded-lg overflow-hidden">

              <!-- normal row -->
              <div v-if="inlineEditId !== c.id"
                class="flex items-center gap-3 px-4 py-3 bg-gray-50 text-sm">
                <span class="text-xs text-gray-400 w-5">#{{ c.id }}</span>
                <span class="font-medium w-24 truncate">{{ c.name }}</span>
                <span class="text-gray-500 flex-1 truncate">{{ c.comment }}</span>
                <span class="text-xs px-2 py-0.5 rounded-full"
                  :class="c.status === 'active' ? 'bg-green-50 text-green-700' : 'bg-gray-100 text-gray-500'">
                  {{ c.status }}
                </span>
                <span class="text-xs text-gray-400 w-12 text-right">{{ c.score }} pts</span>
                <div class="flex gap-1 ml-auto">
                  <button @click="startInlineEdit(c); selectTopEdit(c)"
                    class="px-3 py-1 text-xs border border-gray-200 rounded-md hover:bg-gray-100 transition-colors">
                    Edit
                  </button>
                  <button @click="deleteCustomer(c.id)"
                    class="px-3 py-1 text-xs border border-red-200 text-red-500 rounded-md hover:bg-red-50 transition-colors">
                    Delete
                  </button>
                </div>
              </div>

              <!-- inline edit row -->
              <div v-else class="flex items-center gap-2 px-4 py-3 bg-blue-50 text-sm">
                <span class="text-xs text-gray-400 w-5">#{{ c.id }}</span>
                <input v-model="inlineName" placeholder="Name"
                  class="flex-1 px-2 py-1 text-sm border border-blue-200 rounded-md bg-white focus:outline-none focus:border-blue-400" />
                <input v-model="inlineComment" placeholder="Comment"
                  class="flex-1 px-2 py-1 text-sm border border-blue-200 rounded-md bg-white focus:outline-none focus:border-blue-400" />
                <select v-model="inlineStatus"
                  class="px-2 py-1 text-sm border border-blue-200 rounded-md bg-white focus:outline-none focus:border-blue-400">
                  <option value="active">active</option>
                  <option value="inactive">inactive</option>
                </select>
                <input v-model.number="inlineScore" type="number" placeholder="Score"
                  class="w-16 px-2 py-1 text-sm border border-blue-200 rounded-md bg-white focus:outline-none focus:border-blue-400" />
                <div class="flex gap-1 ml-auto">
                  <button @click="updateCustomer(c.id, inlineName, inlineComment, inlineStatus, inlineScore)"
                    class="px-3 py-1 text-xs bg-gray-900 text-white rounded-md hover:bg-gray-700 transition-colors">
                    Save
                  </button>
                  <button @click="cancelInline"
                    class="px-3 py-1 text-xs border border-gray-200 rounded-md hover:bg-gray-100 transition-colors">
                    Cancel
                  </button>
                </div>
              </div>

          </li>
        </ul>
      </div>

      <p v-if="error" class="mt-4 text-xs text-red-500">{{ error }}</p>
    </div>
  </div>
</template>