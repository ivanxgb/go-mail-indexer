<template>
  <section class="w-full px-4 rounded-r-3xl">
    <div class="flex justify-between items-center h-32 border-b-2 mb-8">
      <div class="flex space-x-4 items-center">
        <div
          class="flex items-center justify-center w-12 h-12 rounded-full shadow-lg"
        >
          <p class="text-gray-600 font-medium">{{ nameInitials }}</p>
        </div>
        <div>
          <h3 class="font-semibold text-lg">{{ mail?.mail?.["x-from"] }}</h3>
          <p class="text-light text-gray-400">{{ mail?.mail.from }}</p>
        </div>
      </div>
      <div>
        <ul class="flex text-gray-400 space-x-4">
          <li class="w-6 h-6 cursor-pointer hover:text-black">
            <img alt="Delete" src="@/assets/trash.svg" />
          </li>
          <li class="w-6 h-6">
            <MailToolTip :mail="mail?.mail" />
          </li>
        </ul>
      </div>
    </div>
    <section>
      <h1 class="font-bold text-2xl">{{ mail?.mail.subject }}</h1>
      <article class="mt-8 text-gray-500 leading-7 tracking-wider">
        <p>
          {{ mail?.mail.content }}
        </p>
      </article>
    </section>
  </section>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import MailToolTip from "@/components/mail/MailToolTip.vue";
import { getInitialsName, parseDate } from "@/utils/utils";
import { useEmailStore } from "@/stores/email_store";

const mailStore = useEmailStore();
const mail = computed(() => mailStore.mailSelected);
const dateParsed = computed(() => parseDate(mail?.value?.mail.date));
const nameInitials = computed(() => getInitialsName(mail?.value?.mail.from));
</script>