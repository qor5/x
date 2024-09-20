import { getCurrentInstance } from 'vue'

// return a utils function: hasEventListener
// this function is used to check if some event callback is applied to some components
/**
 * hasEventListener('click:ok');   // Should log 'onClick:ok'
 * hasEventListener('click-ok');   // Should log 'onClickOk'
 * hasEventListener('my-custom:event'); // Should log 'onMyCustom:event'
 * hasEventListener('multiple-parts-to-test'); // Should log 'onMultiplePartsToTest'
 */
export function useHasEventListener() {
  const instance = getCurrentInstance()

  const hasEventListener = (event: string): boolean => {
    // Match the separator '-' or ':'
    const separator = event.match(/[-:]/)?.[0]

    // Split the event name by '-' or ':'
    const parts = event.split(/[-:]/)

    let eventName = ''

    if (separator === '-') {
      // Capitalize each part and join without a separator for '-'
      const capitalizedParts = parts.map((part) => part.charAt(0).toUpperCase() + part.slice(1))
      eventName = 'on' + capitalizedParts.join('')
    } else if (separator === ':') {
      // Capitalize the first part and keep the second part as is for ':'
      eventName = 'on' + parts[0].charAt(0).toUpperCase() + parts[0].slice(1) + ':' + parts[1]
    }

    // console.log(eventName) // For testing purpose

    // Check if the event listener exists in vnode.props (assumed instance)
    return !!instance?.vnode.props?.[eventName]
  }

  return { hasEventListener }
}
