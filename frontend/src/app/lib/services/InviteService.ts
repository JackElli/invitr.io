import NetworkService from "./NetworkService";

export type Person = {
    name: string;
    is_going?: boolean;
}

export type Invite = {
    id: string;
    title: string;
    organiser: string;
    location: string;
    notes: string;
    date: string;
    passphrase: string;
    invitees: Person[];
}

const PORT = "3202"
const SSR_IP = "invites:" + PORT;
const CLIENT_IP = "localhost:" + PORT

class InviteService {
    async new(title: string, location: string, date: string, passphrase: string, people: string) {
        const invitees = people.split(",").map((p) => {
            return {
                "name": p.trim(),
            }
        });

        NetworkService.post(`http://${CLIENT_IP}/invites/invite`, {
            title: title,
            organiser: "123",
            location: location,
            date: date,
            passphrase: passphrase,
            invitees: invitees
        });
    }

    async addNotes(inviteId: string, notes: string) {
        NetworkService.post(`http://${CLIENT_IP}/invites/invite/${inviteId}/note`, {
            notes: notes
        });
    }

    async getById(id: string): Promise<Invite> {
        return NetworkService.get(`http://${SSR_IP}/invites/invite/${id}`);
    }

    async getByUser(userId: string): Promise<{ finished: Invite[], ongoing: Invite[] }> {
        return NetworkService.get(`http://${SSR_IP}/invites/user/${userId}`);
    }

    async isUserGoing(inviteId: string, user: string, tag?: string): Promise<boolean | null> {
        return NetworkService.get(`http://${SSR_IP}/invites/invite/${inviteId}/user/${user}`, tag);
    }

    async respondToEvent(response: boolean, inviteId: string, user: string) {
        return NetworkService.post(`http://${CLIENT_IP}/invites/invite/${inviteId}/user/${user}`, {
            going: response
        });
    }
}


export default new InviteService();