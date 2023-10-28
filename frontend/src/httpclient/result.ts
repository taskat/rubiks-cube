import { Side } from "../cube/side";
import { Error } from "../error/error";

export class Result {
    colorPalette: Map<Side, string[][]> = this.createEmptyPalette();
    errors: Error[] = [];
    turns: Map<string, string[]> = new Map();
    steps: string[] = [];
    size: number = 0;
    constructor(data: any) {
        if (data.state) {
            const stringToColors: Map<string, string[][]> = new Map(Object.entries(data.state.sides));
            stringToColors.forEach((colors, key) => {
                const side = Side[key as unknown as keyof typeof Side];
                this.colorPalette.set(side, colors);
            });
            this.size = data.state.size;
        }
        if (data.errors) {
            this.errors = data.errors.map((error: any) => new Error(error));
        }
        if (data.turns) {
            this.turns = new Map(Object.entries(data.turns));
        }
        if (data.steps) {
            this.steps = data.steps;
        }
        if (data.steps && data.turns) {
            let orderedTurns = new Map();
            this.steps.forEach((step: string) => {
                orderedTurns.set(step, this.turns.get(step));
            });
            this.turns = orderedTurns;
        }
    }

    createEmptyPalette(): Map<Side, string[][]> {
        const sides = [Side.Front, Side.Back, Side.Left, Side.Right, Side.Up, Side.Down];
        const emptySide = Array(3).fill(Array(3).fill("-"));
        return new Map(sides.map(side => [side, emptySide]));
    }
}