'use client'

import Button from "@/app/lib/components/Button";
import Textbox from "@/app/lib/components/Textbox";
import InviteService from "@/app/lib/services/InviteService";
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

    const createEvent = () => {
        if (inputs.title == "" || inputs.location == "" || inputs.date == "" || inputs.passphrase == "" || inputs.people == "") {
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
                            <Textbox name="date" value={inputs.date} onChange={dataOnChange}>Date</Textbox>
                        </div>

                        <div className="mt-4">
                            <Textbox name="passphrase" value={inputs.passphrase} onChange={dataOnChange}>Passphrase</Textbox>
                        </div>

                    </div>
                    <Button className="mt-4" onClick={createEvent}>Create invite</Button>
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