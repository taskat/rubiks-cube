import{ matrix, multiply } from "mathjs";
import { Identity } from "./rotations";
import Coord from "./coord";
import SideCoord from "./sidecoord";

export default class Piece {
    constructor(id, coord, cubeData) {
      this.id = id;
      this.coord = coord;
      this.faces = this.getFaces(cubeData.vmin, cubeData.vmax);
      this.accTransform3 = Identity;
      this.sideCoords = createSideCoords(cubeData, coord);
    }
  
    getFaces(vmin, vmax) {
      return {
        up: this.coord.y === vmax ? "U" : "-",
        down: this.coord.y === vmin ? "D" : "-",
        left: this.coord.x === vmin ? "L" : "-",
        right: this.coord.x === vmax ? "R" : "-",
        front: this.coord.z === vmax ? "F" : "-",
        back: this.coord.z === vmin ? "B" : "-"
      };
    }
  
    rotatePiece(rotationMatrix3) {
      const vector3 = matrix([this.coord.x, this.coord.y, this.coord.z]);
      const rotatedVector3 = multiply(vector3, rotationMatrix3);
      const x = rotatedVector3.get([0]);
      const y = rotatedVector3.get([1]);
      const z = rotatedVector3.get([2]);
      this.coord = new Coord(x, y, z);
      this.accTransform3 = multiply(this.accTransform3, rotationMatrix3);
    }
  
    isInCoordsList(coordsList) {
      return coordsList.coords.findIndex(coord => this.coord.equal(coord)) >= 0;
    }

    getSideCoord(side) {
      return this.sideCoords.find(sideCoord => sideCoord.side === side);
    }
  }

  function createSideCoords(cubeData, coord) {
    const sideCoords = [];
    const vmin = cubeData.vmin;
    const vmax = cubeData.vmax;
    const maxCoord = cubeData.cubeSize - 1;
    if (coord.y === vmin) {
      sideCoords.push(new SideCoord("D", coord.x - vmin, coord.z - vmin));
    }
    if (coord.y === vmax) {
      sideCoords.push(new SideCoord("U", coord.x - vmin, coord.z - vmin));
    }
    if (coord.x === vmin) {
      sideCoords.push(new SideCoord("L", coord.y - vmin, coord.z - vmin));
    }
    if (coord.x === vmax) {
      sideCoords.push(new SideCoord("R", coord.y - vmin, coord.z - vmin));
    }
    if (coord.z === vmin) {
      sideCoords.push(new SideCoord("B", coord.x - vmin, coord.y - vmin));
    }
    if (coord.z === vmax) {
      sideCoords.push(new SideCoord("F", coord.x - vmin, coord.y - vmin));
    }
    return sideCoords;
  }