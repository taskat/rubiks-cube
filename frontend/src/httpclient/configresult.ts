import { Side } from "../cube/side";
import { Error } from "../error/error";

export class ConfigResult {
    colorPalette: Map<Side, string[][]>;
    errors: Error[];
    constructor(data: any) {
        this.colorPalette = new Map();
        const stringToColors: Map<string, string[][]> = new Map(Object.entries(data.data.sides));
        stringToColors.forEach((colors, key) => {
            const side = Side[key as unknown as keyof typeof Side];
            this.colorPalette.set(side, colors);
        });
        if (data.errors) {
            this.errors = data.errors.map((error: any) => new Error(error));
        } else {
            this.errors = [];
        }
    }
}