<template>
  <div class="item cursor-pointer shadow-3" @click="itemClicked(item)">
    <div v-if="!noThumbnail" class="thumbnail-wrapper">
      <img
        referrerPolicy="no-referrer"
        :src="item.Thumbnail"
        loading="lazy"
        class="thumbnail"
        @error="noThumbnail = true"
      />
    </div>
    <div :class="{ name: true, 'no-thumbnail': noThumbnail }">
      {{ item.Name }}
    </div>
    <div
      class="metadata"
      v-if="item.Metadata && Object.keys(item.Metadata).length !== 0"
    >
      <div class="custom-chip" v-for="(info, i) in item.Metadata" :key="i">
        {{ info }}
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import { copyLink } from 'src/helpers/helpers'

export default defineComponent({
  name: 'SearchResult',
  components: {},
  data() {
    return {
      noThumbnail: true,
    }
  },
  props: {
    item: { type: Object },
  },
  methods: {
    itemClicked(item) {
      // window['runtime']['BrowserOpenURL'](item.Link)
      copyLink(item.Link)
    },
  },
  computed: {},
  created() {
    this.noThumbnail = this.item.Thumbnail === ''
  },
})
</script>
<style lang="scss" scoped>
.item {
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
  text-align: center;
  margin-bottom: 5px;
  margin-right: 15px;
  height: calc(var(--thumbnails-size) + 120px);
  padding: 15px;
  border-radius: 15px;
  .thumbnail-wrapper {
    width: var(--thumbnails-size);
    height: var(--thumbnails-size);
    .thumbnail {
      width: 100%;
      height: 100%;
      object-fit: contain;
    }
  }
  .name {
    width: 200px !important;
    margin-bottom: 5px;
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
      height: var(--thumbnails-size);
      width: var(--thumbnails-size) !important;
      display: flex;
      align-items: center;
      justify-content: center;
      // border: var(--q-primary) 2px solid;
      // border-radius: 15px;
      padding: 5px;
    }
  }
  .metadata {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
    font-size: 0.85em;
  }
}
</style>
