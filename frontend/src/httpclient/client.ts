import axios from 'axios';
import { Result } from './result';

export default class Client {
    private baseUrl: string = "http://localhost:8080";
    constructor() {}
    
    public async postConfig(content: string): Promise<Result> {
        let body = {config: content};
        return this.sendPostRequest("/config", body);
    }

    public async postAll(config: string, algo: string): Promise<Result> {
        let body = {config: config, algo: algo};
        return this.sendPostRequest("/all", body);
    }

    private async sendPostRequest(endpoint: string, body: any): Promise<Result> {
        const response = await axios.post(this.baseUrl + endpoint, JSON.stringify(body));
        console.log(response.data);
        return new Result(response.data);
    }
}