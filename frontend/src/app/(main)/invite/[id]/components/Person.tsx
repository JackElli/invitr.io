import { Person } from "@/app/lib/services/InviteService";
import IsGoing from "./IsGoing";
import { USER } from "@/app/(main)/layout";

type Props = {
    person: Person;
}

const Person = ({ person }: Props) => {
    if (person.name != USER) {
        return (
            <div className="flex items-center justify-between bg-white px-4 py-3 rounded-md border border-gray-200 shadow-sm hover:border-gray-300">
                <div className="flex gap-2 items-center">
                    <div className="w-6 h-6 rounded-full bg-gray-200"></div>
                    <p className="text-lg font-bold">{person.name}</p>
                </div>

                <IsGoing isGoing={person.is_going} />
            </div>
        )
    }
}

export default Person;