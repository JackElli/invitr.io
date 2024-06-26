import Logo from "./Logo";

const Header = () => {
    return (
        <div className='w-full h-10 bg-gray-200 flex items-center sticky top-0 border-b border-b-gray-300'>
            <div className='flex items-center justify-between w-3/4 mx-auto'>
                <a href='/'>
                    <Logo />
                </a>

                <div className="flex gap-3 items-center">
                    <a href='/join'>
                        <p className="text-sm text-gray-700 hover:underline cursor-pointer">Join event</p>
                    </a>

                    <p className="text-sm text-gray-400">About</p>
                </div>
            </div>
        </div>
    )
}

export default Header;