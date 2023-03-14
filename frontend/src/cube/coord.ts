export default class Coord {
    x: number;
    y: number;
    z: number;
    constructor(x: number, y: number, z: number) {
        this.x = x;
        this.y = y;
        this.z = z;
    }

    equal(other: Coord): boolean {
        return this.x === other.x && this.y === other.y && this.z === other.z;
    }
}