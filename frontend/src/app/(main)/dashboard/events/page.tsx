import InviteService from "@/app/lib/services/InviteService";
import Invite from "../components/Invite";
import ErrorCard from "@/app/lib/components/ErrorCard";

export default async function EventsPage() {
    try {
        const invites = await InviteService.getByUser("123");
        return (
            <>
                <div className="flex flex-col gap-4 mt-2">
                    {
                        invites.length == 0 &&
                        <p>No invites found</p>
                    }
                    {
                        invites.length > 0 && invites.map((invite, count) => {
                            return (
                                <a href={`/invite/${invite.id}`}>
                                    <Invite _key={invite.date + count} invite={invite} />
                                </a>
                            )
                        })
                    }
                </div>
            </>
        )
    } catch (e) {
        return (
            <ErrorCard />
        )
    }
}