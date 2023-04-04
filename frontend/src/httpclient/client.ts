import axios from 'axios';
import { ConfigResult } from './configresult';

export default class Client {
    private baseUrl: string = "http://localhost:8080";
    constructor() {}
    
    public async postConfig(content: string): Promise<ConfigResult> {
        let body = {config: content};
        const response = await axios.post(this.baseUrl + "/config", JSON.stringify(body));
        console.log(response.data);
        return new ConfigResult(response.data);
    }
}