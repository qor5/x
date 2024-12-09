export function useColor(colorString: string) {
  const isRGBorHexColor = (colorStr: string) => /rgb|#/.test(colorStr)

  return {
    color: isRGBorHexColor(colorString) ? colorString : `rgb(var(--v-theme-${colorString}))`
  }
}
