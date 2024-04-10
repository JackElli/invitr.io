'use client'
import ActionButton from "@/app/lib/components/ActionButton";
import Textbox from "@/app/lib/components/Textbox";
import { useRouter } from "next/navigation";
import { useEffect, useRef, useState } from "react";

export default function JoinPage() {
    const [inviteId, setInviteId] = useState("");
    const [key, setKey] = useState("");
    const router = useRouter();
    const inviteIDTextbox = useRef<HTMLInputElement>(null);

    const join = () => {
        if (inviteId == "" || key == "") {
            return
        }

        router.push(`/invite/${inviteId}?key=${key}`);
    }

    useEffect(() => {
        inviteIDTextbox.current?.focus();
    }, [])

    return (
        <>
            <h1 className="text-3xl font-bold">Join an event</h1>
            <h1 className="text-md">All you need is the event ID and your key.</h1>

            <div className="mt-4">
                <Textbox _ref={inviteIDTextbox} name='invite_id' className="w-96" onChange={(e) => setInviteId(e.target.value)}>Invite ID</Textbox>
            </div>

            <div className="mt-4">
                <Textbox name='key' className="w-96" onChange={(e) => setKey(e.target.value)}>Key</Textbox>
            </div>

            <ActionButton className="mt-4" onClick={join}>Join event</ActionButton>
        </>
    )
}