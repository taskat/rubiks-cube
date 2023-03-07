import { matrix, identity } from "mathjs";

export const Identity = identity(3)

export class CosSin {
  constructor(cos, sin) {
    this.cos = cos
    this.sin = sin
  }
}

function calcCosSin(degrees) {
  const radians = degrees * Math.PI / 180
  const cos = Math.trunc(Math.cos(radians))
  const sin = Math.trunc(Math.sin(radians))
  return new CosSin(cos, sin)
}

export function rotateX(degrees) {
  const cs = calcCosSin(degrees)
  return matrix([
    [1, 0, 0],
    [0, cs.cos, cs.sin],
    [0, -cs.sin, cs.cos]
  ])
}

export function rotateY(degrees) {
  const cs = calcCosSin(degrees)
  return matrix([
    [cs.cos, 0, -cs.sin],
    [0, 1, 0],
    [cs.sin, 0, cs.cos]
  ])
}

export function rotateZ(degrees) {
  const cs = calcCosSin(degrees)
  return matrix([
    [cs.cos, -cs.sin, 0],
    [cs.sin, cs.cos, 0],
    [0, 0, 1]
  ])
}
