<script lang="ts" setup>
const props = defineProps<{ summary: string }>();
</script>

<template>
  <article class="w-2/3 border-2 border-gray-200 rounded-xl">
    <h3
      class="bg-slate-900 w-fit px-4 py-1 text-sm font-bold text-white rounded-tl-lg rounded-br-xl"
    >
      SUMMARY
    </h3>

    <p class="p-4 text-gray-700 text-justify">{{ props.summary }}</p>
  </article>
</template>
