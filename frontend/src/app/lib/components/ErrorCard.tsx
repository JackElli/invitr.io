type Props = {
    error?: string;
}

const ErrorCard = ({ error }: Props) => {
    return (
        <div className="bg-red-300 p-4">
            <h1>Something went wrong! {error}</h1>
        </div>
    )
}

export default ErrorCard;