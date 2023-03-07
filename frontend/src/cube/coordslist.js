export default class CoordsList {
    constructor(coords) {
        this.coords = coords;
    }

    concat(other) {
        return new CoordsList(this.coords.concat(other.coords));
    }

    xSlice(xSlice) {
        return new CoordsList(this.coords.filter((coord) => coord.x === xSlice));
    }

    ySlice(ySlice) {
        return new CoordsList(this.coords.filter((coord) => coord.y === ySlice));
    }

    zSlice(zSlice) {
        return new CoordsList(this.coords.filter((coord) => coord.z === zSlice));
    }
}