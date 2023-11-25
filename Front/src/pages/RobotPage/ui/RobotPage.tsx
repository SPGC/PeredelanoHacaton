import { TbRobot } from "react-icons/tb";

const RobotPage = () => {
    return (
        <div className="flex flex-col w-full gap-2 text-base bg-gray-700 p-2 rounded-md">
        <div className="flex flex-row gap-3">
            <TbRobot size="56px"/>
        </div>
        </div>
    );
};

export default RobotPage;