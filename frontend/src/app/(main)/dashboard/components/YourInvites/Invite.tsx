import { getDate } from "@/app/lib/Date";
import { Invite } from "../../../../lib/services/InviteService";

type Props = {
    _key: string;
    invite: Invite;
}

const Invite_ = ({ _key, invite }: Props) => {
    return (
        <div key={_key} className="bg-white p-4 rounded-md border border-gray-200 shadow-sm hover:border-gray-300 cursor-pointer">
            <h1 className="font-bold text-md">{invite.title}</h1>
            <p><span className="font-bold">{getDate(invite.date)}</span> at {invite.location}</p>
        </div>
    )
}

export default Invite_;