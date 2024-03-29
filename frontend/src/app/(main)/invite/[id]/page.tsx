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

        // event is NOT owned by the user that's logged in
        // need to check if they've got the correct key
        if (invite.organiser != USER) {
            return (
                <div>
                    <div>
                        <div className="flex justify-between items-center">
                            <InviteOverview invite={invite} />
                            <ChangeGoing invite={invite} searchParams={searchParams} />
                        </div>
                    </div>

                    <PeopleInvited invite={invite} />

                    <Notes invite={invite} />
                </div>
            )
        }

        // event IS owned by the user that's logged in
        return (
            <div>
                <div className="flex justify-center items-center rounded-b-lg w-full h-10 bg-gradient-to-r from-[#e1f9fc] to-white border border-gray-200 -mt-10 mb-4">
                    <h1 className="text-xl text-center font-bold">This is your event!</h1>
                </div>
                <div>
                    <div className="flex justify-between items-center">
                        <InviteOverview invite={invite} />
                        <DeleteButton disabled>Delete event</DeleteButton>
                    </div>
                </div>

                <PeopleInvited invite={invite} />

                <Notes invite={invite} editable />
            </div>
        )
    } catch (e: any) {
        return (
            <ErrorCard />
        )
    }
}

