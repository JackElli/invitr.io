import Button from "@/app/lib/components/Button";

export default function HomeLayout({ children }: Readonly<{ children: React.ReactNode; }>) {
    return (
        <>
            <div className="bg-gradient-to-r from-[#e1f9fc] to-white border-b border-b-gray-200 py-4">
                <div className="w-3/4 mx-auto flex justify-between items-center">
                    <div>
                        <h1 className='text-4xl font-bold w-96 lg:w-auto'>Welcome to invitr.io</h1>
                        <p className="text-sm">Start inviting people to your events.</p>
                    </div>
                    <div className="flex gap-2 items-center">
                        <a href='/join' className="inline-block">
                            <Button className="bg-green-600 text-white">Join event</Button>
                        </a>

                        <a href='/login' className="inline-block">
                            <Button className="bg-blue-500 text-white">Log in</Button>
                        </a>
                    </div>

                </div>
            </div>
            <div className="w-3/4 mx-auto mt-6">
                {children}
            </div>
        </>
    );
}