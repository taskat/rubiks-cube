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

function colorsChecker(): Map<Side, THREE.Color[][]> {
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

export const ColorsChecker: Map<Side, THREE.Color[][]> = colorsChecker();

function colorsSolved(): Map<Side, THREE.Color[][]> {
    let map = new Map<Side, THREE.Color[][]>();
    const front = [
        ["b", "b", "b"],
        ["b", "b", "b"],
        ["b", "b", "b"]
    ];
    map.set(Side.Front, front.map(row => row.map(color => COLORS.get(color) as THREE.Color)));
    const back = [
        ["g", "g", "g"],
        ["g", "g", "g"],
        ["g", "g", "g"]
    ];
    map.set(Side.Back, back.map(row => row.map(color => COLORS.get(color) as THREE.Color)));
    const up = [
        ["y", "y", "y"],
        ["y", "y", "y"],
        ["y", "y", "y"]
    ];
    map.set(Side.Up, up.map(row => row.map(color => COLORS.get(color) as THREE.Color)));
    const down = [
        ["w", "w", "w"],
        ["w", "w", "w"],
        ["w", "w", "w"]
    ];
    map.set(Side.Down, down.map(row => row.map(color => COLORS.get(color) as THREE.Color)));
    const left = [
        ["o", "o", "o"],
        ["o", "o", "o"],
        ["o", "o", "o"]
    ];
    map.set(Side.Left, left.map(row => row.map(color => COLORS.get(color) as THREE.Color)));
    const right = [
        ["r", "r", "r"],
        ["r", "r", "r"],
        ["r", "r", "r"]
    ];
    map.set(Side.Right, right.map(row => row.map(color => COLORS.get(color) as THREE.Color)));
    return map;
}

export const ColorsSolved: Map<Side, THREE.Color[][]> = colorsSolved();
