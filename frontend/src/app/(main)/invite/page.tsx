'use client'

import Button from "@/app/lib/components/Button";
import Textbox from "@/app/lib/components/Textbox";
import { createRef, useEffect } from "react"

export default function Invite() {
    const firstInput = createRef<HTMLInputElement>();

    useEffect(() => {
        firstInput.current?.focus();
    }, [])

    const createEvent = () => {
        alert("CLICKED");
    }

    return (
        <>
            <h1 className='text-3xl font-bold'>Create an invite</h1>
            <p className="text-sm">Fill in the details to start.</p>

            <div className="mt-4 border-t border-t-gray-100 pt-2 border-b border-b-gray-100 pb-4">
                <div>
                    <Textbox _ref={firstInput}>Location</Textbox>
                </div>

                <div className="mt-4">
                    <Textbox>Date</Textbox>
                </div>

                <div className="mt-4">
                    <Textbox>Passphrase</Textbox>
                </div>

            </div>
            <Button className="mt-4" onClick={createEvent}>Create invite</Button>
        </>
    )
}