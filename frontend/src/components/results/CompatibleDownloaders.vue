<template>
  <div class="compatible-downloaders">
    <q-btn
      icon="download"
      tooltip
      color="primary"
      no-caps
      dense
      @click="displayPopup = true"
    >
      <q-tooltip>
        {{ $t('results.compatible_downloaders') }}
      </q-tooltip>
    </q-btn>
    <q-dialog v-model="displayPopup"
      ><div class="popup">
        <div class="explanation">
          {{ $t('results.downloaders_explanation') }}
        </div>
        <div class="downloaders-list">
          <div
            class="downloader cursor-pointer"
            v-for="downloader in downloaders"
            :key="downloader.Name"
            @click="downloaderClicked(downloader.Link)"
          >
            {{ downloader.Name }}
          </div>
        </div>
      </div>
    </q-dialog>
  </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'CompatibleDownloaders',
  components: {},
  data() {
    return {
      displayPopup: false,
    }
  },
  props: {
    downloaders: { type: Array },
  },
  methods: {
    downloaderClicked(link) {
      window['runtime']['BrowserOpenURL'](link)
    },
  },
  computed: {},
  created() {},
})
</script>
<style lang="scss" scoped>
.downloaders-list {
  margin-top: 15px;
  display: flex;
  justify-content: center;
  .downloader {
    color: var(--q-primary);
    font-weight: bold;
    font-size: 1.2em;
  }
}
</style>
