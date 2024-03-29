'use server'
import ErrorCard from "@/app/lib/components/ErrorCard";
import InviteService, { Invite } from "@/app/lib/services/InviteService"
import AcceptButton from "./AcceptButton";
import RejectButton from "./RejectButton";
import { revalidateTag } from "next/cache";

// TODO also need to check if user has actually been invited
// VERY IMPORTANT!!!
// for now, we'll just assume the user has been invited
export async function ChangeGoing({ invite, searchParams }: { invite: Invite, searchParams: { [key: string]: string | string[] | undefined } }) {
    try {

        // Do this again here for security (better to check key twice)
        const key = searchParams['key'] as string;
        const username = await InviteService.getUserFromKey(invite.id, key)

        if (!username) {
            return <p>You are not allowed to be here :)</p>
        }

        const isGoing = await InviteService.isUserGoing(invite.id, username, "going");

        if (isGoing == true) {
            return <p className="bg-green-200 px-2 py-1 rounded-xl shadow-md">You accepted</p>
        }

        if (isGoing == false) {
            return <p className="bg-red-200 px-2 py-1 rounded-xl shadow-md">You declined</p>
        }

        return (
            <div>
                <p className="text-sm text-gray-600 text-center">RSVP to this event</p>

                <div className="bg-white p-4 border border-gray-200 rounded-md shadow-lg">
                    <div className="flex gap-2 justify-center">
                        <form action={action}>
                            <AcceptButton invite={invite} username={username} />
                        </form>

                        <form action={action}>
                            <RejectButton invite={invite} username={username} />
                        </form>
                    </div>
                </div >
            </div>
        )
    } catch (e) {
        return <ErrorCard />
    }
}

export default async function action() {
    revalidateTag('going')
}
