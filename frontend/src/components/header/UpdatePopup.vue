<template>
  <div class="update-popup popup">
    <div class="title">
      {{ $t('update.update_available') }}
    </div>
    <div class="changelog" v-if="!updating">
      {{ details.ReleaseNotes }}
    </div>
    <div class="updating" v-else>
      {{ $t('update.updating') }}
      <q-spinner-tail color="primary" size="2em" />
    </div>

    <div class="buttons">
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
      updating: false,
    }
  },
  created() {},
  methods: {
    update() {
      this.updating = true
      window['go']['main']['App']['SelfUpdate']().then((result) => {
        if (result) {
          this.$q.notify(this.$t('update.update_successfull'))
        } else {
          this.$q.notify(this.$t('update.update_error'))
        }
        this.$emit('close')
      })
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
  .buttons {
    display: flex;
    justify-content: space-around;
  }
}
</style>
