<script lang="ts" setup>
import LeftColumn from "@/components/side/LeftColumn.vue";
import Mail from "@/components/mail/Mail.vue";
import { useEmailStore } from "@/stores/email_store";
</script>

<template>
  <div class="flex">
    <LeftColumn class="ml-1" />
    <Mail v-if="useEmailStore().mailSelected !== null" class="flex-1" />
  </div>
</template>