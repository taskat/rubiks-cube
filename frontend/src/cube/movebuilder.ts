import CoordsList from "./coordslist";
import Move from "./move";

export default class MoveBuilder {
    rotationMatrix3: math.Matrix[];
    coordsList: CoordsList;
    name: string;
    reversed: boolean;
    constructor(rotationMatrix3: math.Matrix[], coordsList: CoordsList, name: string) {
      this.rotationMatrix3 = rotationMatrix3;
      this.coordsList = coordsList;
      this.name = name;
      this.reversed = false;
    }

    build(): Move[] {
      let move90 = this.getMove(0, 1, "");
      const move180 = this.getMove(1, 2, "2");
      let move270 = this.getMove(2, 1, "'");
      if (this.reversed) {
        move90 = this.getMove(2, 1, "");
        move270 = this.getMove(0, 1, "'");
      }
      return [move90, move180, move270];
    }
  
    getMove(index: number, numTurns: number, suffix: string): Move {
      return new Move(this.rotationMatrix3[index], this.coordsList, numTurns, this.name + suffix);
    }

    reverse() {
      this.reversed = true;
    }
  }