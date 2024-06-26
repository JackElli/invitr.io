'use client'

import ActionButton from "@/lib/components/ActionButton";
import Button from "@/lib/components/Button";
import Textbox from "@/lib/components/Textbox";
import UnderConstruction from "@/lib/components/UnderConstruction";
import InviteService from "@/lib/services/InviteService";
import { useRouter } from "next/navigation";

import { ChangeEvent, useEffect, useRef, useState } from "react"

type Inputs = {
    title: string;
    location: string;
    date: string;
    passphrase: string;
    people: string;
}

export default function NewInvitePage() {

    const titleInput = useRef<HTMLInputElement>(null);
    const router = useRouter();

    const [inputs, setInputs] = useState<Inputs>({
        title: "",
        location: "",
        date: "",
        passphrase: "",
        people: ""
    })

    useEffect(() => {
        titleInput.current?.focus();
    }, [])

    const dataOnChange = (e: ChangeEvent<HTMLInputElement>) => {
        const name = e.target.name;
        const value = e.target.value;

        setInputs(values => ({ ...values, [name]: value }));
    }

    const notValid = inputs.title == "" || inputs.location == "" || inputs.date == "" || inputs.passphrase == "" || inputs.people == ""

    const createEvent = () => {
        if (notValid) {
            return;
        }

        try {
            InviteService.new(
                inputs.title,
                inputs.location,
                inputs.date,
                inputs.passphrase,
                inputs.people
            );

            router.push('/dashboard');
        } catch (e) {
            console.log(e)
        }
    }

    return (
        <>
            <UnderConstruction />
            <h1 className='text-3xl font-bold'>Create an invite</h1>
            <p className="text-sm">Fill in the details to start.</p>
            <div className="flex gap-32 items-center">
                <div>
                    <div className="mt-4 border-t border-t-gray-100 pt-2 border-b border-b-gray-100 pb-4">
                        <div>
                            <Textbox name="title" _ref={titleInput} value={inputs.title} onChange={dataOnChange}>Title</Textbox>
                        </div>
                        <div className="mt-4">
                            <Textbox name="location" value={inputs.location} onChange={dataOnChange}>Location</Textbox>
                        </div>

                        <div className="mt-4">
                            <Textbox name="date" value={inputs.date} onChange={dataOnChange}>Date (YYYY-mm-dd)</Textbox>
                        </div>

                        <div className="mt-4">
                            <Textbox name="passphrase" value={inputs.passphrase} onChange={dataOnChange}>Passphrase</Textbox>
                        </div>

                    </div>
                    <ActionButton className="mt-2" onClick={createEvent} disabled={notValid}>Create invite</ActionButton>
                </div>

                <div>
                    <h1 className="font-bold text-xl">Invite your people</h1>
                    <div className="mt-2">
                        <Textbox name="people" className="w-96" value={inputs.people} onChange={dataOnChange}>Add people (comma separated for now)</Textbox>
                    </div>
                </div>
            </div>
        </>
    )
}