import { rotateX, rotateY, rotateZ } from "./rotations"
import CubeData from "./cubedata"
import Piece from "./piece"

export default class Cube {
  constructor(cubeSize) {
    this.cubeData = new CubeData(cubeSize);
    const allCoordsList = this.cubeData.makeAllCoordsList();
    this.pieces = allCoordsList.map(([x, y, z], index) => new Piece(index, x, y, z, this.cubeData));
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
      ...this.cubeData.values.map(xSlice => [xRotationMatrices3, xSliceCoordsList(allCoordsList, xSlice)]),
      ...this.cubeData.values.map(ySlice => [yRotationMatrices3, ySliceCoordsList(allCoordsList, ySlice)]),
      ...this.cubeData.values.map(zSlice => [zRotationMatrices3, zSliceCoordsList(allCoordsList, zSlice)]),
      [xRotationMatrices3, xSliceCoordsList(allCoordsList, -1).concat(xSliceCoordsList(allCoordsList, 0))],
      [xRotationMatrices3, xSliceCoordsList(allCoordsList, 0).concat(xSliceCoordsList(allCoordsList, 1))],
      [yRotationMatrices3, ySliceCoordsList(allCoordsList, -1).concat(ySliceCoordsList(allCoordsList, 0))],
      [yRotationMatrices3, ySliceCoordsList(allCoordsList, 0).concat(ySliceCoordsList(allCoordsList, 1))],
      [zRotationMatrices3, zSliceCoordsList(allCoordsList, -1).concat(zSliceCoordsList(allCoordsList, 0))],
      [zRotationMatrices3, zSliceCoordsList(allCoordsList, 0).concat(zSliceCoordsList(allCoordsList, 1))],
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
    const value = {
      id,
      oppositeMoveId,
      makeMove: this.rotatePieces(coordsList, rotationMatrix3),
      rotationMatrix3,
      coordsList,
      numTurns
    };
    return [key, value];
  }

  rotatePieces(coordsList, rotationMatrix3) {
    return pieces => {
      pieces.forEach(piece => {
        if(piece.isInCoordsList(coordsList)) {
          piece.rotatePiece(rotationMatrix3)
        }
      });
      return pieces;
    };
  }

  getPieces(coordsList) {
    return coordsList.map(coords =>
      this.pieces.find(piece => piece.hasCoords(coords)));
  }
}

function xSliceCoordsList(allCoordsList, xSlice) {
  return allCoordsList.filter(([x]) => x === xSlice);
} 

function ySliceCoordsList(allCoordsList, ySlice) {
  return allCoordsList.filter(([, y]) => y === ySlice);
}

function zSliceCoordsList(allCoordsList, zSlice) {
  return allCoordsList.filter(([, , z]) => z === zSlice);
}

function flatten(xss) {
  return [].concat(...xss)
}
