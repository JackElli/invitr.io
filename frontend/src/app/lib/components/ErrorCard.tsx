type Props = {
    error?: string;
}

const ErrorCard = ({ error }: Props) => {
    return (
        <div className="bg-red-200 p-4 shadow-md rounded-md border border-gray-400 ">
            <h1>Something went wrong! {error}</h1>
        </div>
    )
}

export default ErrorCard;