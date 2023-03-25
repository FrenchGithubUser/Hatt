<template>
  <div class="result">
    <div class="top-line">
      <div class="source-name shadow-3">{{ result.Website }}</div>
      <CompatibleDownloaders
        v-if="result.CompatibleDownloaders.length !== 0"
        :downloaders="result.CompatibleDownloaders"
      />
    </div>
    <div class="website-description shadow-3">
      {{ $t('website_descriptions.' + result.Website) }}
    </div>
    <br />
    <div class="custom-chip results-amount" v-if="result.Items">
      {{ result.Items.length }} {{ $t('results.results') }}
    </div>
    <div v-if="result.Items === null" class="no-result">No result</div>
    <div
      class="items"
      v-if="result.Items !== null && result.Items[0].Name !== 'error'"
    >
      <SearchResult
        v-for="item in result.Items"
        :key="item.Link"
        :item="item"
      />
    </div>
    <div class="error" v-else>
      <div
        class="error login-required"
        v-if="
          result.Items !== null &&
          result.Items[0].Metadata.name === 'login_required'
        "
      >
        login_required
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import CompatibleDownloaders from './CompatibleDownloaders.vue'
import SearchResult from './SearchResult.vue'

export default defineComponent({
  name: 'WebsiteResults',
  components: { CompatibleDownloaders, SearchResult },
  data() {
    return {}
  },
  props: {
    result: { type: Object },
  },
  methods: {},
  computed: {},
  created() {},
})
</script>
<style lang="scss" scoped>
.result {
  max-width: 100%;
  margin-top: 40px;
  .top-line {
    display: flex;
    align-items: center;
    margin-bottom: 10px;
    .source-name {
      font-size: 1.4em;
      width: fit-content;
      // background-color: $primary;
      // color: white;
      padding: 7px;
      border-radius: 15px;
      margin-right: 10px;
      text-transform: capitalize;
    }
  }
  .website-description {
    display: inline-block;
    margin-bottom: 5px;
    padding: 3px;
    border-radius: 7px;
  }
  .results-amount {
    display: inline-block;
  }
  .items {
    overflow-y: hidden;
    overflow-x: scroll;
    display: flex;
    align-items: center;
    padding: 15px;
    border-radius: 15px;
  }
}
</style>
