type Props = {
    isGoing?: boolean;
}

const IsGoing = ({ isGoing }: Props) => {
    if (isGoing == true) {
        return <p className="bg-green-200 px-2 py-1 rounded-xl">Accepted</p>
    }

    if (isGoing == false) {
        return <p className="bg-red-200 px-2 py-1 rounded-xl">Declined</p>
    }

    return <p className="bg-gray-200 px-2 py-1 rounded-xl">not responded</p>
}

export default IsGoing;