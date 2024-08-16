import { createI18n } from 'vue-i18n'
import { zhHans, en, ja } from 'vuetify/locale'

const messages = {
  zhHans: {
    $vuetify: {
      ...zhHans,
      datePicker: {
        ...zhHans.datePicker,
        title: '',
        header: '选择日期'
      }
    }
  },
  en: {
    $vuetify: {
      ...en,
      datePicker: {
        ...en.datePicker,
        title: '',
        header: 'Select date'
      }
    }
  },
  ja: {
    $vuetify: {
      ...ja,
      datePicker: {
        ...ja.datePicker,
        title: '',
        header: '日付を選択'
      }
    }
  }
}

export default createI18n({
  legacy: false, // Vuetify does not support the legacy mode of vue-i18n
  locale: 'en',
  fallbackLocale: 'en',
  messages
})
