import { Level } from "./level";
import { Position } from "./pos";

export class Error {
    text: string;
    file: string;
    pos: Position;
    level: Level;
    constructor(data : any) {
        this.text = data.text;
        this.file = data.file;
        this.pos = data.pos;
        this.level = data.level;
    }
}