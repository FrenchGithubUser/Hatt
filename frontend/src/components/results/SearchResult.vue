<template>
  <div class="result">
    <div class="top-line">
      <div class="source-name">{{ result.Website }}</div>
      <CompatibleDownloaders
        v-if="result.CompatibleDownloaders.length !== 0"
        :downloaders="result.CompatibleDownloaders"
      />
    </div>
    <div v-if="result.Items === null" class="no-result">No result</div>
    <div
      class="items shadow-1"
      v-if="result.Items !== null && result.Items[0].Name !== 'error'"
    >
      <div
        class="item cursor-pointer"
        v-for="item in result.Items"
        :key="item.Link"
        @click="itemClicked(item)"
      >
        <div v-if="item.Thumbnail !== ''" class="thumbnail-wrapper">
          <img
            referrerPolicy="no-referrer"
            :src="item.Thumbnail"
            loading="lazy"
            class="thumbnail"
          />
        </div>
        <div :class="{ name: true, 'no-thumbnail': item.Thumbnail === '' }">
          {{ item.Name }}
        </div>
        <div class="metadata">
          <div class="info" v-for="(info, i) in item.Metadata" :key="i">
            {{ info }}
          </div>
        </div>
      </div>
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

export default defineComponent({
  name: 'SearchResult',
  components: { CompatibleDownloaders },
  data() {
    return {}
  },
  props: {
    result: { type: Object },
  },
  methods: {
    itemClicked(item) {
      window['runtime']['BrowserOpenURL'](item.Link)
    },
  },
  computed: {},
  created() {},
})
</script>
<style lang="scss" scoped>
.result {
  max-width: 100%;
  .top-line {
    display: flex;
    align-items: center;
    margin-bottom: 10px;
    .source-name {
      font-size: 1.4em;
      width: fit-content;
      background-color: $primary;
      color: white;
      padding: 7px;
      border-radius: 15px;
      margin-right: 10px;
    }
  }
  .items {
    overflow-y: hidden;
    overflow-x: scroll;
    display: flex;
    align-items: center;
    vertical-align: middle;
    padding: 15px;
    border-radius: 15px;
    .item {
      display: flex;
      flex-direction: column;
      align-items: center;
      text-align: center;
      margin-bottom: 5px;
      margin-right: 15px;
      .thumbnail-wrapper {
        width: 200px;
        height: 200px;
        .thumbnail {
          width: 100%;
          height: 100%;
          object-fit: contain;
        }
      }
      .name {
        width: 200px !important;
        &:not(.no-thumbnail) {
          overflow: hidden;
          text-overflow: ellipsis;
          display: -webkit-box;
          -webkit-line-clamp: 2; /* number of lines to show */
          line-clamp: 2;
          -webkit-box-orient: vertical;
          margin-top: 5px;
        }
        &.no-thumbnail {
          height: 150px;
          display: flex;
          align-items: center;
          justify-content: center;
          border: $primary 2px solid;
          border-radius: 15px;
          padding: 5px;
        }
      }
      .metadata {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-wrap: wrap;
        font-size: 0.85em;
        .info {
          background-color: $primary;
          color: white;
          padding: 3px;
          border-radius: 5px;
          margin: 2px;
        }
      }
    }
  }
}
</style>
