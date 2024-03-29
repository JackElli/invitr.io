'use client'
import InviteService, { Invite } from "@/app/lib/services/InviteService";
import { useState } from "react";

type Props = {
    invite: Invite;
    editable?: boolean;
}

const Notes = ({ invite, editable = false }: Props) => {
    const [editing, setEditing] = useState(false);
    const [notes, setNotes] = useState(invite.notes == "" ? "No notes found" : invite.notes);

    const saveNote = () => {
        setEditing(false);
        InviteService.addNotes(invite.id, notes);
    }

    if (!editable) {
        return (
            <>
                <h1 className="text-xl font-bold mt-10">Notes</h1>
                <p className="text-xl">{notes}</p>
            </>
        )
    }

    return (
        <div className="pb-6">
            <h1 className="text-xl font-bold mt-10">Notes</h1>
            <p className="text-sm text-gray-400">(Double click to edit)</p>
            <div className="mt-4">
                {editing ?
                    <textarea onBlur={() => saveNote()} value={notes} onChange={(e) => { setNotes(e.target.value) }} className="w-full rounded-md p-4 outline-none border border-gray-300" autoFocus></textarea> :
                    <p className="text-xl bg-white p-4 rounded-md border border-gray-200" onDoubleClick={() => setEditing(true)}>{notes}</p>
                }
            </div>
        </div>
    )
}

export default Notes;