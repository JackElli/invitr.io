import { redirect } from "next/navigation";

class NetworkService {
    async get(url: string, tag?: string): Promise<any> {
        return this.call(url, 'GET', tag);
    }

    async post(url: string, data: object, tag?: string): Promise<any> {
        return this.call(url, 'POST', tag, data);
    }

    async call(url: string, method: string, tag?: string, data?: object,) {
        const response = await fetch(url, {
            method: method,
            credentials: 'include',
            body: JSON.stringify(data),
            cache: 'no-store',
            next: {
                tags: [tag ?? '']
            }
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
