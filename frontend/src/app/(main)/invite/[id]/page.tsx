import Button from "@/app/lib/components/Button";
import ErrorCard from "@/app/lib/components/ErrorCard";
import InviteService from "@/app/lib/services/InviteService";
import { getDate, getDateRelative, isOver } from "@/app/lib/components/Date";

export default async function Invite({ params }: { params: { id: string } }) {
    try {
        const invite = await InviteService.getById(params.id);
        const hasPassed = isOver(invite.date);
        return (
            <>
                <div className="border-b border-b-zinc-200 pb-4">
                    <div className="flex justify-between items-center">
                        <h1 className="font-bold text-5xl flex-grow-0">{invite.title}</h1>
                        <Button disabled className="flex-grow-0 min-w-32 border-red-300 text-red-500">Delete event</Button>
                    </div>

                    <h1 className="font-semibold text-2xl mt-2">{invite.location}</h1>
                    <div className="flex gap-2 items-center mt-2">
                        <h1 className="font-semibold text-xl">{getDate(invite.date)}</h1>
                        <p className={`px-4 py-1  rounded-2xl border shadow-sm bg-gray-100 ${!hasPassed ? ' animate-pulse  border-green-400' : 'border-red-200'}`}>{getDateRelative(invite.date)}</p>
                    </div>

                </div>

                <h1 className="text-xl font-bold mt-10">People invited</h1>
                <div className="flex flex-col gap-2 mt-2">
                    {
                        invite.invitees.map((person, count) => {
                            return (
                                <p key={person + count}>{person}</p>
                            )
                        })
                    }
                </div>
            </>
        )
    } catch (e: any) {
        return (
            <ErrorCard />
        )
    }
}

