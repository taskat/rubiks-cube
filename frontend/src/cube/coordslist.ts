import Coord from "./coord";

export default class CoordsList {
    coords: Coord[];
    constructor(coords: Coord[]) {
        this.coords = coords;
    }

    concat(other: CoordsList): CoordsList {
        return new CoordsList(this.coords.concat(other.coords));
    }

    xSlice(xSlice: number): CoordsList {
        return new CoordsList(this.coords.filter((coord) => coord.x === xSlice));
    }

    ySlice(ySlice: number): CoordsList {
        return new CoordsList(this.coords.filter((coord) => coord.y === ySlice));
    }

    zSlice(zSlice: number): CoordsList {
        return new CoordsList(this.coords.filter((coord) => coord.z === zSlice));
    }
}