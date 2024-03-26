import { redirect } from "next/navigation";

class NetworkService {
    async get(url: string): Promise<any> {
        return this.call(url, 'GET');
    }

    async post(url: string, data: object): Promise<any> {
        return this.call(url, 'POST', data);
    }

    async call(url: string, method: string, data?: object) {
        const response = await fetch(url, {
            method: method,
            credentials: 'include',
            body: JSON.stringify(data),
            cache: 'no-store'
        });

        if (!response.ok) {
            if (response.status == 401) {
                throw redirect('/login');
            }
            throw new Error('Oops, something wrong has happened.');
        }

        return await response.json();
    }
}

export default new NetworkService();
