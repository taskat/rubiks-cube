export function range(n) {
  Array.from(Array(n).keys())
}

export function flatten(xss) {
  [].concat(...xss)
}

export function closeTo(a, b) {
  Math.abs(a - b) <= 1e-12
}
