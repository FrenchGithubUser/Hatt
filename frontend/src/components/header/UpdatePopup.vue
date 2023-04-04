<template>
  <div class="update-popup popup">
    <div class="title">
      {{ $t('update.update_available') }}
    </div>
    <div class="changelog" v-if="status === ''">
      {{ details.ReleaseNotes }}
    </div>
    <div class="updating" v-if="status === 'updating'">
      {{ $t('update.updating') }}
      <q-spinner-tail color="primary" size="2em" />
    </div>
    <div class="error" v-if="status === 'error'">
      <div class="explanation">
        <q-icon name="error" color="negative" size="3em" />
        <div>
          {{ $t('update.error_explanation') }}
        </div>
      </div>
      <ul>
        <li class="solution">
          {{ $t('update.launch_as_admin') }}
        </li>
        <li class="solution">
          {{ $t('update.package_manager') }}
        </li>
        <li class="solution">
          {{ $t('update.manual_download') }}
        </li>
      </ul>
      <q-btn
        :label="$t('update.download_manually')"
        color="primary"
        no-caps
        @click="downloadManually"
      />
    </div>

    <div class="buttons" v-if="status === ''">
      <q-btn
        :label="$t('update.update_now')"
        color="primary"
        no-caps
        @click="update"
      />
      <q-btn
        :label="$t('update.later')"
        color="primary"
        no-caps
        @click="$emit('close')"
      />
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'UpdatePopup',
  components: {},
  props: {
    details: { type: Object },
  },
  data() {
    return {
      status: '',
    }
  },
  created() {},
  methods: {
    update() {
      this.status = 'updating'
      window['go']['main']['App']['SelfUpdate']().then((result) => {
        if (result) {
          this.$q.notify(this.$t('update.update_successfull'))
          this.$emit('close')
        } else {
          this.status = 'error'
        }
      })
    },
    downloadManually() {
      window['runtime']['BrowserOpenURL'](this.details.URL)
    },
  },
  watch: {},
  computed: {},
})
</script>
<style lang="scss" scoped>
.update-popup {
  width: 500px;
  .title {
    text-align: center;
    font-size: 1.3em;
  }
  .changelog {
    white-space: pre-line;
    margin: 15px;
  }
  .updating {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin: 25px;
  }
  .error {
    text-align: center;
    .explanation {
      margin: 15px 0px;
      font-weight: bold;
    }
    ul {
      text-align: left;
      li {
        margin: 3px 0px;
      }
    }
  }
  .buttons {
    display: flex;
    justify-content: space-around;
  }
}
</style>
