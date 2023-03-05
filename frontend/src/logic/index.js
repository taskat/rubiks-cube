import * as math from "mathjs"
import * as CL from "./coordsLists"
import * as R from "./rotations"
import * as U from "./utils"
import CubeData from "./coordsLists"

function coordsToFaces(vmin, vmax, x, y, z) {
  return {
    up: y === vmax ? "U" : "-",
    down: y === vmin ? "D" : "-",
    left: x === vmin ? "L" : "-",
    right: x === vmax ? "R" : "-",
    front: z === vmax ? "F" : "-",
    back: z === vmin ? "B" : "-"
  };
} 

function pieceHasCoords(piece, coords) {
  return piece.x === coords[0] &&
  piece.y === coords[1] &&
  piece.z === coords[2];
}
  
export function getPieces(cube, coordsList) {
  return coordsList.map(coords =>
    cube.pieces.find(piece => pieceHasCoords(piece, coords)));
}
  

function isPieceInCoordsList(piece, coordsList) {
  return coordsList.findIndex(coords => pieceHasCoords(piece, coords)) >= 0;
}
  

function rotatePiece(piece, rotationMatrix3) {
  const vector3 = math.matrix([piece.x, piece.y, piece.z])
  const rotatedVector3 = math.multiply(vector3, rotationMatrix3)
  return {
    ...piece,
    x: rotatedVector3.get([0]),
    y: rotatedVector3.get([1]),
    z: rotatedVector3.get([2]),
    accTransform3: math.multiply(piece.accTransform3, rotationMatrix3)
  }
}

function rotatePieces(coordsList, rotationMatrix3) {
  return pieces => {
    const result = pieces.map(piece => isPieceInCoordsList(piece, coordsList)
      ? rotatePiece(piece, rotationMatrix3)
      : piece);
    return result;
  };
}
  

function makeKvp(id, oppositeMoveId, rotationMatrix3, coordsList, numTurns) {
  const key = id
  const value = {
    id,
    oppositeMoveId,
    makeMove: rotatePieces(coordsList, rotationMatrix3),
    rotationMatrix3,
    coordsList,
    numTurns
  }
  return [key, value]
}

function makeKvpsForSlice([rotationMatrices3, coordsList], index) {
  const baseId = index * 3
  const move90Id = baseId
  const move180Id = baseId + 1
  const move270Id = baseId + 2
  const move90 = makeKvp(move90Id, move270Id, rotationMatrices3[0], coordsList, 1)
  const move180 = makeKvp(move180Id, move180Id, rotationMatrices3[1], coordsList, 2)
  const move270 = makeKvp(move270Id, move90Id, rotationMatrices3[2], coordsList, 1)
  return [move90, move180, move270]
}

const angles = [90, 180, 270]
const xRotationMatrices3 = angles.map(R.rotateX)
const yRotationMatrices3 = angles.map(R.rotateY)
const zRotationMatrices3 = angles.map(R.rotateZ)

function makeMoveIdsToMoves(cubeSize) {
  const cubeData = new CubeData(cubeSize)
  const allCoordsList = cubeData.makeAllCoordsList();
  const slices = [
    ...cubeData.values.map(xSlice => [xRotationMatrices3, CL.xSliceCoordsList(allCoordsList, xSlice)]),
    ...cubeData.values.map(ySlice => [yRotationMatrices3, CL.ySliceCoordsList(allCoordsList, ySlice)]),
    ...cubeData.values.map(zSlice => [zRotationMatrices3, CL.zSliceCoordsList(allCoordsList, zSlice)]),
    [xRotationMatrices3, CL.xSliceCoordsList(allCoordsList, -1).concat(CL.xSliceCoordsList(allCoordsList, 0))],
    [xRotationMatrices3, CL.xSliceCoordsList(allCoordsList, 0).concat(CL.xSliceCoordsList(allCoordsList, 1))],
    [yRotationMatrices3, CL.ySliceCoordsList(allCoordsList, -1).concat(CL.ySliceCoordsList(allCoordsList, 0))],
    [yRotationMatrices3, CL.ySliceCoordsList(allCoordsList, 0).concat(CL.ySliceCoordsList(allCoordsList, 1))],
    [zRotationMatrices3, CL.zSliceCoordsList(allCoordsList, -1).concat(CL.zSliceCoordsList(allCoordsList, 0))],
    [zRotationMatrices3, CL.zSliceCoordsList(allCoordsList, 0).concat(CL.zSliceCoordsList(allCoordsList, 1))],
    [xRotationMatrices3, allCoordsList],
    [yRotationMatrices3, allCoordsList],
    [zRotationMatrices3, allCoordsList],
  ]
  const nestedKvps = slices.map(makeKvpsForSlice)
  return new Map(U.flatten(nestedKvps))
}

export class Piece {
  constructor(id, x, y, z, cubeData) {
    this.id = id;
    this.x = x;
    this.y = y;
    this.z = z;
    this.faces = coordsToFaces(cubeData.vmin, cubeData.vmax, x, y, z);
    this.accTransform3 = R.Identity;
  }
}

export class Cube {
  constructor(pieces) {
    this.pieces = pieces;
  }
}

function makeSolvedCube(cubeSize) {
  const cubeData = new CubeData(cubeSize)
  const allCoordsList = cubeData.makeAllCoordsList();
  const pieces = allCoordsList.map(([x, y, z], index) => new Piece(index, x, y, z, cubeData));
  return new Cube(pieces);
}

export function makePerCubeSizeDataEntry(cubeSize) {
  const moveIdsToMoves = makeMoveIdsToMoves(cubeSize)
  const moves = Array.from(moveIdsToMoves.values())
  const solvedCube = makeSolvedCube(cubeSize)
  return [cubeSize, { moveIdsToMoves, moves, solvedCube }]
}

const PER_CUBE_SIZE_DATA = new Map([
  makePerCubeSizeDataEntry(2),
  makePerCubeSizeDataEntry(3),
  makePerCubeSizeDataEntry(4),
  makePerCubeSizeDataEntry(5),
  makePerCubeSizeDataEntry(6)
])

export function getSolvedCube(cubeSize) {
  const perCubeSizeData = PER_CUBE_SIZE_DATA.get(cubeSize)
  return perCubeSizeData.solvedCube
}

export function lookupMoveId(cubeSize, id) {
  const perCubeSizeData = PER_CUBE_SIZE_DATA.get(cubeSize)
  return perCubeSizeData.moveIdsToMoves.get(id)
}

export function getRandomMove(cubeSize) {
  const perCubeSizeData = PER_CUBE_SIZE_DATA.get(cubeSize)
  const randomIndex = Math.floor(Math.random() * perCubeSizeData.moves.length)
  return perCubeSizeData.moves[randomIndex]
}

export function removeRedundantMoves(moves) {
  for (; ;) {
    let removedSomething = false
    const indexes = U.range(moves.length)
    for (const index of indexes.slice(1)) {
      const move = moves[index]
      const previousMove = moves[index - 1]
      if (move.id === previousMove.oppositeMoveId) {
        moves.splice(index, 1)
        removedSomething = true
        break
      }
    }
    if (!removedSomething) break
  }
}

export function makeMoves(moves, initialCube) {
  return moves.reduce((cube, move) => move.makeMove(cube), initialCube.pieces);
}
