<template>
  <div class="item cursor-pointer" @click="itemClicked(item)">
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
    <div class="metadata">
      <div class="custom-chip" v-for="(info, i) in item.Metadata" :key="i">
        {{ info }}
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import { copyToClipboard } from 'quasar'

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

      copyToClipboard(item.Link).then(() => {
        this.$q.notify(this.$t('notifications.link_copied'))
      })
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
  align-items: center;
  text-align: center;
  margin-bottom: 5px;
  margin-right: 15px;
  height: 100%;
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
  }
}
</style>
