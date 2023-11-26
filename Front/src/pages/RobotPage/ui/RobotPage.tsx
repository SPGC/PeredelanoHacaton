import { TbRobot } from "react-icons/tb";

const RobotPage = () => {
    return (
        <div className="flex flex-col w-full gap-2 text-base bg-gray-700 p-2 rounded-md">
        <div className="flex flex-row gap-3">
            <TbRobot size="56px"/>
            <p className="text-xl"></p>
        </div>
            <div className="flex flex-col gap-2 bg-gray-600 rounded-md h-20 p-2">
               </div>
        </div>
    );
};

export default RobotPage;