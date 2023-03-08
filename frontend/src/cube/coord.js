export default class Coord {
    constructor(x, y, z) {
        this.x = x;
        this.y = y;
        this.z = z;
    }

    equal(other) {
        return this.x === other.x && this.y === other.y && this.z === other.z;
    }
}