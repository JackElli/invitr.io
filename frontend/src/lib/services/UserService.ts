import NetworkService from "./NetworkService";

export type User = {
    id: string;
    firstName: string;
    lastName: string;
}

const PORT = "3200"
const SSR_IP = "users:" + PORT;
const CLIENT_IP = "localhost:" + PORT

class UserService {
    async getById(id: string): Promise<User> {
        return NetworkService.get(`http://${SSR_IP}/user/${id}`);
    }
}


export default new UserService();