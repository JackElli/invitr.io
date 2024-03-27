import InviteService from "@/app/lib/services/InviteService";

import ErrorCard from "@/app/lib/components/ErrorCard";
import Events from "./Events";

export default async function EventsPage() {
    try {
        const invites = await InviteService.getByUser("123");
        return (
            <div>
                <div>
                    <h1 className="font-semibold text-sm text-gray-500">Upcoming events</h1>
                    <Events invites={invites.ongoing} />
                </div>

                <div className="mt-4">
                    <h1 className="font-semibold text-sm text-gray-500">Past events</h1>
                    <Events invites={invites.finished} />
                </div>

            </div>
        )
    } catch (e) {
        return (
            <ErrorCard />
        )
    }
}