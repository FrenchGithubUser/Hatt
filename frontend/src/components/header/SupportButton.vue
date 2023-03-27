<template>
  <div class="support">
    <q-icon
      name="volunteer_activism"
      size="2.5em"
      class="icon"
      @click="popupOpened = true"
    />
    <q-dialog v-model="popupOpened">
      <div class="popup" id="support-popup">
        <div class="title">{{ $t('support.support') }}</div>
        <div class="explanation">{{ $t('support.explanation') }} :</div>
        <div class="section shadow-3">
          <div class="section-title">
            {{ $t('support.donate') }}
          </div>
          <div class="section-explanation">
            {{ $t('support.donate_explanation') }}
          </div>
          <div class="addresses">
            <CryptocurrencyAddress
              v-for="crypto in cryptocurrencies"
              :key="crypto.name"
              :crypto="crypto"
            />
          </div>
          <div class="fiat-methods">
            <img
              v-for="method in fiatMethods"
              :key="method.name"
              :src="'images/fiat-methods/' + method.name + '.png'"
              :style="'width:' + method.width + ';'"
              @click="openLink(method.url)"
              class="cursor-pointer"
            />
          </div>
        </div>
        <div class="section shadow-3">
          <div class="section-title">
            {{ $t('support.contribute_on_gh') }}
          </div>
          <div class="section-explanation">
            {{ $t('support.contribute_on_gh_explanation') }}
          </div>
        </div>
        <div class="section shadow-3">
          <div class="section-title">
            {{ $t('support.get_famous') }}
          </div>
          <div class="section-explanation">
            {{ $t('support.get_famous_explanation') }}
          </div>
        </div>
      </div>
    </q-dialog>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import CryptocurrencyAddress from './CryptocurrencyAddress.vue'

export default defineComponent({
  name: 'SupportButton',
  components: { CryptocurrencyAddress },
  props: {},
  data() {
    return {
      popupOpened: false,
      cryptocurrencies: [
        {
          name: 'Monero',
          address:
            '46NLLW7dzu5jo2eZ3SiAKgQuVL1Jw8wPMSBAYA3eh4h334HzwMNFSXQ3V3PmXYEoMFXkt24pTHcD1X57KRePN8ehQXn3Ggt',
        },
        {
          name: 'Bitcoin',
          address: 'bc1qf4a44ae76txhmfxfsa875u8mv6murdwawt7msx',
        },
        {
          name: 'Ethereum',
          address: '0x134a0974f2fefF0F76276fBdD44439717B2b587B',
        },
      ],
      fiatMethods: [
        {
          name: 'BuyMeACoffee',
          url: 'https://www.buymeacoffee.com/FrGithubUser',
          width: '150px',
        },
        {
          name: 'kofi',
          url: 'https://ko-fi.com/frenchgithubuser',
          width: '120px',
        },
      ],
    }
  },
  created() {},
  methods: {
    openLink(link) {
      window['runtime']['BrowserOpenURL'](link)
    },
  },
  watch: {},
  computed: {},
})
</script>
<style lang="scss" scoped></style>

<style lang="scss">
#support-popup {
  text-align: center;
  width: 80vw !important;
  max-width: 80vw !important;
  .title {
    font-size: 1.3em;
    font-weight: bold;
    color: var(--q-primary);
    margin-bottom: 5px;
  }
  .section {
    text-align: left;
    margin-top: 15px;
    padding: 10px;
    border-radius: 15px;
    .section-title {
      font-size: 1.1em;
      font-weight: bold;
      margin-bottom: 5px;
    }
  }
  .addresses {
    display: flex;
    justify-content: center;
    margin-top: 10px;
  }
  .fiat-methods {
    margin-top: 25px;
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
    img {
      margin: 0px 7px;
    }
  }
}
</style>
