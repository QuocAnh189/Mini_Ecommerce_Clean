const formatNumber = (num: number): string => {
  return new Intl.NumberFormat('de-DE').format(num)
}

export default formatNumber
