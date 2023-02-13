<template>
  <div class="result">
    <div class="source-name">{{ result.Website }}</div>
    <div class="items">
      <div
        class="item cursor-pointer"
        v-for="item in result.Items"
        :key="item.Link"
        @click="itemClicked(item)"
      >
        <img
          referrerPolicy="no-referrer"
          :src="item.Thumbnail"
          loading="lazy"
          class="thumbnail"
        />
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
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import { copyToClipboard, Notify } from 'quasar'

export default defineComponent({
  name: 'SearchResult',
  data() {
    return {}
  },
  props: {
    result: { type: Object },
  },
  methods: {
    itemClicked(item) {
      copyToClipboard(item.Link).then(() => {
        Notify.create('Link copied to clipboard')
      })
    },
  },
  computed: {},
  created() {},
})
</script>
<style lang="scss" scoped>
.result {
  max-width: 100%;
  .source-name {
    font-size: 1.4em;
    width: fit-content;
    background-color: $primary;
    color: white;
    padding: 7px;
    border-radius: 15px;
    margin-bottom: 10px;
  }
  .items {
    overflow-y: scroll;
    display: flex;
    .item {
      display: flex;
      flex-direction: column;
      align-items: center;
      text-align: center;
      margin-bottom: 5px;
      margin-right: 15px;
      .thumbnail {
        max-width: 200px;
        max-height: 200px;
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
