import {HiOutlineExclamation} from "react-icons/hi";
import {Link} from "react-router-dom";

const PageError = () => {
    return (
        <div className="flex items-center justify-center h-screen w-screen text-3xl">
            <div className="flex flex-col gap-5 items-center justify-center">
                <HiOutlineExclamation className="text-red-500" size="72px"/>
                <p className="text-red-500">
                    Something went wrong...
                </p>
                <Link to={'/'}>
                    <p className="
                    text-[16px] border-2 px-2 rounded-md border-red-200
                    ">{'-> HOME <-'}</p>
                </Link>
            </div>

        </div>
    );
};

export default PageError;