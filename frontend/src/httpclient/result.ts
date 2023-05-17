import { Side } from "../cube/side";
import { Error } from "../error/error";

export class Result {
    colorPalette: Map<Side, string[][]> = this.createEmptyPalette();
    errors: Error[] = [];
    turns: string[] = [];
    constructor(data: any) {
        if (data.data) {
            const stringToColors: Map<string, string[][]> = new Map(Object.entries(data.state.sides));
            stringToColors.forEach((colors, key) => {
                const side = Side[key as unknown as keyof typeof Side];
                this.colorPalette.set(side, colors);
            });
        }
        if (data.errors) {
            this.errors = data.errors.map((error: any) => new Error(error));
        }
        if (data.turns) {
            this.turns = data.turns;
        }
    }

    createEmptyPalette(): Map<Side, string[][]> {
        const sides = [Side.Front, Side.Back, Side.Left, Side.Right, Side.Up, Side.Down];
        const emptySide = Array(3).fill(Array(3).fill("-"));
        return new Map(sides.map(side => [side, emptySide]));
    }
}