import { Notify } from 'quasar'
import { i18n } from 'src/boot/i18n.js'
import { emitter } from 'src/boot/mitt.js'

export function updateSettings(showNotification = true) {
  window['go']['main']['App']
    ['UpdateUserSettings'](window.settings)
    .then(() => {
      if (showNotification) {
        Notify.create(i18n.global.t('settings.saved'))
      }
      emitter.emit('settingsSaved')
      return true
    })
    .catch((e) => {
      return e
    })
}
