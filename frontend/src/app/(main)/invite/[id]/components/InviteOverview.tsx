import { getDate, getDateRelative } from "@/app/lib/Date";
import { Invite } from "@/app/lib/services/InviteService";

type Props = {
    invite: Invite;
}

const InviteOverview = ({ invite }: Props) => {
    return (
        <div>
            <h1 className="font-bold text-5xl">{invite.title}</h1>
            <h1 className="font-semibold text-2xl mt-2">@ {invite.location}</h1>
            <h1 className="font-semibold text-xl mt-2">{getDate(invite.date)}</h1>
            <p className="text-sm bg-gray-200 inline-block px-2 rounded-sm animate-pulse">{getDateRelative(invite.date)}</p>
        </div>
    )
}

export default InviteOverview;