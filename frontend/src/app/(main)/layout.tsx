import Header from "../../lib/components/Header";
import "../globals.css";

export default function RootLayout({ children, }: Readonly<{ children: React.ReactNode; }>) {
    return (
        <html lang="en">
            <body className="bg-[#f6f6f6] pb-20">
                <Header />
                <div>
                    {children}
                </div>
            </body>
        </html>
    );
}
