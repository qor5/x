import { ChartOptions } from './presets.config'

export function useVxChartMergeOptsCallback(props: any) {
  const invokeMergeOptionsCallback = (
    options: ChartOptions,
    mergeCallbackOptions: { seriesData: any[] }
  ) => {
    if (props.mergeOptionsCallback) {
      props.mergeOptionsCallback(options, mergeCallbackOptions)
    }
  }

  return {
    invokeMergeOptionsCallback
  }
}
