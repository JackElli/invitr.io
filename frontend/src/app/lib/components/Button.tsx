import { PropsWithChildren } from "react";

interface Props extends PropsWithChildren {
    className?: string;
    loading?: boolean;
    disabled?: boolean;
    onClick?: React.MouseEventHandler<HTMLElement>;
}

const Button = ({ children, className, loading, disabled, onClick }: Props) => {
    return (
        <button onClick={onClick} className={`${className} disabled:opacity-70 flex gap-2 items-center justify-center px-4 py-2 rounded-lg bg-stone-200 border border-gray-300 shadow-sm ${!disabled ? 'hover:shadow-md' : ''}`} disabled={loading || disabled}>
            {children}
            {loading && <p>loading..</p>}
        </button>
    )
}

export default Button;