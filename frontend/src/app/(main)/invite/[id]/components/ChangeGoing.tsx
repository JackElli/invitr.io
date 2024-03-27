'use server'
import ErrorCard from "@/app/lib/components/ErrorCard";
import InviteService, { Invite } from "@/app/lib/services/InviteService"
import AcceptButton from "./AcceptButton";
import RejectButton from "./RejectButton";
import { revalidateTag } from "next/cache";

type Props = {
    invite: Invite
}

// TODO get this from URL
const username = "jack"

// TODO also need to check if user has actually been invited
// VERY IMPORTANT!!!
// for now, we'll just assume the user has been invited
export async function ChangeGoing({ invite }: Props) {
    try {
        const isGoing = await InviteService.isUserGoing(invite.id, username, "going");

        if (isGoing == true) {
            return <p className="bg-green-200 px-2 py-1 rounded-xl">Accepted</p>
        }

        if (isGoing == false) {
            return <p className="bg-red-200 px-2 py-1 rounded-xl">Declined</p>
        }

        return (
            <div>
                <p className="bg-gray-200 px-2 py-1 rounded-xl">You have not responded</p>
                <div className="flex gap-2 mt-2 justify-center">
                    <form action={action}>
                        <AcceptButton invite={invite} username={username} />
                    </form>

                    <form action={action}>
                        <RejectButton invite={invite} username={username} />
                    </form>
                </div>
            </div >

        )
    } catch (e) {
        return <ErrorCard />
    }
}

export default async function action() {
    revalidateTag('going')
}
