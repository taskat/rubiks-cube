import { rotateX, rotateY, rotateZ } from "./rotations";
import CubeData from "./cubedata";
import Piece from "./piece";
import Move from "./move";
import CoordsList from "./coordslist";
import MoveBuilder from "./movebuilder";
import * as THREE from "three";
import { Side } from "./side";
import { SideCoordCalculator } from "./sidecoordcalculator";

const COLORS: Map<String, THREE.Color> = new Map([
  ["b", new THREE.Color("blue")],
  ["g", new THREE.Color("green")],
  ["r", new THREE.Color("red")],
  ["o", new THREE.Color("darkorange")],
  ["y", new THREE.Color("yellow")],
  ["w", new THREE.Color("ghostwhite")],
  ["-", new THREE.Color(0x282828)]
]);

export default class Cube {
  cubeData: CubeData;
  pieces: Piece[];
  moves: Map<string, Move>;
  colorPalette: Map<Side, string[][]>
  constructor(cubeSize: number, colorPalette: Map<Side, string[][]>) {
    this.cubeData = new CubeData(cubeSize);
    const calculator = new SideCoordCalculator(this.cubeData);
    const allCoordsList = this.cubeData.makeAllCoordsList();
    this.pieces = allCoordsList.coords.map((coord, index) => new Piece(index, coord, calculator));
    this.moves = new Map(this.generateMoves().map(move => [move.name, move]));
    this.colorPalette = colorPalette;
  }

  closeTo(a: number, b: number) {
    return Math.abs(a - b) <= 1e-12
  }

  makeMoves(moves: Move[]) {
    this.pieces = moves.reduce((pieces, move) => move.makeMove(pieces), this.pieces);
  }

  generateMoves(): Move[] {
    const angles = [90, 180, 270];
    const xRotationMatrices3 = angles.map(rotateX);
    const yRotationMatrices3 = angles.map(rotateY);
    const zRotationMatrices3 = angles.map(rotateZ);
    const allCoordsList = this.cubeData.makeAllCoordsList();
    const xNames = ["L", "M", "R"];
    const yNames = ["D", "E", "U"];
    const zNames = ["B", "S", "F"];
    const builders: MoveBuilder[] = [
      ...this.cubeData.values.map((xSlice, index) => new MoveBuilder(xRotationMatrices3, allCoordsList.xSlice(xSlice), xNames[index])),
      ...this.cubeData.values.map((ySlice, index) => new MoveBuilder(yRotationMatrices3, allCoordsList.ySlice(ySlice), yNames[index])),
      ...this.cubeData.values.map((zSlice, index) => new MoveBuilder(zRotationMatrices3, allCoordsList.zSlice(zSlice), zNames[index])),
      new MoveBuilder(xRotationMatrices3, allCoordsList.xSlice(-1).concat(allCoordsList.xSlice(0)), "Lw"),
      new MoveBuilder(xRotationMatrices3, allCoordsList.xSlice(0).concat(allCoordsList.xSlice(1)), "Rw"),
      new MoveBuilder(yRotationMatrices3, allCoordsList.ySlice(-1).concat(allCoordsList.ySlice(0)), "Dw"),
      new MoveBuilder(yRotationMatrices3, allCoordsList.ySlice(0).concat(allCoordsList.ySlice(1)), "Uw"),
      new MoveBuilder(zRotationMatrices3, allCoordsList.zSlice(-1).concat(allCoordsList.zSlice(0)), "Bw"),
      new MoveBuilder(zRotationMatrices3, allCoordsList.zSlice(0).concat(allCoordsList.zSlice(1)), "Fw"),
      new MoveBuilder(xRotationMatrices3, allCoordsList, "x"),
      new MoveBuilder(yRotationMatrices3, allCoordsList, "y"),
      new MoveBuilder(zRotationMatrices3, allCoordsList, "z"),
    ];
    const toReverse = ["U", "R", "B", "Uw", "Rw", "Bw", "x", "y"];
    builders.forEach(builder => { 
      if (toReverse.includes(builder.name)) {
        builder.reverse();
      }
    });
    const nestedMoves = builders.map((moveBuilder: MoveBuilder) => moveBuilder.build());
    return flatten(nestedMoves);
  }

  getColor(piece: Piece, normalX: number, normalY: number, normalZ: number): THREE.Color {
    let side: Side = Side.None;
    if (this.closeTo(normalY, -1)) {
      side = Side.Down;
    } else if (this.closeTo(normalY, 1)) {
      side = Side.Up;
    } else if (this.closeTo(normalX, -1)) {
      side = Side.Left;
    } else if (this.closeTo(normalX, 1)) {
      side = Side.Right;
    } else if (this.closeTo(normalZ, 1)) {
      side = Side.Front;
    } else if (this.closeTo(normalZ, -1)) {
      side = Side.Back;
    }
    const sideCoord = piece.getSideCoord(side);
    if (sideCoord) {
      let color = (this.colorPalette.get(side) as string[][])[sideCoord.i][sideCoord.j]
      let color2 = COLORS.get(color);
      if (!color2) {
        console.log(sideCoord);
        console.log(color);
        console.log(piece.coord);
      }
      return COLORS.get(color) as THREE.Color;
    } 
    return COLORS.get("-") as THREE.Color;
  }

  getPieces(coordsList: CoordsList): Piece[] {
    return coordsList.coords.map(coord =>
      this.pieces.find(piece => piece.coord.equal(coord))).filter(piece => piece !== undefined) as Piece[];
  }

  setColorPalette(colorPalette: Map<Side, string[][]>) {
    this.colorPalette = colorPalette;
  }
}

function flatten(xss: Move[][]): Move[] {
  return ([] as Move[]).concat(...xss)
}
