import * as THREE from "three";
import { Side } from "./side";

const COLORS: Map<String, THREE.Color> = new Map([
    ["b", new THREE.Color("blue")],
    ["g", new THREE.Color("green")],
    ["r", new THREE.Color("red")],
    ["o", new THREE.Color("darkorange")],
    ["y", new THREE.Color("yellow")],
    ["w", new THREE.Color("ghostwhite")]
]);

function colors(): Map<Side, THREE.Color[][]> {
    let map = new Map<Side, THREE.Color[][]>();
    const front = [
        ["y", "w", "y"],
        ["w", "y", "w"],
        ["y", "w", "y"]
    ];
    map.set(Side.Front, front.map(row => row.map(color => COLORS.get(color) as THREE.Color)));
    const back = [
        ["w", "y", "w"],
        ["y", "w", "y"],
        ["w", "y", "w"]
    ];
    map.set(Side.Back, back.map(row => row.map(color => COLORS.get(color) as THREE.Color)));
    const up = [
        ["b", "g", "b"],
        ["g", "b", "g"],
        ["b", "g", "b"]
    ];
    map.set(Side.Up, up.map(row => row.map(color => COLORS.get(color) as THREE.Color)));
    const down = [
        ["g", "b", "g"],
        ["b", "g", "b"],
        ["g", "b", "g"]
    ];
    map.set(Side.Down, down.map(row => row.map(color => COLORS.get(color) as THREE.Color)));
    const left = [
        ["r", "o", "r"],
        ["o", "r", "o"],
        ["r", "o", "r"]
    ];
    map.set(Side.Left, left.map(row => row.map(color => COLORS.get(color) as THREE.Color)));
    const right = [
        ["o", "r", "o"],
        ["r", "o", "r"],
        ["o", "r", "o"]
    ];
    map.set(Side.Right, right.map(row => row.map(color => COLORS.get(color) as THREE.Color)));
    return map;
}

const Colors: Map<Side, THREE.Color[][]> = colors();
export default Colors;