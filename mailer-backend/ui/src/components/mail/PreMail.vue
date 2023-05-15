<template>
  <div
    :class="isMailSelected ? 'bg-slate-100' : ''"
    class="w-full pb-4 border-y border-gray-300 hover:bg-slate-100 hover:cursor-pointer"
    role="button"
    @click="handleMailClick"
  >
    <div class="flex items-center py-2">
      <div
        class="flex items-center justify-center w-12 h-12 rounded-full shadow-lg"
      >
        <p class="text-gray-600 font-medium">{{ nameInitials }}</p>
      </div>
      <div class="w-full ml-2">
        <div class="flex items-center justify-between">
          <h3 class="text-sm line-clamp-1 text-gray-800">
            {{ mail?.["x-from"] }}
          </h3>
          <p class="text-xs text-gray-600">{{ dateParsed }}</p>
        </div>
        <div class="flex items-center">
          <p class="mr-1 text-xs text-indigo-600">
            {{ mail?.subject }}
          </p>
        </div>
      </div>
    </div>
    <p class="text-justify text-xs line-clamp-2 text-gray-600">
      {{ mail?.content }}
    </p>
  </div>
</template>

<script lang="ts" setup>
import { computed, defineProps } from "vue";
import type { MailData } from "@/model/SearchResponseModel";
import { getInitialsName, parseDate } from "@/utils/utils";
import { useEmailStore } from "@/stores/email_store";

const props = defineProps<{ mailData: MailData }>();

const store = useEmailStore();

const { mailData } = props;
const mail = computed(() => mailData?.mail);

const handleMailClick = () => mailData && store.selectMail(mailData);
const dateParsed = computed(() => parseDate(mail?.value?.date));
const isMailSelected = computed(() => mailData?.id === store.mailSelected?.id);
const nameInitials = computed(() => getInitialsName(mail?.value?.from));
</script>