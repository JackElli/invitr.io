import { redirect } from "next/navigation";

export let USER = "";

export default function Main() {
    const loggedIn = USER != "";
    if (!loggedIn) {
        redirect('/home')
    }
    redirect('/dashboard')
}