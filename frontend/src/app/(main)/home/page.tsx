import ActionButton from "@/app/lib/components/ActionButton";

export default async function HomePage() {
    return (
        <>
            <h1 className="text-xl font-semibold">What is invitr.io?</h1>
            <p className="text-lg">invitr.io is an application allowing you to create and send invites easily in either email or physical form.</p>

            <h1 className="text-xl font-semibold mt-8">Why invitr.io?</h1>
            <p className="text-lg">Easy to use UI, easy calendar integration, no need to log in to respond to invites.</p>

            <ul className="list-disc mt-8 text-xl list-inside">
                <li>Free if you're responding to invites</li>
                <li>Perfect integration for organisations</li>
                <li>Super easy to invite people to your events</li>
                <li>Add notes to your events</li>
            </ul>

            <h1 className="text-xl font-semibold mt-8">Want invitr.io for your organisation?</h1>
            <p className="text-lg">Talk to us at <span className="font-medium">test@test.com!</span></p>

            <a href='/login' className="inline-block mt-4">
                <ActionButton className="bg-blue-500 text-white">Talk to us</ActionButton>
            </a>
        </>

    )
}