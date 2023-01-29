<template>
  <div class="result">
    <div class="source-name">{{ result.Website }}</div>
    <div class="items">
      <div class="item" v-for="item in result.Items" :key="item.Link">
        <img :src="item.Thumbnail" alt="thumbnail" class="thumbnail" />
        <div class="name">{{ item.Name }}</div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'SearchResult',
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
  .source-name {
    font-size: 1.5em;
  }
  .items {
    overflow-y: scroll;
    display: flex;
    .item {
      text-align: center;
      margin: 5px 10px;
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
