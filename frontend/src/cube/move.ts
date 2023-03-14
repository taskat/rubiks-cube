import math from "mathjs";
import CoordsList from "./coordslist";
import Piece from "./piece";

export default class Move {
    rotationMatrix3: math.Matrix;
    makeMove: (pieces: Piece[]) => Piece[];
    coordsList: CoordsList;
    numTurns: number;
    name: string;
    constructor(rotationMatrix3: math.Matrix, coordsList: CoordsList, numTurns: number, name: string) {
        this.rotationMatrix3 = rotationMatrix3;
        this.makeMove = rotatePieces(coordsList, rotationMatrix3);
        this.coordsList = coordsList;
        this.numTurns = numTurns;
        this.name = name;
    }

    toString(): string {
        return this.name;
    }
}

function rotatePieces(coordsList: CoordsList, rotationMatrix3: math.Matrix): (pieces: Piece[]) => Piece[] {
    return pieces => {
        pieces.forEach((piece: Piece) => {
            if (piece.isInCoordsList(coordsList)) {
                piece.rotatePiece(rotationMatrix3)
            }
        });
        return pieces;
    };
}