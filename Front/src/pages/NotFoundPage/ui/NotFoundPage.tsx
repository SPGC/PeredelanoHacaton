import {TbError404} from "react-icons/tb";

const NotFoundPage = () => {
    return (
        <div className="flex flex-col justify-center
        items-center flex-wrap gap-2 text-base bg-gray-700
        p-2 rounded-md h-80">
            <TbError404 size="72px"/>
            <p className="text-2xl font-black">Page was not found</p>
        </div>
    );
};

export default NotFoundPage;