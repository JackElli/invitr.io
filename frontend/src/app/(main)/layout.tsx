import type { Metadata } from "next";
import Header from "../lib/components/Header";
import "../globals.css";

export const metadata: Metadata = {
    title: "invitr.io",
};
export const USER = "123";

export default function RootLayout({ children, }: Readonly<{ children: React.ReactNode; }>) {
    return (
        <html lang="en">
            <body className="bg-[#f6f6f6]">
                <Header />
                <div className="w-3/4 mx-auto mt-10">
                    {children}
                </div>
            </body>
        </html>
    );
}
