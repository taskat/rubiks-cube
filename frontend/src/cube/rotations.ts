import { matrix, identity } from "mathjs";

export const Identity = identity(3) as math.Matrix;

class CosSin {
  cos: number;
  sin: number;

  constructor(cos: number, sin: number) {
    this.cos = cos
    this.sin = sin
  }
}

function calcCosSin(degrees: number): CosSin {
  const radians = degrees * Math.PI / 180
  const cos = Math.trunc(Math.cos(radians))
  const sin = Math.trunc(Math.sin(radians))
  return new CosSin(cos, sin)
}

export function rotateX(degrees: number): math.Matrix {
  const cs = calcCosSin(degrees)
  return matrix([
    [1, 0, 0],
    [0, cs.cos, cs.sin],
    [0, -cs.sin, cs.cos]
  ])
}

export function rotateY(degrees: number): math.Matrix {
  const cs = calcCosSin(degrees)
  return matrix([
    [cs.cos, 0, -cs.sin],
    [0, 1, 0],
    [cs.sin, 0, cs.cos]
  ])
}

export function rotateZ(degrees: number): math.Matrix {
  const cs = calcCosSin(degrees)
  return matrix([
    [cs.cos, -cs.sin, 0],
    [cs.sin, cs.cos, 0],
    [0, 0, 1]
  ])
}
