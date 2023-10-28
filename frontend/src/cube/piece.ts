import{ matrix, multiply } from "mathjs";
import { Identity } from "./rotations";
import SideCoord from "./sidecoord";
import { Side } from "./side";
import Coord from "./coord";
import CoordsList from "./coordslist";
import { SideCoordCalculator } from "./sidecoordcalculator";

export default class Piece {
    id: number;
    coord: Coord;
    accTransform3: math.Matrix;
    sideCoords: SideCoord[];
    constructor(id: number, coord: Coord, calculator: SideCoordCalculator) {
      this.id = id;
      this.coord = coord;
      this.accTransform3 = Identity;
      this.sideCoords = calculator.calculate(coord);
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
