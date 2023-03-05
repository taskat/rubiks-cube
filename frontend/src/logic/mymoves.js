import {Cube, lookupMoveId} from "./index"

const CUBE = new Cube(3);
const MOVES = CUBE.moves;

export const getMoves = cubeSize => {
    let moves = []
    for (let i = 0; i < 3; i++) {
        for (let j = 0; j < 4; j++) {
            moves.push(middlePrime())
            moves.push(upPrime())
        }
        moves.push(xPrime())
        moves.push(zPrime())
    }
    return moves;
}

const left = () => {
    return MOVES[0]
}

const left2 = () => {
    return MOVES[1]
}

const leftPrime = () => {
    return MOVES[2]
}

const middle = () => {
    return MOVES[3]
}

const middle2 = () => {
    return MOVES[4]
}

const middlePrime = () => {
    return MOVES[5]
}

const rightPrime = () => {
    return MOVES[6]
}

const right2 = () => {
    return MOVES[7]
}

const right = () => {
    return MOVES[8]
}

const down = () => {
    return MOVES[9]
}

const down2 = () => {
    return MOVES[10]
}

const downPrime = () => {
    return MOVES[11]
}

const equator = () => {
    return MOVES[12]
}

const equator2 = () => {
    return MOVES[13]
}

const equatorPrime = () => {
    return MOVES[14]
}

const upPrime = () => {
    return MOVES[15]
}

const up2 = () => {
    return MOVES[16]
}

const up = () => {
    return MOVES[17]
}

const backPrime = () => {
    return MOVES[18]
}

const back2 = () => {
    return MOVES[19]
}

const back = () => {
    return MOVES[20]
}

const side = () => {
    return MOVES[21]
}

const side2 = () => {
    return MOVES[22]
}

const sidePrime = () => {
    return MOVES[23]
}

const front = () => {
    return MOVES[24]
}

const front2 = () => {
    return MOVES[25]
}

const frontPrime = () => {
    return MOVES[26]
}

const leftWide = () => {
    return MOVES[27]
}

const leftWide2 = () => {
    return MOVES[28]
}

const lefttWidePrime = () => {
    return MOVES[29]
}

const rightWidePrime = () => {
    return MOVES[30]
}

const rightWide2 = () => {
    return MOVES[31]
}

const rightWide = () => {
    return MOVES[32]
}

const downWide = () => {
    return MOVES[33]
}

const downWide2 = () => {
    return MOVES[34]
}

const downWidePrime = () => {
    return MOVES[35]
}

const upWidePrime = () => {
    return MOVES[36]
}

const upWide2 = () => {
    return MOVES[37]
}

const upWide = () => {
    return MOVES[38]
}

const backWidePrime = () => {
    return MOVES[39]
}

const backWide2 = () => {
    return MOVES[40]
}

const backWide = () => {
    return MOVES[41]
}

const frontWide = () => {
    return MOVES[42]
}

const frontWide2 = () => {
    return MOVES[43]
}

const frontWidePrime = () => {
    return MOVES[44]
}

const xPrime = () => {
    return MOVES[45]
}

const x2 = () => {
    return MOVES[46]
}

const x = () => {
    return MOVES[47]
}

const yPrime = () => {
    return MOVES[48]
}

const y2 = () => {
    return MOVES[49]
}

const y = () => {
    return MOVES[50]
}

const z = () => {
    return MOVES[51]
}

const z2 = () => {
    return MOVES[52]
}

const zPrime = () => {
    return MOVES[53]
}

export const reverseMoves = moves => {
    return moves
      .map(move => move.oppositeMoveId)
      .map(id => CUBE.lookupMoveId(id))
      .reverse()
}