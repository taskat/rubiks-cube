import { Side } from "./side";

export default class SideCoord {
    side: Side;
    i: number;
    j: number;
    constructor(side: Side, i: number, j: number) {
        this.side = side;
        this.i = i;
        this.j = j;
    }
}