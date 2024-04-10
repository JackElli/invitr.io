import { redirect } from "next/navigation";

export let USER = "123";

export default function Main() {
    const loggedIn = USER != "";

    if (!loggedIn) {
        redirect('/home')
    }

    redirect('/dashboard')
}