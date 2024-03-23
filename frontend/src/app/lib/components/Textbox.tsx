import { PropsWithChildren, Ref } from "react";

interface Props extends PropsWithChildren {
    className?: string;
    _ref?: Ref<HTMLInputElement>;
}

const Textbox = ({ children, _ref }: Props) => {
    return (
        <>
            <p className="text-sm text-gray-800">{children}</p>
            <input ref={_ref} className="px-2 py-1 border border-gray-300 rounded-sm outline-none" />
        </>
    )
}

export default Textbox;