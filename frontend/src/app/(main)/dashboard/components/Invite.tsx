import { getDate, getDateRelative } from "@/lib/Date";
import InviteService, { Invite } from "../../../../lib/services/InviteService";
import { USER } from "@/app/page";

type Props = {
    _key: string;
    invite: Invite;
}

async function Invite({ _key, invite }: Props) {
    // Careful here!
    const key = await InviteService.getOrganiserKey(invite.id, USER);

    return (
        <div key={_key}>
            <a href={`/invite/${invite.id}?key=${key}`}>
                <div className="flex justify-between items-center bg-white p-4 rounded-md border border-gray-200 shadow-sm hover:border-gray-300 cursor-pointer">
                    <div>
                        <div className="flex gap-2 items-center">
                            <h1 className="font-bold text-2xl">{invite.title}</h1>
                            {invite.organiser == USER &&
                                <div className='h-3 w-3 bg-gradient-to-r from-[#e1f9fc] to-gray-100 border border-yellow-600 rounded-sm mt-1' title="You are the organiser of this event"></div>}
                        </div>
                        <div className="flex gap-1 items-center mt-2">
                            <p className="text-sm bg-gray-200 inline-block px-2 rounded-sm">{getDateRelative(invite.date)}</p>
                            <p><span className="font-bold"> - {getDate(invite.date)}</span> @ {invite.location}</p>
                        </div>
                    </div>

                    <p className="text-gray-500"><span className="font-bold">?</span> going</p>
                </div>
            </a>
        </div>
    )
}

export default Invite;