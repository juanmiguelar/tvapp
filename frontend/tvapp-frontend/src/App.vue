<template>
  <div>
    <NewsTable />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import NewsTable from './components/NewsTable.vue';

export default defineComponent({
  components: {
    NewsTable,
  },
});
</script>
