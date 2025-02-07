<template>
  <Teleport to="body">
    <div
      v-if="visible"
      class="fixed top-4 right-4 z-50 flex items-center justify-center"
    >
      <div
        class="rounded-md bg-green-50 p-4 shadow-lg transition-all duration-300"
        :class="{
          'translate-x-0 opacity-100': visible,
          'translate-x-full opacity-0': !visible
        }"
      >
        <div class="flex">
          <div class="flex-shrink-0">
            <svg
              class="h-5 w-5 text-green-400"
              viewBox="0 0 20 20"
              fill="currentColor"
              aria-hidden="true"
            >
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-green-800">{{ message }}</p>
          </div>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const props = defineProps({
  message: {
    type: String,
    required: true
  },
  duration: {
    type: Number,
    default: 3000
  }
})

const visible = ref(false)
let timer = null

onMounted(() => {
  visible.value = true
  timer = setTimeout(() => {
    visible.value = false
  }, props.duration)
})

onUnmounted(() => {
  if (timer) {
    clearTimeout(timer)
  }
})
</script>