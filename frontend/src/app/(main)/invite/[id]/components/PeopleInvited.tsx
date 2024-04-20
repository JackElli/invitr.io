import { Invite } from "@/lib/services/InviteService"
import Person from "./Person";

type Props = {
    invite: Invite;
}

const PeopleInvited = ({ invite }: Props) => {
    return (
        <>
            <h1 className="text-xl font-bold mt-10">People invited</h1>
            <div className="flex flex-col gap-4 mt-2">
                {
                    invite.invitees.map((person, count) => {
                        return (
                            <Person person={person} invite={invite} />
                        )
                    })
                }
            </div>
        </>
    )
}

export default PeopleInvited;