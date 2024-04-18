import { getDate, getDateRelative } from "@/lib/Date";
import { Invite } from "@/lib/services/InviteService";

type Props = {
    invite: Invite;
}

const InviteOverview = ({ invite }: Props) => {
    return (
        <div>
            <h1 className="font-bold text-5xl">{invite.title}</h1>
            <h1 className="font-semibold text-2xl mt-2">@ {invite.location}</h1>
            <div className="flex gap-1 items-center mt-2">
                <p className="text-sm bg-gray-200 inline-block px-2 py-1 rounded-sm animate-pulse">{getDateRelative(invite.date)}</p>
                <h1 className="font-semibold text-xl">- {getDate(invite.date)}</h1>
            </div>
        </div>
    )
}

export default InviteOverview;