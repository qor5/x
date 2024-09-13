export interface SelectOption {
  text: string
  value: string
}

export interface IndependentTranslations {
  filterBy: string
}

export interface linkageSelectData {
  items: Array<any>
  labels: Array<string>
  selectOutOfOrder: Boolean
  linkageSelectRemoteOptions: linkageSelectRemoteOptions
}

export interface autocompleteDataSource {
  remoteUrl: string
  isPaging: Boolean
  hasIcon: Boolean
  itemTitle: string
  itemValue: string
  page: number
  pageSize: number
  separator: string

  itemIcon: string

  pageField: string
  pagesField: string
  pageSizeField: string
  totalField: string
  itemsField: string
  currentField: string
  searchField: string
}

export interface linkageSelectRemoteOptions {
  remoteUrl: string
  isPaging: Boolean
  itemTitle: string
  itemValue: string
  page: number
  pageSize: number
  separator: string

  levelStart: number
  levelStep: number

  pageField: string
  pagesField: string
  pageSizeField: string
  totalField: string
  itemsField: string
  currentField: string
  searchField: string

  parentField: string
  parentIdField: string
  levelField: string
}

export interface dateOptions {
  label: string
  disabled: boolean
  loading: boolean
  dateFormat: string
  clearText: string
  okText: string
  datePickerProps: object
}

export interface FilterItem {
  key: string
  label: string
  folded: boolean
  itemType: string
  modifier: string
  valueIs: string
  disableChooseModifier?: boolean
  valuesAre: string[]
  selected?: boolean
  valueFrom?: string
  valueTo?: string
  inTheLastValue?: string
  inTheLastUnit?: string
  options?: SelectOption[]
  translations?: IndependentTranslations
  linkageSelectData?: linkageSelectData
  autocompleteDataSource: autocompleteDataSource
  dateOptions: dateOptions[]
}
