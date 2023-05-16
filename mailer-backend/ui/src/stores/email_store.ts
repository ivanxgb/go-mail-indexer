import { ref } from "vue";
import { defineStore } from "pinia";
import type { MailData } from "@/model/SearchResponseModel";
import { fetchSearch, fetchSummary } from "@/services/api_router";

export const useEmailStore = defineStore("emails", () => {
  const mails = ref<MailData[]>([]);
  const mailSelected = ref<MailData | null>(null);
  const mailSummary = ref<string | null>(null);

  async function getSummary(mailContent: string) {
    const resp = await fetchSummary(mailContent);
    mailSummary.value = resp.content;
  }

  async function getMails(searchTerm: string = "") {
    mails.value = await fetchSearch(searchTerm);
  }

  function selectMail(mail: MailData) {
    mailSelected.value = mail;
    mailSummary.value = null;
  }

  return {
    mails,
    mailSelected,
    mailSummary,
    getSummary,
    selectMail,
    fetchMails: getMails
  };
});
