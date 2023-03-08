import math from "mathjs";
import CoordsList from "./coordslist";
import Piece from "./piece";

export default class Move {
    id: number;
    oppositeMoveId: number;
    rotationMatrix3: math.Matrix;
    makeMove: (pieces: any) => any;
    coordsList: CoordsList;
    numTurns: number;
    constructor(id: number, oppositeId: number, rotationMatrix3: math.Matrix, coordsList: CoordsList, numTurns: number) {
        this.id = id;
        this.oppositeMoveId = oppositeId;
        this.rotationMatrix3 = rotationMatrix3;
        this.makeMove = rotatePieces(coordsList, rotationMatrix3);
        this.coordsList = coordsList;
        this.numTurns = numTurns;
    }
}

function rotatePieces(coordsList: CoordsList, rotationMatrix3: math.Matrix): (pieces: any) => any {
    return pieces => {
        pieces.forEach((piece: Piece) => {
            if (piece.isInCoordsList(coordsList)) {
                piece.rotatePiece(rotationMatrix3)
            }
        });
        return pieces;
    };
}