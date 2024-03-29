'use client'
import Button from "@/app/lib/components/Button"
import InviteService, { Invite } from "@/app/lib/services/InviteService";

type Props = {
    invite: Invite;
    // TODO get rid of the username
    user: string;
}
const RejectButton = ({ invite, user }: Props) => {

    const reject = () => {
        InviteService.respondToEvent(false, invite.id, user);
    }

    return <Button className="bg-red-200" onClick={reject}>Reject</Button>
}

export default RejectButton;