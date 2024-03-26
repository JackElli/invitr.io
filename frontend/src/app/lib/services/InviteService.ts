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

const ip = "invites:3202";

class InviteService {
    async new(title: string, location: string, date: string, passphrase: string) {
        NetworkService.post(`http://localhost:3202/invites/invite`, {
            title: title,
            organiser: "123",
            location: location,
            date: date,
            passphrase: passphrase,
            invitees: ["jackellis"]
        })
    };

    async getByUser(userId: string): Promise<Invite[]> {
        return NetworkService.get(`http://${ip}/invites/user/${userId}`);
    }
}


export default new InviteService();