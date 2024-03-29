export default function JoinLayout({ children }: Readonly<{ children: React.ReactNode; }>) {
    return (
        <div className="w-3/4 mx-auto mt-10">
            {children}
        </div>
    );
}