import { redirect } from "next/navigation";

export default function Main() {
    const loggedIn = true;
    if (!loggedIn) {
        redirect('/login')
    }
    redirect('/dashboard')
}