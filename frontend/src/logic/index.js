import * as math from "mathjs"
import * as CL from "./coordsLists"
import * as R from "./rotations"
import * as U from "./utils"
import CubeData from "./coordsLists"

export class Piece {
  constructor(id, x, y, z, cubeData) {
    this.id = id;
    this.x = x;
    this.y = y;
    this.z = z;
    this.faces = this.getFaces(cubeData.vmin, cubeData.vmax, x, y, z);
    this.accTransform3 = R.Identity;
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
    const vector3 = math.matrix([this.x, this.y, this.z]);
    const rotatedVector3 = math.multiply(vector3, rotationMatrix3);
    this.x = rotatedVector3.get([0]);
    this.y = rotatedVector3.get([1]);
    this.z = rotatedVector3.get([2]);
    this.accTransform3 = math.multiply(this.accTransform3, rotationMatrix3);
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

export class Cube {
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
    const xRotationMatrices3 = angles.map(R.rotateX);
    const yRotationMatrices3 = angles.map(R.rotateY);
    const zRotationMatrices3 = angles.map(R.rotateZ);
    const allCoordsList = this.cubeData.makeAllCoordsList();
    const slices = [
      ...this.cubeData.values.map(xSlice => [xRotationMatrices3, CL.xSliceCoordsList(allCoordsList, xSlice)]),
      ...this.cubeData.values.map(ySlice => [yRotationMatrices3, CL.ySliceCoordsList(allCoordsList, ySlice)]),
      ...this.cubeData.values.map(zSlice => [zRotationMatrices3, CL.zSliceCoordsList(allCoordsList, zSlice)]),
      [xRotationMatrices3, CL.xSliceCoordsList(allCoordsList, -1).concat(CL.xSliceCoordsList(allCoordsList, 0))],
      [xRotationMatrices3, CL.xSliceCoordsList(allCoordsList, 0).concat(CL.xSliceCoordsList(allCoordsList, 1))],
      [yRotationMatrices3, CL.ySliceCoordsList(allCoordsList, -1).concat(CL.ySliceCoordsList(allCoordsList, 0))],
      [yRotationMatrices3, CL.ySliceCoordsList(allCoordsList, 0).concat(CL.ySliceCoordsList(allCoordsList, 1))],
      [zRotationMatrices3, CL.zSliceCoordsList(allCoordsList, -1).concat(CL.zSliceCoordsList(allCoordsList, 0))],
      [zRotationMatrices3, CL.zSliceCoordsList(allCoordsList, 0).concat(CL.zSliceCoordsList(allCoordsList, 1))],
      [xRotationMatrices3, allCoordsList],
      [yRotationMatrices3, allCoordsList],
      [zRotationMatrices3, allCoordsList],
    ];
    const nestedKvps = slices.map(this.makeKvpsForSlice.bind(this));
    return new Map(U.flatten(nestedKvps));
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
