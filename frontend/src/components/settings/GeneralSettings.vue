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
          :bg-color="$q.dark.isActive ? 'blue-grey-8' : 'blue-grey-3'"
          v-model.number="values.thumbnailsSize"
          class="input"
          suffix="px"
          outlined
          dense
        />
      </div>
    </div>
    <div class="item dark-mode">
      <div class="name">{{ $t('settings.dark_mode') }}</div>
      <div class="setters">
        <q-checkbox v-model="values.darkMode" @click="toggleDarkMode">
          {{
            values.darkMode
              ? $t('expressions.activated')
              : $t('expressions.off')
          }}
        </q-checkbox>
      </div>
    </div>
    <q-btn
      class="save"
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
import { toggleDarkMode, updateSettings } from 'src/helpers/helpers.js'

export default defineComponent({
  name: 'GeneralSettings',
  data() {
    return {
      values: {
        thumbnailsSize: 5,
        darkMode: false,
      },
    }
  },
  props: {
    originalValues: { Type: Object },
  },
  methods: {
    toggleDarkMode() {
      toggleDarkMode()
      this.save()
    },
    save() {
      document.body.style.setProperty(
        '--thumbnails-size',
        this.values.thumbnailsSize + 'px'
      )
      window.settings.general = this.values
      updateSettings()
      this.$emit('updated')
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
    color: var(--q-primary);
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
  .save {
    margin-top: 30px;
  }
}
</style>
