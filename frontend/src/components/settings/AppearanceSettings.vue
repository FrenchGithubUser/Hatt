<template>
  <div class="appearance">
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
  </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'AppearanceSettings',
  data() {
    return {
      values: {
        thumbnailsSize: 5,
      },
    }
  },
  props: {
    originalValues: { Type: Object },
    saved: { Type: Boolean },
  },
  methods: {
    save() {
      this.root = document.documentElement
      this.root.style.setProperty(
        '--thumbnails-size',
        this.values.thumbnailsSize + 'px'
      )
    },
  },
  computed: {},
  watch: {
    saved(newVal) {
      if (newVal) {
        this.save()
      }
    },
  },
  created() {
    this.values = { ...this.originalValues }
  },
})
</script>
<style lang="scss" scoped>
.appearance {
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
