import Coord from "./coord";
import CubeData from "./cubedata";
import { Side } from "./side";
import SideCoord from "./sidecoord";

export class SideCoordCalculator {
    vmin: number = 0;
    vmax: number = 0;
    isEven: boolean = false;
    sideCoords: SideCoord[] = [];

    constructor(cubedata: CubeData) {
        this.vmin = cubedata.vmin;
        this.vmax = cubedata.vmax;
        this.isEven = cubedata.isEvenSizedCube;
        this.sideCoords = [];
    }

    calculate(coord: Coord): SideCoord[] {
        this.sideCoords = [];
        if (coord.y === this.vmin) {
            this.getDown(coord);
        }
        if (coord.y === this.vmax) {
            this.getUp(coord);
        }
        if (coord.x === this.vmin) {
            this.getLeft(coord);
        }
        if (coord.x === this.vmax) {
            this.getRight(coord);
        }
        if (coord.z === this.vmin) {
            this.getBack(coord);
        }
        if (coord.z === this.vmax) {
            this.getFront(coord);
        }
        return this.sideCoords;
    }

    getDown(coord: Coord): void {
        let z: number;
        let x: number;
        if (this.isEven) {
            z = coord.z < 0 ? this.vmax - coord.z - 1 : this.vmax - coord.z;
            x = coord.x < 0 ? this.vmax + coord.x : this.vmax + coord.x - 1;
        } else {
            z = this.vmax - coord.z;
            x = this.vmax + coord.x;
        }
        this.sideCoords.push(new SideCoord(Side.Down, z, x));
    }

    getUp(coord: Coord): void {
        let z: number;
        let x: number;
        if (this.isEven) {
            z = coord.z < 0 ? this.vmax + coord.z : this.vmax + coord.z - 1;
            x = coord.x < 0 ? this.vmax + coord.x : this.vmax + coord.x - 1;
        } else {
            z = this.vmax + coord.z;
            x = this.vmax + coord.x;
        }
        this.sideCoords.push(new SideCoord(Side.Up, z, x));
    }

    getLeft(coord: Coord): void {
        let y: number;
        let z: number;
        if (this.isEven) {
            y = coord.y < 0 ? this.vmax - coord.y - 1 : this.vmax - coord.y;
            z = coord.z < 0 ? this.vmax + coord.z : this.vmax + coord.z - 1;
        } else {
            y = this.vmax - coord.y;
            z = this.vmax + coord.z;
        }
        this.sideCoords.push(new SideCoord(Side.Left, y, z));
    }

    getRight(coord: Coord): void {
        let y: number;
        let z: number;
        if (this.isEven) {
            y = coord.y < 0 ? this.vmax - coord.y - 1 : this.vmax - coord.y;
            z = coord.z < 0 ? this.vmax - coord.z - 1 : this.vmax - coord.z;
        } else {
            y = this.vmax - coord.y;
            z = this.vmax - coord.z;
        }
        this.sideCoords.push(new SideCoord(Side.Right, y, z));
    }

    getBack(coord: Coord): void {
        let y: number;
        let x: number;
        if (this.isEven) {
            y = coord.y < 0 ? this.vmax - coord.y - 1 : this.vmax - coord.y;
            x = coord.x < 0 ? this.vmax - coord.x - 1 : this.vmax - coord.x;
        } else {
            y = this.vmax - coord.y;
            x = this.vmax - coord.x;
        }
        this.sideCoords.push(new SideCoord(Side.Back, y, x));
    }

    getFront(coord: Coord): void {
        let y: number;
        let x: number;
        if (this.isEven) {
            y = coord.y < 0 ? this.vmax - coord.y - 1 : this.vmax - coord.y;
            x = coord.x < 0 ? this.vmax + coord.x : this.vmax + coord.x - 1;
        } else {
            y = this.vmax - coord.y;
            x = this.vmax + coord.x;
        }
        this.sideCoords.push(new SideCoord(Side.Front, y, x));
    }
}