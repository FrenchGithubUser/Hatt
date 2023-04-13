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
          @change="save"
        />
        <q-input
          :bg-color="$q.dark.isActive ? 'blue-grey-8' : 'blue-grey-3'"
          v-model.number="values.thumbnailsSize"
          class="input"
          suffix="px"
          outlined
          dense
          @keyup.enter="save"
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
    <div class="item open-link-action">
      <div class="name">{{ $t('settings.item_clicked_action') }}</div>
      <q-select
        class="select"
        outlined
        v-model="values.itemClickedAction"
        :options="itemClickedActions"
        dense
        @update:model-value="save"
      >
        <template v-slot:option="scope">
          <q-item v-bind="scope.itemProps">
            <q-item-section>
              <q-item-label>{{ $t('settings.' + scope.opt) }}</q-item-label>
            </q-item-section>
          </q-item>
        </template>
        <template v-slot:selected-item="action">
          {{ $t('settings.' + action.opt) }}
        </template>
      </q-select>
    </div>
    <div class="item">
      <div class="name">{{ $t('settings.xxx') }}</div>
      <div class="setters">
        <q-checkbox v-model="values.xxx" @click="save">
          {{ values.xxx ? $t('expressions.activated') : $t('expressions.off') }}
        </q-checkbox>
      </div>
    </div>
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
        thumbnailsSize: 0,
        darkMode: false,
        itemClickedAction: '',
        xxx: false,
      },
      itemClickedActions: ['open_browser', 'copy_link'],
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
    // if the user just updated hatt from a version that didn't have this setting
    if (this.values.itemClickedAction === '') {
      this.values.itemClickedAction = this.itemClickedActions[0]
    }
  },
})
</script>
<style lang="scss" scoped>
.general {
  .item {
    margin-bottom: 20px;
    .name {
      color: var(--q-primary);
      font-weight: bold;
      margin-bottom: 7px;
    }
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
  .open-link-action {
    .select {
      width: 200px;
    }
  }
}
</style>
