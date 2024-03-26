import NetworkService from "./NetworkService";

export type Invite = {
    id: string;
    title: string;
    organiser: string;
    location: string;
    date: string;
    passphrase: string;
    invitees: string[];
}

const PORT = "3202"
const SSR_IP = "invites:" + PORT;
const CLIENT_IP = "localhost:" + PORT

class InviteService {
    async new(title: string, location: string, date: string, passphrase: string) {
        NetworkService.post(`http://${CLIENT_IP}/invites/invite`, {
            title: title,
            organiser: "123",
            location: location,
            date: date,
            passphrase: passphrase,
            invitees: ["jackellis"]
        });
    }

    async getById(id: string): Promise<Invite> {
        return NetworkService.get(`http://${SSR_IP}/invites/invite/${id}`);
    }

    async getByUser(userId: string): Promise<Invite[]> {
        return NetworkService.get(`http://${SSR_IP}/invites/user/${userId}`);
    }
}


export default new InviteService();