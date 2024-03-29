import { Invite, Person } from "@/app/lib/services/InviteService";
import IsGoing from "./IsGoing";
import UserService, { User } from "@/app/lib/services/UserService";

type Props = {
    invite: Invite;
    person: Person;
}

export default async function Person({ invite, person }: Props) {

    let user: User | undefined;

    try {
        user = await UserService.getById(person.id);
    } catch (e) {
        // WE NEED TO SORT ERROR MESSAGES OUT!!!
        user = undefined
    }

    return (
        <div className="flex items-center justify-between bg-white px-4 py-3 rounded-md border border-gray-200 shadow-sm hover:border-gray-300">
            <div className="flex gap-2 items-center">
                <div className={`w-6 h-6 ${user ? 'rounded-full' : ''} bg-gray-200`}></div>
                {user &&
                    < p className="text-lg font-bold">{user.firstName} {user.lastName}  {person.id == invite.organiser ? <span className="font-medium">(organiser)</span> : ''}</p>
                }

                {!user &&
                    <p className="text-lg font-bold">{person.id}</p>
                }

            </div>

            {
                person.id != invite.organiser &&
                <IsGoing isGoing={person.is_going} />
            }

        </div >
    )

}