<script lang="ts" setup>
import Search from "@/components/side/Search.vue";
import { useEmailStore } from "@/stores/email_store";
import { onMounted } from "vue";
import PreMail from "@/components/mail/PreMail.vue";
import { storeToRefs } from "pinia";

const mailStore = useEmailStore();
const { mails } = storeToRefs(mailStore);

onMounted(() => {
  mailStore.fetchMails("ivan");
});
</script>

<template>
  <div class="border-x-2 mt-6 w-1/3 xl:w-1/4 2xl:w-1/5">
    <Search />
    <PreMail v-for="mail in mails" :key="mail.id" :mailData="mail" />
  </div>
</template>

<style scoped></style>