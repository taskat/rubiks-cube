import axios from 'axios';
import { ConfigResult } from './configresult';
import { AlgoResult } from './algoresult';

export default class Client {
    private baseUrl: string = "http://localhost:8080";
    constructor() {}
    
    public async postConfig(content: string): Promise<ConfigResult> {
        let body = {config: content};
        const response = await axios.post(this.baseUrl + "/config", JSON.stringify(body));
        console.log(response.data);
        return new ConfigResult(response.data);
    }

    public async postAlgo(content: string): Promise<AlgoResult> {
        let body = {algo: content};
        const response = await axios.post(this.baseUrl + "/algo", JSON.stringify(body));
        console.log(response.data);
        return new AlgoResult(response.data);
    }
}