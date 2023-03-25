<template>
  <div class="general">
    <div class="item thumbnails-size">
      <div class="name">{{ $t('settings.thumbnails_size') }}</div>
      <div class="setters">
        <q-slider
          v-model="values.thumbnailsSize"
          :min="50"
          :max="300"
          class="slider"
        />
        <q-input
          bg-color="blue-grey-3"
          v-model.number="values.thumbnailsSize"
          class="input"
          suffix="px"
          outlined
          dense
        />
      </div>
    </div>
    <q-btn
      color="primary"
      :label="$t('settings.save')"
      icon="save"
      no-caps
      @click="save"
    />
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import { updateSettings } from 'src/helpers/helpers.js'

export default defineComponent({
  name: 'GeneralSettings',
  data() {
    return {
      values: {
        thumbnailsSize: 5,
      },
    }
  },
  props: {
    originalValues: { Type: Object },
  },
  methods: {
    save() {
      this.root = document.documentElement
      this.root.style.setProperty(
        '--thumbnails-size',
        this.values.thumbnailsSize + 'px'
      )
      window.settings.general = this.values
      updateSettings()
    },
  },
  computed: {},
  watch: {},
  created() {
    this.values = { ...this.originalValues }
  },
})
</script>
<style lang="scss" scoped>
.general {
  .name {
    color: $primary;
    font-weight: bold;
  }
  .thumbnails-size {
    .setters {
      display: flex;
      align-items: center;
      .slider {
        width: 40%;
        margin-top: 10px;
        margin-right: 25px;
      }
      .input {
        width: 80px;
      }
    }
  }
}
</style>
