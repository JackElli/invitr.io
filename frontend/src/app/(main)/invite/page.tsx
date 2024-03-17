'use client'
import { createRef, useEffect } from "react"

export default function Invite() {
    const firstInput = createRef<HTMLInputElement>();

    useEffect(() => {
        firstInput.current?.focus();
    }, [])

    return (
        <>
            <h1 className='text-3xl font-bold'>Create an invite</h1>
            <p className="text-sm">Fill in the details to start.</p>

            <div className="mt-4 border-t border-t-gray-100 pt-2 border-b border-b-gray-100 pb-4">
                <div>
                    <p className="text-sm text-gray-800">Location</p>
                    <input ref={firstInput} className="px-2 py-1 border border-gray-300 rounded-sm outline-none" />
                </div>

                <div className="mt-4">
                    <p className="text-sm text-gray-800">Date</p>
                    <input className="px-2 py-1 border border-gray-300 rounded-sm outline-none" />
                </div>

                <div className="mt-4">
                    <p className="text-sm text-gray-800">Passphrase</p>
                    <input className="px-2 py-1 border border-gray-300 rounded-sm outline-none" />
                </div>


            </div>

            <button className="px-4 py-2 rounded-lg bg-stone-200 mt-4 border border-gray-300 shadow-sm hover:shadow-md">Create invite</button>

        </>
    )
}