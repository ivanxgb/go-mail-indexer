<template>
  <div
    class="relative w-full flex justify-start items-center text-gray-800 border-t-2"
  >
    <input
      v-model="searchInput"
      class="w-full pl-16 py-3 rounded placeholder-gray-500 focus:outline-none focus:ring-1 focus:ring-gray-600"
      placeholder="Search"
      trim="true"
      type="text"
    />
    <img
      alt="search"
      class="absolute w-4 left-4"
      src="../../assets/magnifier.svg"
    />
  </div>
</template>

<script lang="ts" setup>
import { ref, watch } from "vue";
import { useEmailStore } from "@/stores/email_store";
import { debounce } from "@/utils/utils";

const searchInput = ref<string>("");

const debounceWatch = debounce(() => {
  useEmailStore().fetchMails(searchInput.value);
});

watch(searchInput, (value, oldValue) => {
  if (value !== oldValue && value.trim() !== "") {
    debounceWatch();
  }
});
</script>