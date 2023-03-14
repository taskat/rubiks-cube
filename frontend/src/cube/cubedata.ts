import CoordsList from "./coordslist";
import Coord from "./coord";

export default class CubeData {
  cubeSize: number;
  isEvenSizedCube: boolean;
  values: number[];
  vmin: number;
  vmax: number;
  constructor(cubeSize: number) {
    this.cubeSize = cubeSize;
    this.isEvenSizedCube = cubeSize % 2 === 0;
    this.values = range(cubeSize)
      .map(v => v - Math.floor(cubeSize / 2))
      .map(v => this.isEvenSizedCube && v >= 0 ? v + 1 : v);
    this.vmin = Math.min(...this.values);
    this.vmax = Math.max(...this.values);
  }

  * allCoordsGenerator() {
    const isFace = (v: number) => v === this.vmin || v === this.vmax
    for (const x of this.values) {
      for (const y of this.values) {
        for (const z of this.values) {
          if (isFace(x) || isFace(y) || isFace(z)) {
            yield new Coord(x, y, z);
          }
        }
      }
    }
  }

  makeAllCoordsList(): CoordsList {
    return new CoordsList(Array.from(this.allCoordsGenerator()));
  }
}

function range(n: number): number[] {
  return Array.from(Array(n).keys())
}
