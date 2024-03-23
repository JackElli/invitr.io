import { PropsWithChildren } from "react";

interface Props extends PropsWithChildren {
    className?: string;
    loading?: boolean;
    onClick?: React.MouseEventHandler<HTMLElement>;
}

const Button = ({ children, className, loading, onClick }: Props) => {
    return (
        <button onClick={onClick} className={`${className} disabled:opacity-70 flex gap-2 items-center justify-center px-4 py-2 rounded-lg bg-stone-200 border border-gray-300 shadow-sm ${!loading ? 'hover:shadow-md' : ''}`} disabled={loading}>
            {children}
            {loading && <p>loading..</p>}
        </button>
    )
}

export default Button;