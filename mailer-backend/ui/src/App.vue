<script lang="ts" setup>
import { RouterView } from 'vue-router'
import Header from '@/components/header/Header.vue'
</script>

<template>
  <header>
    <Header />
  </header>

  <main>
    <RouterView />
  </main>
</template>