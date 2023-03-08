import{ matrix, multiply } from "mathjs";
import { Identity } from "./rotations";
import SideCoord from "./sidecoord";
import { Side } from "./side";
import CubeData from "./cubedata";
import Coord from "./coord";
import CoordsList from "./coordslist";

export default class Piece {
    id: number;
    coord: Coord;
    accTransform3: math.Matrix;
    sideCoords: SideCoord[];
    constructor(id: number, coord: Coord, cubeData: CubeData) {
      this.id = id;
      this.coord = coord;
      this.accTransform3 = Identity;
      this.sideCoords = createSideCoords(cubeData, coord);
    }
  
    rotatePiece(rotationMatrix3: math.Matrix): void {
      const vector3 = matrix([this.coord.x, this.coord.y, this.coord.z]);
      const rotatedVector3 = multiply(vector3, rotationMatrix3);
      const x = rotatedVector3.get([0]);
      const y = rotatedVector3.get([1]);
      const z = rotatedVector3.get([2]);
      this.coord = new Coord(x, y, z);
      this.accTransform3 = multiply(this.accTransform3, rotationMatrix3);
    }
  
    isInCoordsList(coordsList: CoordsList): boolean {
      return coordsList.coords.findIndex(coord => this.coord.equal(coord)) >= 0;
    }

    getSideCoord(side: Side): SideCoord | undefined {
      return this.sideCoords.find(sideCoord => sideCoord.side === side);
    }
  }

  function createSideCoords(cubeData: CubeData, coord: Coord) {
    const sideCoords = [];
    const vmin = cubeData.vmin;
    const vmax = cubeData.vmax;
    if (coord.y === vmin) {
      sideCoords.push(new SideCoord(Side.Down, -coord.z + vmax, coord.x - vmin));
    }
    if (coord.y === vmax) {
      sideCoords.push(new SideCoord(Side.Up, coord.z - vmin, coord.x - vmin));
    }
    if (coord.x === vmin) {
      sideCoords.push(new SideCoord(Side.Left, -coord.y + vmax, coord.z - vmin));
    }
    if (coord.x === vmax) {
      sideCoords.push(new SideCoord(Side.Right, -coord.y + vmax, -coord.z + vmax));
    }
    if (coord.z === vmin) {
      sideCoords.push(new SideCoord(Side.Back, coord.y - vmin, coord.x - vmin));
    }
    if (coord.z === vmax) {
      sideCoords.push(new SideCoord(Side.Front, -coord.y + vmax, coord.x - vmin));
    }
    return sideCoords;
  }