import UserService from "@/app/lib/services/UserService";
import ErrorCard from "@/app/lib/components/ErrorCard";
import { USER } from "@/app/page";
import ActionButton from "@/app/lib/components/ActionButton";
import { Err } from "@/app/lib/services/Err";

export default async function DashboardLayout({ children }: Readonly<{ children: React.ReactNode; }>) {
    try {
        const user = await UserService.getById(USER);
        return (
            <>
                <div className="bg-gradient-to-r from-[#e1f9fc] to-white border-b border-b-gray-200 py-4">
                    <div className="w-3/4 mx-auto flex justify-between items-center">
                        <div>
                            <h1 className='text-4xl font-bold w-96 lg:w-auto'>Welcome to invitr.io, <span className="font-semibold">{user.firstName} {user.lastName}</span></h1>
                            <p className="text-sm">Start inviting people to your events.</p>
                        </div>

                        <a href='/invite' className="inline-block">
                            <ActionButton>Create invite</ActionButton>
                        </a>
                    </div>
                </div>

                <div className="w-3/4 mx-auto mt-8">
                    <div className="flex gap-4">
                        <h1 className={`text-xl font-bold border-b-2 pb-1 border-b-gray-600 `}>Your events</h1>
                    </div>
                    <div className="mt-4">
                        {children}
                    </div>
                </div>
            </>
        )
    } catch (e) {
        const err = e as Err;
        return (
            <div className="w-3/4 mx-auto mt-2 flex justify-center items-center">
                <ErrorCard error={err.message} />
            </div>
        )
    }
}