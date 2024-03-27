'use client'
import InviteService, { Invite } from "@/app/lib/services/InviteService";
import { useEffect, useRef, useState } from "react";

type Props = {
    invite: Invite;
}

const Notes = ({ invite }: Props) => {
    const [editing, setEditing] = useState(false);
    const [notes, setNotes] = useState(invite.notes == "" ? "Double click here to add event notes." : invite.notes);

    const saveNote = () => {
        setEditing(false);
        InviteService.addNotes(invite.id, notes);
    }

    return (
        <>
            <h1 className="text-xl font-bold mt-10">Notes</h1>
            <p className="text-sm text-gray-400">(Double click to edit)</p>
            <div className="mt-4">
                {editing ?
                    <textarea onBlur={() => saveNote()} value={notes} onChange={(e) => { setNotes(e.target.value) }} className="mt-2 w-full rounded-sm p-4 outline-none border border-gray-300" autoFocus></textarea> :
                    <p className="text-xl" onDoubleClick={() => setEditing(true)}>{notes}</p>
                }
            </div>

        </>
    )
}

export default Notes;