'use client'
import Button from "@/app/lib/components/Button";
import { usePathname } from 'next/navigation'

export default function DashboardLayout({ children }: Readonly<{ children: React.ReactNode; }>) {
    const pathname = usePathname();
    const isEventsSelected = pathname.includes("/dashboard/events");
    return (
        <>
            <div className="bg-gradient-to-r from-[#e1f9fc] to-white border-b border-b-gray-200 py-4">
                <div className="w-3/4 mx-auto flex justify-between items-center">
                    <div>
                        <h1 className='text-4xl font-bold w-96 lg:w-auto'>Welcome to invitr.io, Jack Ellis</h1>
                        <p className="text-sm">Start inviting people to your events.</p>
                    </div>

                    <a href='/invite' className="inline-block">
                        <Button className="bg-green-100 min-w-32">Create invite</Button>
                    </a>
                </div>
            </div>

            <div className="w-3/4 mx-auto mt-8">
                <div className="flex gap-4">
                    {/* can we break these into their own components? */}
                    <a href="/dashboard/events">
                        <h1 className={`text-xl font-bold border-b-2 pb-1 ${isEventsSelected ? ' border-b-gray-600' : 'text-gray-800 border-transparent hover:border-b-gray-600'}`}>Your events</h1>
                    </a>
                    <a href="/dashboard/invites">
                        <h1 className={`text-xl font-bold border-b-2 pb-1 ${!isEventsSelected ? ' border-b-gray-600 ' : 'text-gray-800 border-transparent hover:border-b-gray-600'}`} >Pending invites</h1>
                    </a>

                </div>
                <div className="mt-4">
                    {children}
                </div>
            </div>
        </>
    )
}