import { rotateX, rotateY, rotateZ } from "./rotations";
import CubeData from "./cubedata";
import Piece from "./piece";
import Move from "./move";

export default class Cube {
  constructor(cubeSize) {
    this.cubeData = new CubeData(cubeSize);
    const allCoordsList = this.cubeData.makeAllCoordsList();
    this.pieces = allCoordsList.coords.map((coord, index) => new Piece(index, coord, this.cubeData));
    this.moveIdsToMoves = this.makeMoveIdsToMoves();
    this.moves = Array.from(this.moveIdsToMoves.values());
  }

  lookupMoveId(id) {
    return this.moveIdsToMoves.get(id);
  }

  makeMoves(moves) {
    this.pieces = moves.reduce((pieces, move) => move.makeMove(pieces), this.pieces);
  }

  makeMoveIdsToMoves() {
    const angles = [90, 180, 270];
    const xRotationMatrices3 = angles.map(rotateX);
    const yRotationMatrices3 = angles.map(rotateY);
    const zRotationMatrices3 = angles.map(rotateZ);
    const allCoordsList = this.cubeData.makeAllCoordsList();
    const slices = [
      ...this.cubeData.values.map(xSlice => [xRotationMatrices3, allCoordsList.xSlice(xSlice)]),
      ...this.cubeData.values.map(ySlice => [yRotationMatrices3, allCoordsList.ySlice(ySlice)]),
      ...this.cubeData.values.map(zSlice => [zRotationMatrices3, allCoordsList.zSlice(zSlice)]),
      [xRotationMatrices3, allCoordsList.xSlice(-1).concat(allCoordsList.xSlice(0))],
      [xRotationMatrices3, allCoordsList.xSlice(0).concat(allCoordsList.xSlice(1))],
      [yRotationMatrices3, allCoordsList.ySlice(-1).concat(allCoordsList.ySlice(0))],
      [yRotationMatrices3, allCoordsList.ySlice(0).concat(allCoordsList.ySlice(1))],
      [zRotationMatrices3, allCoordsList.zSlice(-1).concat(allCoordsList.zSlice(0))],
      [zRotationMatrices3, allCoordsList.zSlice(0).concat(allCoordsList.zSlice(1))],
      [xRotationMatrices3, allCoordsList],
      [yRotationMatrices3, allCoordsList],
      [zRotationMatrices3, allCoordsList],
    ];
    const nestedKvps = slices.map(this.makeKvpsForSlice.bind(this));
    return new Map(flatten(nestedKvps));
  }

  makeKvpsForSlice([rotationMatrices3, coordsList], index) {
    const baseId = index * 3;
    const move90Id = baseId;
    const move180Id = baseId + 1;
    const move270Id = baseId + 2;
    const move90 = this.makeKvp(move90Id, move270Id, rotationMatrices3[0], coordsList, 1);
    const move180 = this.makeKvp(move180Id, move180Id, rotationMatrices3[1], coordsList, 2);
    const move270 = this.makeKvp(move270Id, move90Id, rotationMatrices3[2], coordsList, 1);
    return [move90, move180, move270];
  }

  makeKvp(id, oppositeMoveId, rotationMatrix3, coordsList, numTurns) {
    const key = id;
    const value = new Move(id, oppositeMoveId,
      rotationMatrix3, coordsList, numTurns);
    return [key, value];
  }

  getPieces(coordsList) {
    return coordsList.coords.map(coord =>
      this.pieces.find(piece => piece.coord.equal(coord)));
  }
}

function flatten(xss) {
  return [].concat(...xss)
}
