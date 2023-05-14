import { Error } from "../error/error";

export class AlgoResult {
    errors: Error[] = [];
    constructor(data: any) {
        if (data.errors) {
            this.errors = data.errors.map((error: any) => new Error(error));
        }
    }
}