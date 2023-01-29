<template>
  <div class="results">
    <SearchResult
      class="result"
      v-for="source in results"
      :key="source.Website"
      :result="source"
    />
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import SearchResult from 'src/components/results/SearchResult.vue'

export default defineComponent({
  name: 'SearchResults',
  components: { SearchResult },
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
  .result {
    margin-bottom: 15px;
  }
}
</style>
