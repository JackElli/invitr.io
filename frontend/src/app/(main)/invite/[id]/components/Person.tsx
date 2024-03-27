import { Person } from "@/app/lib/services/InviteService";

type Props = {
    person: Person;
}

const Person = ({ person }: Props) => {
    return (
        <div className="flex items-center justify-between bg-white px-4 py-3 rounded-md border border-gray-200 shadow-sm">
            <div className="flex gap-2 items-center">
                <div className="w-6 h-6 rounded-full bg-gray-200"></div>
                <p className="text-lg font-bold">{person.name}</p>
            </div>

            {person.is_going ?
                person.is_going == true ? <p className="bg-green-200 px-2 py-1 rounded-xl">Accepted</p> :
                    <p className="bg-red-200 px-2 py-1 rounded-xl">Declined</p>
                :
                <p className="bg-gray-200 px-2 py-1 rounded-xl">Not responded</p>
            }

        </div>
    )
}

export default Person;