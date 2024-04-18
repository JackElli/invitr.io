import Button from "@/lib/components/Button";
import { PropsWithChildren } from "react";

interface Props extends PropsWithChildren {
    disabled?: boolean;
}

const DeleteButton = ({ children, disabled }: Props) => {
    return (
        <Button disabled={disabled} className="min-w-32 border-red-300 text-red-500">
            {children}
        </Button>
    )
}

export default DeleteButton;