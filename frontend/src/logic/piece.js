import{ matrix, multiply } from "mathjs";
import { Identity } from "./rotations";

export default class Piece {
    constructor(id, x, y, z, cubeData) {
      this.id = id;
      this.x = x;
      this.y = y;
      this.z = z;
      this.faces = this.getFaces(cubeData.vmin, cubeData.vmax, x, y, z);
      this.accTransform3 = Identity;
    }
  
    getFaces(vmin, vmax, x, y, z) {
      return {
        up: y === vmax ? "U" : "-",
        down: y === vmin ? "D" : "-",
        left: x === vmin ? "L" : "-",
        right: x === vmax ? "R" : "-",
        front: z === vmax ? "F" : "-",
        back: z === vmin ? "B" : "-"
      };
    }
  
    rotatePiece(rotationMatrix3) {
      const vector3 = matrix([this.x, this.y, this.z]);
      const rotatedVector3 = multiply(vector3, rotationMatrix3);
      this.x = rotatedVector3.get([0]);
      this.y = rotatedVector3.get([1]);
      this.z = rotatedVector3.get([2]);
      this.accTransform3 = multiply(this.accTransform3, rotationMatrix3);
    }
  
    isInCoordsList(coordsList) {
      return coordsList.findIndex(coords => this.hasCoords(coords)) >= 0;
    }
  
    hasCoords(coords) {
      return this.x === coords[0] &&
        this.y === coords[1] &&
        this.z === coords[2];
    }
  }