import InviteService from "@/lib/services/InviteService";
import ErrorCard from "@/lib/components/ErrorCard";
import Events from "./Events";
import { USER } from "@/app/page";

export default async function EventsPage() {
    try {
        const invites = await InviteService.getByUser(USER);

        if (!invites.finished && !invites.ongoing) {
            return <p>No events found</p>
        }

        return (
            <div>
                {invites.ongoing &&
                    <div>
                        <h1 className="font-semibold text-sm text-gray-500">Upcoming events</h1>
                        <Events invites={invites.ongoing} />
                    </div>
                }

                {invites.finished &&
                    <div className="mt-4">
                        <h1 className="font-semibold text-sm text-gray-500">Past events</h1>
                        <Events invites={invites.finished} />
                    </div>
                }

            </div>
        )
    } catch (e) {
        return (
            <ErrorCard />
        )
    }
}