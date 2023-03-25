import { i18n } from 'src/boot/i18n.js'
import { emitter } from 'src/boot/mitt.js'
import { copyToClipboard, Dark, Notify } from 'quasar'

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
export function copyLink(link) {
  copyToClipboard(link).then(() => {
    Notify.create(i18n.global.t('notifications.link_copied'))
  })
}
export function toggleDarkMode() {
  Dark.toggle()
  document.documentElement.style.setProperty(
    '--q-primary',
    Dark.isActive ? '#4d68aa' : '#1f2e5'
  )
}
