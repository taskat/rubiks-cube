import * as U from "./utils"

export default class CubeData {
  constructor(cubeSize) {
    this.cubeSize = cubeSize;
    this.isEvenSizedCube = cubeSize % 2 === 0;
    this.values = U.range(cubeSize)
      .map(v => v - Math.floor(cubeSize / 2))
      .map(v => this.isEvenSizedCube && v >= 0 ? v + 1 : v);
    this.vmin = Math.min(...this.values);
    this.vmax = Math.max(...this.values);
  }

  * allCoordsGenerator() {
    const isFace = v => v === this.vmin || v === this.vmax
    for (const x of this.values) {
      for (const y of this.values) {
        for (const z of this.values) {
          if (isFace(x) || isFace(y) || isFace(z)) {
            yield [x, y, z];
          }
        }
      }
    }
  }

  makeAllCoordsList() {
    return Array.from(this.allCoordsGenerator());
  }
}
