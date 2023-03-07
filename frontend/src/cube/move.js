export default class Move {
    constructor(id, oppositeId, rotationMatrix3, coordsList, numTurns) {
        this.id = id;
        this.oppositeMoveId = oppositeId;
        this.rotationMatrix3 = rotationMatrix3;
        this.makeMove = rotatePieces(coordsList, rotationMatrix3);
        this.coordsList = coordsList;
        this.numTurns = numTurns;
    }
}

function rotatePieces(coordsList, rotationMatrix3) {
    return pieces => {
        pieces.forEach(piece => {
            if (piece.isInCoordsList(coordsList)) {
                piece.rotatePiece(rotationMatrix3)
            }
        });
        return pieces;
    };
}