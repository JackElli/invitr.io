import { getDate, getDateRelative } from "@/app/lib/Date";
import { Invite } from "../../../../lib/services/InviteService";
import { USER } from "@/app/(main)/layout";

type Props = {
    _key: string;
    invite: Invite;
}

const Invite = ({ _key, invite }: Props) => {
    return (
        <div key={_key} className="flex justify-between items-center bg-white p-4 rounded-md border border-gray-200 shadow-sm hover:border-gray-300 cursor-pointer">
            <div>
                <div className="flex gap-2 items-center">
                    <h1 className="font-bold text-lg">{invite.title}</h1>
                    {invite.organiser == USER &&
                        <div className='h-3 w-3 bg-yellow-300 rounded-sm' title="You are the organiser of this event"></div>}
                </div>

                <p><span className="font-bold">{getDate(invite.date)}</span> at {invite.location}</p>
            </div>

            <p className="text-sm bg-gray-200 inline-block px-2 rounded-sm">{getDateRelative(invite.date)}</p>
        </div>
    )
}

export default Invite;