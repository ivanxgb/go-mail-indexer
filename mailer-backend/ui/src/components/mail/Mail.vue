<script lang="ts" setup>
import { computed } from "vue";
import MailToolTip from "@/components/mail/MailToolTip.vue";
import Summary from "@/components/mail/Summary.vue";
import { cleanEnronFormat, getInitialsName, parseDate } from "@/utils/utils";
import { useEmailStore } from "@/stores/email_store";

const mailStore = useEmailStore();
const mail = computed(() => mailStore.mailSelected);
const fromCleaned = computed(() => cleanEnronFormat(mail?.value?.mail?.from));
const dateParsed = computed(() => parseDate(mail?.value?.mail.date));
const mailSummary = computed(() => mailStore.mailSummary);
const nameInitials = computed(() => getInitialsName(fromCleaned.value));
const xFromCleaned = computed(() =>
  cleanEnronFormat(mail?.value?.mail?.["x-from"])
);

const getSummary = () => mailStore.getSummary(mail.value?.mail.content!);
</script>

<template>
  <section class="w-full h-full overflow-y-auto px-4 rounded-r-3xl">
    <div class="flex justify-between items-center h-32 border-b-2 mb-8">
      <div class="flex space-x-4 items-center">
        <div
          class="flex items-center justify-center w-12 h-12 rounded-full shadow-lg"
        >
          <p class="text-gray-600 font-medium">{{ nameInitials }}</p>
        </div>
        <div>
          <h3 class="font-semibold text-lg">{{ xFromCleaned }}</h3>
          <p class="text-light text-gray-600">{{ fromCleaned }}</p>
        </div>
      </div>
      <div>
        <ul class="flex text-gray-400 space-x-4">
          <li
            v-if="mail?.mail.content"
            class="w-6 h-6 cursor-pointer"
            @click="getSummary()"
          >
            <img alt="summary" src="@/assets/message.svg" />
          </li>
          <li class="w-6 h-6">
            <MailToolTip :mail="mail?.mail" />
          </li>
        </ul>
      </div>
    </div>
    <section>
      <h1 class="font-bold text-2xl">{{ mail?.mail.subject }}</h1>

      <Summary
        v-if="mailSummary !== null"
        :summary="mailSummary"
        class="mt-4"
      />
      <article class="mt-8 text-base text-justify text-gray-700">
        <p class="whitespace-pre-line">
          {{ mail?.mail.content }}
        </p>
      </article>
    </section>
  </section>
</template>

<style scoped>
section {
  height: calc(100vh - 150px);
}
</style>