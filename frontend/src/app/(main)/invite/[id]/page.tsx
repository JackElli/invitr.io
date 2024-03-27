import ErrorCard from "@/app/lib/components/ErrorCard";
import InviteService from "@/app/lib/services/InviteService";
import PeopleInvited from "./components/PeopleInvited";
import Notes from "./components/Notes";
import InviteOverview from "./components/InviteOverview";
import DeleteButton from "./components/DeleteButton";

export default async function InvitePage({ params }: { params: { id: string } }) {
    try {
        const invite = await InviteService.getById(params.id);
        return (
            <div>
                <div className="border-b border-b-zinc-200 pb-4">
                    <div className="flex justify-between items-center">
                        <InviteOverview invite={invite} />
                        <DeleteButton disabled>Delete event</DeleteButton>
                    </div>
                </div>

                <PeopleInvited invite={invite} />

                <Notes invite={invite} />
            </div>
        )
    } catch (e: any) {
        return (
            <ErrorCard />
        )
    }
}

