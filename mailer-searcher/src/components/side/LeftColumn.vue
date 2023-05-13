<script lang="ts" setup>
import Search from '@/components/side/Search.vue'
import PreMail from '@/components/mail/PreMail.vue'
</script>

<template>
  <div class="border-x-2 mt-6 w-1/3 xl:w-1/4 2xl:w-1/5">
    <Search />
    <PreMail
      class="px-4"
      date="Jul 25, 2023"
      msg="Hey! How is everything?"
      name="Stefanny Guzman"
    />
  </div>
</template>

<style scoped></style>