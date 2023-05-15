import { ref } from "vue";
import { defineStore } from "pinia";
import type { MailData } from "@/model/SearchResponseModel";
import { search } from "@/services/api_router";

export const useEmailStore = defineStore("emails", () => {
  const mails = ref<MailData[]>([]);
  const mailSelected = ref<MailData | null>(null);

  async function fetchMails(searchTerm: string = "") {
    mails.value = await search(searchTerm);
  }

  function selectMail(mail: MailData) {
    mailSelected.value = mail;
  }

  return { mails, mailSelected, selectMail, fetchMails };
});
