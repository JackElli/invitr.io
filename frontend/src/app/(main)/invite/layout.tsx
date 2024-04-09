export default function InviteLayout({ children }: Readonly<{ children: React.ReactNode; }>) {
    return (
        <div className="w-3/4 mx-auto mt-6">
            {children}
        </div>
    );
}