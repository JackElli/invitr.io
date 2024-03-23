import Button from "@/app/lib/components/Button";


export default function Dashboard() {
    return (
        <>
            <div className="border-b border-b-gray-200 pb-4">
                <h1 className='text-3xl font-bold'>Welcome to invitr.io</h1>
                <p className="text-sm">Start inviting people to your events.</p>

                <a href='/invite' className="inline-block">
                    <Button className="mt-4">Start inviting</Button>
                </a>
            </div>

            <div className="mt-12">
                <h1 className='text-xl font-bold'>Your invites</h1>
                <p>No invites found.</p>
            </div>

        </>
    )
} 