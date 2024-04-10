import { PropsWithChildren } from "react"
import Button from "./Button"

interface Props extends PropsWithChildren {
    className?: string;
    disabled?: boolean;
    onClick?: React.MouseEventHandler<HTMLElement>;
}

const ActionButton = ({ className, disabled, onClick, children }: Props) => {
    return <Button onClick={onClick} className={`${className} bg-blue-500 ${!disabled ? 'hover:bg-blue-600 ' : ''} disabled:cursor-not-allowed text-white`} disabled={disabled}>{children}</Button>
}

export default ActionButton;