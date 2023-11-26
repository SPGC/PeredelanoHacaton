import { TbRobot } from "react-icons/tb";

const RobotPage = () => {
    return (
        <div className="flex flex-row w-full gap-2 text-base bg-gray-700 p-2 rounded-md">
        <div className="flex flex-row gap-3">
            <TbRobot size="56px"/>
            <p className="text-xl"></p>
        </div>
                <div className="text-3xl font-bold bg-gray-900 w-fit p-2 rounded-md">
                    @ibedabot
                </div>
        </div>
    );
};

export default RobotPage;