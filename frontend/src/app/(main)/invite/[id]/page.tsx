import ErrorCard from "@/app/lib/components/ErrorCard";
import InviteService from "@/app/lib/services/InviteService";
import PeopleInvited from "./components/PeopleInvited";
import Notes from "./components/Notes";
import InviteOverview from "./components/InviteOverview";
import DeleteButton from "./components/DeleteButton";
import { ChangeGoing } from "./components/ChangeGoing";


export default async function InvitePage({ params, searchParams }: { params: { id: string }, searchParams: { [key: string]: string | string[] | undefined } }) {
    try {
        const invite = await InviteService.getById(params.id);

        const key = searchParams['key'] as string;
        const USER = await InviteService.getUserFromKey(invite.id, key);

        if (!USER) {
            return <p>You are not allowed to be here :)</p>
        }

        const isOrganiser = invite.organiser == USER;

        // need to check if they've got the correct key
        return (
            <div>
                <div>
                    <div className="flex justify-between items-center">
                        <InviteOverview invite={invite} />
                        {!isOrganiser && <ChangeGoing invite={invite} searchParams={searchParams} />}
                        {isOrganiser && <DeleteButton disabled>Delete event</DeleteButton>}
                    </div>
                </div>

                <PeopleInvited invite={invite} />

                <Notes invite={invite} editable={isOrganiser} />
            </div>
        )

    } catch (e: any) {
        return (
            <ErrorCard />
        )
    }
}

