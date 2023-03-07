<template>
  <div class="result">
    <div class="top-line">
      <div class="source-name shadow-3">{{ result.Website }}</div>
      <CompatibleDownloaders
        v-if="result.CompatibleDownloaders.length !== 0"
        :downloaders="result.CompatibleDownloaders"
      />
    </div>
    <div v-if="result.Items === null" class="no-result">No result</div>
    <div
      class="items shadow-3"
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
