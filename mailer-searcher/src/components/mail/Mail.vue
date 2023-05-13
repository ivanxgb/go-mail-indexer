<template>
  <section class="w-full px-4 rounded-r-3xl">
    <div class="flex justify-between items-center h-32 border-b-2 mb-8">
      <div class="flex space-x-4 items-center">
        <div class="flex items-center justify-center w-12 h-12 rounded-full shadow-lg">
          <p class="text-gray-600 font-medium">{{ nameInitials }}</p>
        </div>
        <div>
          <h3 class="font-semibold text-lg">Stefanny Guzman</h3>
          <p class="text-light text-gray-400">stefannyguzbar@gmail.com</p>
        </div>
      </div>
      <div>
        <ul class="flex text-gray-400 space-x-4">
          <li class="w-6 h-6 cursor-pointer hover:text-black">
            <svg
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
              />
            </svg>
          </li>
        </ul>
      </div>
    </div>
    <section>
      <h1 class="font-bold text-2xl">Team Appreciation</h1>
      <article class="mt-8 text-gray-500 leading-7 tracking-wider">
        <p>
          Design and develop enterprise-facing UI and consumer-facing UI as well as REST API
          backends.Work with Product Managers and User Experience designers to create an appealing
          user experience for desktop web and mobile web.
        </p>
      </article>
    </section>
  </section>
</template>

<script lang="ts" setup>
import { computed, defineProps } from 'vue'
import { getInitialsName } from '@/utils/utils'

const props = defineProps({
  msg: String,
  name: String,
  date: String
})

const nameInitials = computed(() => getInitialsName(props.name))
</script>