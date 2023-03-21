<template>
  <div class="popup">
    <div class="title">
      {{ $t('credentials.enter_creds') }}
    </div>
    <div class="fields">
      <q-input
        class="field"
        v-model="this.fields[i]"
        outlined
        v-for="(field, i) in website.Fields"
        :key="field"
        :type="field === 'password' ? 'password' : 'text'"
        :label="field"
      />
    </div>
    <q-btn
      :label="$t('expressions.validate')"
      no-caps
      color="primary"
      @click="save"
    />
  </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'EditCredentials',
  data() {
    return {
      credentials: {},
      fields: ['', ''],
    }
  },
  props: {
    website: { type: Object },
  },
  methods: {
    save() {
      if (this.credentials.LoginInfo === null) {
        this.credentials.LoginInfo = {}
      }
      this.website.Fields.forEach((field, i) => {
        this.credentials.LoginInfo[field] = this.fields[i]
      })
      window['go']['helpers']['Helper']
        ['SaveUpdatedCredentials'](this.website.Name, this.credentials)
        .then(() => {
          this.$q.notify(
            this.$t('notifications.creds_updated', [this.website.Name])
          )
          this.$emit('saved')
        })
    },
  },
  computed: {},
  created() {
    window['go']['helpers']['Helper']
      ['DeserializeCredentials'](this.website.Name)
      .then((data) => {
        this.credentials = data
        this.website.Fields.forEach((field, i) => {
          this.fields[i] = this.credentials.LoginInfo
            ? this.credentials.LoginInfo[field]
            : ''
        })
      })
  },
})
</script>
<style lang="scss" scoped>
.popup {
  width: 400px;
  text-align: center;
  .title {
    font-size: 1.2em;
  }
  .fields {
    .field {
      margin: 20px 0px;
    }
  }
}
</style>
