<template>
  <div class="w-full pb-4 border-y border-gray-300 hover:bg-slate-100 hover:cursor-pointer">
    <div class="flex items-center py-2">
      <div class="flex items-center justify-center w-12 h-12 rounded-full shadow-lg">
        <p class="text-gray-600 font-medium">{{ nameInitials }}</p>
      </div>
      <div class="w-full ml-2">
        <div class="flex items-center justify-between">
          <h3 class="text-sm line-clamp-1 text-gray-800">{{ name }}</h3>
          <p class="text-xs text-gray-600">{{ date }}</p>
        </div>
        <div class="flex items-center">
          <p class="mr-1 text-xs text-indigo-600">Team appreciation</p>
        </div>
      </div>
    </div>
    <p class="text-justify text-xs line-clamp-2 text-gray-600">
      Greatest appreciation to you and your team for the outstanding job you did for us. The website
      is just what we wanted, and we were thrilled with the speed your team exercised. We feel
      privileged to have eLab communications.
    </p>
  </div>
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