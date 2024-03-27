import Invite from "../components/Invite"

type Props = {
    invites: Invite[]
}

const Events = ({ invites }: Props) => {
    return (
        <div>
            <div className="flex flex-col gap-4 mt-2">
                {
                    invites.length == 0 &&
                    <p>No invites found</p>
                }
                {
                    invites.length > 0 && invites.map((invite, count) => {
                        return (
                            <a href={`/invite/${invite.id}`}>
                                <Invite _key={invite.date + count} invite={invite} />
                            </a>
                        )
                    })
                }
            </div>
        </div>
    )
}

export default Events;