import Invite from "../components/Invite"

type Props = {
    invites: Invite[]
}

const Events = ({ invites }: Props) => {
    return (
        <div>
            <div className="flex flex-col gap-4 mt-2">

                {
                    invites.length > 0 && invites.map((invite, count) => {
                        return (
                            <Invite _key={invite.date + count} invite={invite} />
                        )
                    })
                }
            </div>
        </div>
    )
}

export default Events;