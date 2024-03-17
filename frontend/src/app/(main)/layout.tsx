import type { Metadata } from "next";
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });
import "../globals.css";

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
            <body className="bg-[#f7f5f0]">
                <div className='w-full h-10 bg-stone-200 flex items-center border-b border-b-gray-300'>
                    <div className='w-3/4 mx-auto'>
                        <a href='/'>
                            <h1 className="font-bold px-2 py-1 bg-green-200 inline rounded-lg">invitr.io</h1>
                        </a>

                    </div>
                </div>

                <div className="w-3/4 mx-auto mt-10">
                    {children}
                </div>
            </body>
        </html>
    );
}
