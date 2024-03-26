import { ChangeEventHandler, PropsWithChildren, Ref } from "react";

interface Props extends PropsWithChildren {
    name: string;
    className?: string;
    value?: string;
    _ref?: Ref<HTMLInputElement>;
    onChange?: ChangeEventHandler<HTMLInputElement>;
}

const Textbox = ({ children, name, value, _ref, onChange }: Props) => {
    return (
        <>
            <p className="text-sm text-gray-800">{children}</p>
            <input name={name} ref={_ref} value={value} onChange={onChange} className="px-2 py-1 border border-gray-300 rounded-sm outline-none" />
        </>
    )
}

export default Textbox;