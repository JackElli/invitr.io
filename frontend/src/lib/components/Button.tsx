import { PropsWithChildren } from "react";

interface Props extends PropsWithChildren {
    className?: string;
    loading?: boolean;
    disabled?: boolean;
    onClick?: React.MouseEventHandler<HTMLElement>;
}

const Button = ({ children, className, loading, disabled, onClick }: Props) => {
    return (
        <button onClick={onClick} disabled={loading || disabled} className={`font-semibold ${className} ${!className?.includes("bg-") ? 'bg-stone-200' : ''}  disabled:opacity-70 flex gap-2 items-center justify-center px-2 py-2 rounded-lg shadow-sm ${!disabled ? 'hover:shadow-md' : ''}`}>
            {children}
            {loading && <p>loading..</p>}
        </button>
    )
}

export default Button;