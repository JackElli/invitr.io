import ErrorCard from "@/app/lib/components/ErrorCard";
import InviteService from "@/app/lib/services/InviteService";
import PeopleInvited from "./components/PeopleInvited";
import Notes from "./components/Notes";
import InviteOverview from "./components/InviteOverview";
import DeleteButton from "./components/DeleteButton";
import { USER } from "../../layout";
import { ChangeGoing } from "./components/ChangeGoing";


export default async function InvitePage({ params }: { params: { id: string } }) {
    try {
        const invite = await InviteService.getById(params.id);

        // event is NOT owned by the user that's logged in
        if (invite.organiser != USER) {
            return (
                <div>
                    <div className="border-b border-b-zinc-200 pb-4">
                        <div className="flex justify-between items-center">
                            <InviteOverview invite={invite} />
                            <ChangeGoing invite={invite} />
                        </div>
                    </div>

                    <Notes invite={invite} />
                </div>
            )
        }

        // event IS owned by the user that's logged in
        return (
            <div>
                <div className="border-b border-b-zinc-200 pb-4">
                    <div className="flex justify-between items-center">
                        <InviteOverview invite={invite} />
                        <DeleteButton disabled>Delete event</DeleteButton>
                    </div>
                </div>

                <Notes invite={invite} editable />

                <PeopleInvited invite={invite} />
            </div>
        )
    } catch (e: any) {
        return (
            <ErrorCard />
        )
    }
}

