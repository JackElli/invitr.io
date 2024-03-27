import type { Metadata } from "next";

export const metadata: Metadata = {
    title: "invitr.io",
};

export default function RootLayout({ children, }: Readonly<{ children: React.ReactNode; }>) {
    return (
        <html lang="en" className="pb-20">
            {children}
        </html>
    );
}
