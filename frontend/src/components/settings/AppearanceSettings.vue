<template>
  <div class="appearance">
    <div class="item thumbnails-size">
      <div class="name">{{ $t('settings.thumbnails_size') }}</div>
      <div class="setters">
        <q-slider
          v-model="thumbnailsSize"
          :min="25"
          :max="100"
          class="slider"
        />
        <q-input
          bg-color="blue-grey-3"
          v-model="thumbnailsSize"
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
      thumbnailsSize: 5,
    }
  },
  props: {},
  methods: {},
  computed: {},
  created() {
    // get the user settings from the backend
    // update the settings everytime a value changes (on mouse release)
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
