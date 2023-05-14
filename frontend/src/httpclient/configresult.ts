import { Side } from "../cube/side";
import { Error } from "../error/error";

export class ConfigResult {
    colorPalette: Map<Side, string[][]> = new Map();
    errors: Error[] = [];
    constructor(data: any) {
        this.createEmptyPalette();
        if (data.data) {
            const stringToColors: Map<string, string[][]> = new Map(Object.entries(data.data.sides));
            stringToColors.forEach((colors, key) => {
                const side = Side[key as unknown as keyof typeof Side];
                this.colorPalette.set(side, colors);
            });
        }
        if (data.errors) {
            this.errors = data.errors.map((error: any) => new Error(error));
        }
    }

    createEmptyPalette() {
        const sides = [Side.Front, Side.Back, Side.Left, Side.Right, Side.Up, Side.Down];
        const emptySide = Array(3).fill(Array(3).fill("-"));
        this.colorPalette = new Map(sides.map(side => [side, emptySide]));
    }
}