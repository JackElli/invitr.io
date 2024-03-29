'use client'
import Button from "@/app/lib/components/Button"
import InviteService, { Invite } from "@/app/lib/services/InviteService";

type Props = {
    invite: Invite;
    // TODO get rid of the username
    user: string;
}
const AcceptButton = ({ invite, user }: Props) => {

    const accept = () => {
        InviteService.respondToEvent(true, invite.id, user);
    }

    return <Button className="bg-green-200" onClick={accept}>Accept</Button>
}

export default AcceptButton;