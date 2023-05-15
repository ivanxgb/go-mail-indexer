<script lang="ts" setup>
import Search from "@/components/side/Search.vue";
import { useEmailStore } from "@/stores/email_store";
import PreMail from "@/components/mail/PreMail.vue";
import { storeToRefs } from "pinia";

const mailStore = useEmailStore();
const { mails } = storeToRefs(mailStore);
mailStore.fetchMails("ivan");
</script>

<template>
  <div class="border-x-2 mt-6 md:w-1/2 lg:w-1/3 2xl:w-1/5">
    <Search />
    <PreMail v-for="mail in mails" :key="mail.id" :mailData="mail" />
  </div>
</template>