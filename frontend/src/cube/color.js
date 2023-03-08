import * as THREE from "three";

const COLORS = {
    "b": new THREE.Color("blue"),
    "g": new THREE.Color("green"),
    "r": new THREE.Color("red"),
    "o": new THREE.Color("darkorange"),
    "y": new THREE.Color("yellow"),
    "w": new THREE.Color("ghostwhite")
};

function colors() {
    let map = new Map();
    const front = [
        ["y", "w", "y"],
        ["w", "y", "w"],
        ["y", "w", "y"]
    ];
    map.set("F", front.map(row => row.map(color => COLORS[color])));
    const back = [
        ["w", "y", "w"],
        ["y", "w", "y"],
        ["w", "y", "w"]
    ];
    map.set("B", back.map(row => row.map(color => COLORS[color])));
    const up = [
        ["b", "g", "b"],
        ["g", "b", "g"],
        ["b", "g", "b"]
    ];
    map.set("U", up.map(row => row.map(color => COLORS[color])));
    const down = [
        ["g", "b", "g"],
        ["b", "g", "b"],
        ["g", "b", "g"]
    ];
    map.set("D", down.map(row => row.map(color => COLORS[color])));
    const left = [
        ["r", "o", "r"],
        ["o", "r", "o"],
        ["r", "o", "r"]
    ];
    map.set("L", left.map(row => row.map(color => COLORS[color])));
    const right = [
        ["o", "r", "o"],
        ["r", "o", "r"],
        ["o", "r", "o"]
    ];
    map.set("R", right.map(row => row.map(color => COLORS[color])));
    return map;
}

const Colors = colors();
export default Colors;