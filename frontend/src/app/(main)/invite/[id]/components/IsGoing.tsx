type Props = {
    isGoing?: boolean;
    shadow?: boolean;
}

const IsGoing = ({ isGoing, shadow }: Props) => {
    if (isGoing == true) {
        return <p className={`bg-green-200 px-2 py-1 rounded-xl ${shadow ? 'shadow-lg' : ''}`}>Accepted</p>
    }

    if (isGoing == false) {
        return <p className={`bg-red-200 px-2 py-1 rounded-xl ${shadow ? 'shadow-lg' : ''}`}>Declined</p>
    }

    return <p className="bg-gray-200 px-2 py-1 rounded-xl">not responded</p>
}

export default IsGoing;