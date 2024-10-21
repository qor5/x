import { createI18n } from 'vue-i18n'
import { zhHans, en, ja } from 'vuetify/locale'

const messages = {
  zhHans: {
    $vuetify: {
      ...zhHans,
      datePicker: {
        ...zhHans.datePicker,
        title: '',
        header: '选择日期',
        okTips: '选择日期才能保存',
        saveBtn: '保存'
      }
    }
  },
  en: {
    $vuetify: {
      ...en,
      datePicker: {
        ...en.datePicker,
        title: '',
        header: 'Select date',
        okTips: 'You must select a date to save',
        saveBtn: 'Save'
      }
    }
  },
  ja: {
    $vuetify: {
      ...ja,
      datePicker: {
        ...ja.datePicker,
        title: '',
        header: '日付を選択',
        okTips: '日付を選択してから保存してください',
        saveBtn: 'ほぞん'
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
