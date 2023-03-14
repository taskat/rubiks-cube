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
  constructor(cubeSize: number) {
    this.cubeData = new CubeData(cubeSize);
    const allCoordsList = this.cubeData.makeAllCoordsList();
    this.pieces = allCoordsList.coords.map((coord, index) => new Piece(index, coord, this.cubeData));
    this.moveIdsToMoves = this.makeMoveIdsToMoves();
    this.moves = Array.from(this.moveIdsToMoves.values());
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
    const slices: RotCoord[] = [
      ...this.cubeData.values.map(xSlice => new RotCoord(xRotationMatrices3, allCoordsList.xSlice(xSlice))),
      ...this.cubeData.values.map(ySlice => new RotCoord(yRotationMatrices3, allCoordsList.ySlice(ySlice))),
      ...this.cubeData.values.map(zSlice => new RotCoord(zRotationMatrices3, allCoordsList.zSlice(zSlice))),
      new RotCoord(xRotationMatrices3, allCoordsList.xSlice(-1).concat(allCoordsList.xSlice(0))),
      new RotCoord(xRotationMatrices3, allCoordsList.xSlice(0).concat(allCoordsList.xSlice(1))),
      new RotCoord(yRotationMatrices3, allCoordsList.ySlice(-1).concat(allCoordsList.ySlice(0))),
      new RotCoord(yRotationMatrices3, allCoordsList.ySlice(0).concat(allCoordsList.ySlice(1))),
      new RotCoord(zRotationMatrices3, allCoordsList.zSlice(-1).concat(allCoordsList.zSlice(0))),
      new RotCoord(zRotationMatrices3, allCoordsList.zSlice(0).concat(allCoordsList.zSlice(1))),
      new RotCoord(xRotationMatrices3, allCoordsList),
      new RotCoord(yRotationMatrices3, allCoordsList),
      new RotCoord(zRotationMatrices3, allCoordsList),
    ];
    const nestedKvps = slices.map(this.makeKvpsForSlice.bind(this));
    return new Map(flatten(nestedKvps));
  }

  makeKvpsForSlice(rotCoords: RotCoord, index: number): Kv[] {
    const baseId = index * 3;
    const move90Id = baseId;
    const move180Id = baseId + 1;
    const move270Id = baseId + 2;
    const move90 = this.makeKvp(move90Id, move270Id, rotCoords.rotationMatrix3[0], rotCoords.coordsList, 1);
    const move180 = this.makeKvp(move180Id, move180Id, rotCoords.rotationMatrix3[1], rotCoords.coordsList, 2);
    const move270 = this.makeKvp(move270Id, move90Id, rotCoords.rotationMatrix3[2], rotCoords.coordsList, 1);
    return [move90, move180, move270];
  }

  makeKvp(id: number, oppositeMoveId: number, rotationMatrix3: math.Matrix, coordsList: CoordsList, numTurns: number): Kv {
    const key = id;
    const value = new Move(id, oppositeMoveId,
      rotationMatrix3, coordsList, numTurns);
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
  constructor(rotationMatrix3: math.Matrix[], coordsList: CoordsList) {
    this.rotationMatrix3 = rotationMatrix3;
    this.coordsList = coordsList;
  }
}

type Kv = [number, Move];

function flatten(xss: Kv[][]): Kv[] {
  return ([] as Kv[]).concat(...xss)
}
