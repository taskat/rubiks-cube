export function range(n) {
  return Array.from(Array(n).keys())
}

export function flatten(xss) {
  return [].concat(...xss)
}

export function closeTo(a, b) {
  return Math.abs(a - b) <= 1e-12
}
