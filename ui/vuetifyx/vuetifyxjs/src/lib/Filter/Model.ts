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
}

export interface autocompleteDataSource {
  remoteUrl: string
  isPaging: Boolean
  hasIcon: Boolean
  itemTitle: string
  itemValue: string
  itemIcon: string
  pageKey: string
  pagesKey: string
  pageSizeKey: string
  totalKey: string
  itemsKey: string
  currentKey: string
  searchKey: string
  page: number
  pageSize: number
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
}
