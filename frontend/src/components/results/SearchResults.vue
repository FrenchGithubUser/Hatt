<template>
  <div class="results">
    <WebsiteResults
      class="result"
      v-for="source in results"
      :key="source.Website"
      :result="source"
    />
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import WebsiteResults from 'src/components/results/WebsiteResults.vue'

export default defineComponent({
  name: 'SearchResults',
  components: { WebsiteResults },
  data() {
    return {}
  },
  props: {
    results: { type: Object },
  },
  methods: {},
  computed: {},
  created() {},
})
</script>
<style lang="scss" scoped>
.results {
  max-width: 90vw;
  margin-top: 15px;
  .result {
    margin-bottom: 15px;
  }
}
</style>
