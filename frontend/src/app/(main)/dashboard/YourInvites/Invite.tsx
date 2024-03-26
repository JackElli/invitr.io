import { Invite } from "../../../lib/services/InviteService";

type Props = {
    key: string;
    invite: Invite;
}

const Invite_ = ({ key, invite }: Props) => {
    return (
        <div className="bg-white p-4 rounded-md border border-gray-200 shadow-sm hover:border-gray-300 cursor-pointer">
            <h1 className="font-bold text-md">{invite.title}</h1>
            <p key={key}><span className="font-bold">{invite.date}</span> at {invite.location}</p>
        </div>
    )
}

export default Invite_;