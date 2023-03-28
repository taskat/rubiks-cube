import axios from 'axios';

export default class Client {
    private baseUrl: string = "http://localhost:8080";
    constructor() {}
    
    public async postConfig(content: string): Promise<any> {
        const response = await axios.post(this.baseUrl + "/config", content);
        return response.data;
    }
}