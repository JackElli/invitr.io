import Button from "@/app/lib/components/Button";
import ErrorCard from "@/app/lib/components/ErrorCard";
import InviteService from "@/app/lib/services/InviteService";
import { getDate, getDateRelative, isOver } from "@/app/lib/Date";
import Person from "./components/Person";

export default async function InvitePage({ params }: { params: { id: string } }) {
    try {
        const invite = await InviteService.getById(params.id);
        return (
            <div>
                <div className="border-b border-b-zinc-200 pb-4">
                    <div className="flex justify-between items-center">
                        <h1 className="font-bold text-5xl">{invite.title}</h1>
                        <Button disabled className="min-w-32 border-red-300 text-red-500">Delete event</Button>
                    </div>

                    <h1 className="font-semibold text-2xl mt-2">@ {invite.location}</h1>
                    <h1 className="font-semibold text-xl mt-2">{getDate(invite.date)}</h1>
                    <p className="text-sm bg-gray-200 inline-block px-2 rounded-sm">{getDateRelative(invite.date)}</p>
                </div>

                <h1 className="text-xl font-bold mt-10">People invited</h1>
                <div className="flex flex-col gap-4 mt-2">
                    {
                        invite.invitees.map((person, count) => {
                            return (
                                <Person person={person} />
                            )
                        })
                    }
                </div>
            </div>
        )
    } catch (e: any) {
        return (
            <ErrorCard />
        )
    }
}

