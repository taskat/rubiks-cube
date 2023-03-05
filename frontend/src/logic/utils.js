export function range(n) {
  return Array.from(Array(n).keys())
}

export function closeTo(a, b) {
  return Math.abs(a - b) <= 1e-12
}
