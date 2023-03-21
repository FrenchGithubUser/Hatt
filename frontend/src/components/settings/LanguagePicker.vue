<template>
  <div class="language-picker">
    <img
      :src="'images/lang-flags/' + currentLang + '.svg'"
      class="lang-flag"
      @click="popupOpened = true"
    />
    <q-dialog v-model="popupOpened">
      <div class="popup available-langs" id="language-picker">
        <div
          class="available-lang cursor-pointer shadow-3"
          v-for="lang in availableLangs"
          :key="lang"
          @click="updateLang(lang)"
        >
          <img :src="'images/lang-flags/' + lang + '.svg'" class="lang-flag" />
          {{ lang }}
        </div>
      </div>
    </q-dialog>
  </div>
</template>

<script>
import { updateSettings } from 'src/helpers/helpers'
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'LanguagePicker',
  components: {},
  props: {},
  data() {
    return {
      currentLang: '',
      availableLangs: ['en', 'fr'],
      popupOpened: true,
    }
  },
  created() {
    this.currentLang = window.settings.general.lang
    if (this.currentLang === '') {
      try {
        Device.getLanguageCode().then((lang) => {
          this.currentLang = lang.value.slice(0, 2)
        })
      } catch {
        this.currentLang = 'en'
      }
    }
    this.updateLang(this.currentLang)
  },
  methods: {
    updateLang(lang) {
      window.settings.general.lang = lang
      updateSettings(false)
      this.$i18n.locale = lang
      this.currentLang = lang
      this.popupOpened = false
      console.log(this.$i18n.locale)
    },
  },
  watch: {},
  computed: {},
})
</script>
<style lang="scss" scoped>
.language-picker {
  display: flex;
  justify-content: center;
  .lang-flag {
    width: 50px;
    border-radius: 7px;
  }
}
</style>
<style lang="scss">
#language-picker {
  &.available-langs {
    display: flex;
    flex-wrap: wrap;
    .available-lang {
      display: flex;
      align-items: center;
      margin: 5px;
      padding: 8px;
      border-radius: 7px;
      img {
        margin-right: 15px;
        width: 50px;
        border-radius: 7px;
      }
    }
  }
}
</style>
