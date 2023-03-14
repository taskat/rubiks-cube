import { rotateX, rotateY, rotateZ } from "./rotations";
import CubeData from "./cubedata";
import Piece from "./piece";
import Move from "./move";
import CoordsList from "./coordslist";

export default class Cube {
  cubeData: CubeData;
  pieces: Piece[];
  moveIdsToMoves: Map<number, Move>;
  moves: Move[];
  allMoves: Map<string, Move>;
  constructor(cubeSize: number) {
    this.cubeData = new CubeData(cubeSize);
    const allCoordsList = this.cubeData.makeAllCoordsList();
    this.pieces = allCoordsList.coords.map((coord, index) => new Piece(index, coord, this.cubeData));
    this.moveIdsToMoves = this.makeMoveIdsToMoves();
    this.moves = Array.from(this.moveIdsToMoves.values());
    this.allMoves = new Map(this.moves.map(move => [move.name, move]));
  }

  lookupMoveId(id: number): Move | undefined {
    return this.moveIdsToMoves.get(id);
  }

  makeMoves(moves: Move[]) {
    this.pieces = moves.reduce((pieces, move) => move.makeMove(pieces), this.pieces);
  }

  makeMoveIdsToMoves(): Map<number, Move> {
    const angles = [90, 180, 270];
    const xRotationMatrices3 = angles.map(rotateX);
    const yRotationMatrices3 = angles.map(rotateY);
    const zRotationMatrices3 = angles.map(rotateZ);
    const allCoordsList = this.cubeData.makeAllCoordsList();
    const xNames = ["L", "M", "R"];
    const yNames = ["D", "E", "U"];
    const zNames = ["B", "S", "F"];
    const slices: RotCoord[] = [
      ...this.cubeData.values.map((xSlice, index) => new RotCoord(xRotationMatrices3, allCoordsList.xSlice(xSlice), xNames[index])),
      ...this.cubeData.values.map((ySlice, index) => new RotCoord(yRotationMatrices3, allCoordsList.ySlice(ySlice), yNames[index])),
      ...this.cubeData.values.map((zSlice, index) => new RotCoord(zRotationMatrices3, allCoordsList.zSlice(zSlice), zNames[index])),
      new RotCoord(xRotationMatrices3, allCoordsList.xSlice(-1).concat(allCoordsList.xSlice(0)), "Lw"),
      new RotCoord(xRotationMatrices3, allCoordsList.xSlice(0).concat(allCoordsList.xSlice(1)), "Rw"),
      new RotCoord(yRotationMatrices3, allCoordsList.ySlice(-1).concat(allCoordsList.ySlice(0)), "Dw"),
      new RotCoord(yRotationMatrices3, allCoordsList.ySlice(0).concat(allCoordsList.ySlice(1)), "Uw"),
      new RotCoord(zRotationMatrices3, allCoordsList.zSlice(-1).concat(allCoordsList.zSlice(0)), "Bw"),
      new RotCoord(zRotationMatrices3, allCoordsList.zSlice(0).concat(allCoordsList.zSlice(1)), "Fw"),
      new RotCoord(xRotationMatrices3, allCoordsList, "x"),
      new RotCoord(yRotationMatrices3, allCoordsList, "y"),
      new RotCoord(zRotationMatrices3, allCoordsList, "z"),
    ];
    const nestedKvps = slices.map(this.makeKvpsForSlice.bind(this));
    return new Map(flatten(nestedKvps));
  }

  makeKvpsForSlice(rotCoord: RotCoord, index: number): Kv[] {
    const baseId = index * 3;
    const move90Id = baseId;
    const move180Id = baseId + 1;
    const move270Id = baseId + 2;
    const move90 = this.makeKvp(move90Id, rotCoord.rotationMatrix3[0], rotCoord.coordsList, 1, rotCoord.name);
    const move180 = this.makeKvp(move180Id, rotCoord.rotationMatrix3[1], rotCoord.coordsList, 2, rotCoord.name + "2");
    const move270 = this.makeKvp(move270Id, rotCoord.rotationMatrix3[2], rotCoord.coordsList, 1, rotCoord.name + "'");
    return [move90, move180, move270];
  }

  makeKvp(id: number, rotationMatrix3: math.Matrix, coordsList: CoordsList, numTurns: number, name: string): Kv {
    const key = id;
    const value = new Move(rotationMatrix3, coordsList, numTurns, name);
    return [key, value];
  }

  getPieces(coordsList: CoordsList): Piece[] {
    return coordsList.coords.map(coord =>
      this.pieces.find(piece => piece.coord.equal(coord))).filter(piece => piece !== undefined) as Piece[];
  }
}

class RotCoord {
  rotationMatrix3: math.Matrix[];
  coordsList: CoordsList;
  name: string;
  constructor(rotationMatrix3: math.Matrix[], coordsList: CoordsList, name: string) {
    this.rotationMatrix3 = rotationMatrix3;
    this.coordsList = coordsList;
    this.name = name;
  }
}

type Kv = [number, Move];

function flatten(xss: Kv[][]): Kv[] {
  return ([] as Kv[]).concat(...xss)
}
