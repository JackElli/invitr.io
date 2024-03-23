import type { Metadata } from "next";
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });
import "../globals.css";
import Header from "../lib/components/Header";

export const metadata: Metadata = {
    title: "invitr.io",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
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
