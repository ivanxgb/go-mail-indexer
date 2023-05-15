<script lang="ts" setup>
import { computed, ref } from "vue";
import type { Mail } from "@/model/SearchResponseModel";
import { parseDate } from "@/utils/utils";

const props = defineProps<{ mail?: Mail }>();
const mail = computed(() => props.mail);
const mailTo = computed(() => mail?.value?.to.join(", "));
const dateParsed = computed(() => parseDate(mail?.value?.date));
const showTooltip = ref(false);

const toggleTooltip = (show: boolean) => {
  console.log("showing tooltip");
  showTooltip.value = show;
};
</script>

<template>
  <div @mouseout="toggleTooltip(false)" @mouseover="toggleTooltip(true)">
    <img
      alt="Email Details"
      class="w-2 cursor-pointer"
      src="@/assets/info.svg"
    />

    <div
      :class="showTooltip ? 'display' : 'hidden'"
      class="relative top-2 right-60 z-20 w-64 bg-white shadow-lg p-4 rounded"
      role="tooltip"
    >
      <p>
        From: <span>{{ mail?.from }}</span>
      </p>
      <p class="line-clamp-3">
        To: <span>{{ mailTo }}</span>
      </p>
      <p>
        Date: <span>{{ dateParsed }}</span>
      </p>
      <p>
        Subject: <span>{{ mail?.subject }}</span>
      </p>
    </div>
  </div>
</template>

<style scoped>
div > p {
  @apply text-sm text-gray-600;
}

span {
  @apply text-xs text-gray-500;
}
</style>