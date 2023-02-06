<template>
  <div class="result">
    <div class="source-name">{{ result.Website }}</div>
    <div class="items">
      <div
        class="item"
        v-for="item in result.Items"
        :key="item.Link"
        @click="itemClicked(item)"
      >
        <img
          referrerPolicy="no-referrer"
          :src="item.Thumbnail"
          loading="lazy"
          alt="thumbnail"
          class="thumbnail cursor-pointer"
        />
        <div class="name">{{ item.Name }}</div>
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
      text-align: center;
      margin-bottom: 5px;
      margin-right: 15px;
      .thumbnail {
        max-width: 200px;
        max-height: 200px;
      }
      .name {
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 2; /* number of lines to show */
        line-clamp: 2;
        -webkit-box-orient: vertical;
      }
    }
  }
}
</style>
