<script lang="ts" setup>
import { NConfigProvider, NMessageProvider, dateZhCN, zhCN } from "naive-ui";
import HelloWord from './components/HelloWord.vue';
</script>

<template>
  <n-config-provider class="n-config-provider" :date-locale="dateZhCN" :locale="zhCN">
    <n-message-provider>
      <HelloWord></HelloWord>
    </n-message-provider>
  </n-config-provider>
</template>

<style></style>
