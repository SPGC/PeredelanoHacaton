import {TbCalendar, TbClock, TbDatabaseImport, TbRobot, TbServer } from "react-icons/tb";
import {useEffect, useState} from "react";

const TileRowInfoGeneral = () => {

    const [currentDate, setCurrentDate] = useState(new Date());

    useEffect(() => {
        const timer = setInterval(() => {
            setCurrentDate(new Date());
        }, 1000);

        return () => clearInterval(timer);
    }, []);

    const formattedDate = currentDate.toLocaleDateString('en-US', {
        month: 'short',
        day: 'numeric',
        year: 'numeric'
    });
    const formattedTime = currentDate.toLocaleTimeString('ru-RU');

    // TODO: 'Add functionality to change status + optimize responsive design for mobile'
    return (
        <>
            <div className="flex flex-row flex-auto gap-6 h-10 p-2 w-fit bg-gray-800 justify-center rounded-md">
                <div className="flex whitespace-nowrap text-xl font-black flex-row gap-1 items-center">
                    ADMIN PANEL
                </div>
            </div>
            <div className="flex flex-auto flex-row gap-6 h-10 whitespace-nowrap p-2 w-64 justify-center bg-gray-800 rounded-md">
                <div className="flex flex-row gap-1 items-center"><TbCalendar /> {formattedDate}</div>
                <div className="flex flex-row gap-1 items-center "><TbClock /> {formattedTime}</div>
            </div>
            <div className="flex flex-auto flex-row justify-center gap-6 h-10 p-2 w-fit bg-gray-800 rounded-md">
                <div className="flex flex-row gap-1 items-center ">
                    Status:
                    <TbServer size="24px" color="green"/>
                    |
                    <TbDatabaseImport size="24px" color="green"/>
                    |
                    <TbRobot size="24px" color="green"/>
                </div>
            </div>

            <div className="flex flex-row gap-6 h-10 p-2 w-fit bg-gray-800 rounded-md">
                <div className="flex flex-row gap-1 items-center text-xs">
                    v.0.0.4
                </div>
            </div>
        </>

    );
};

export default TileRowInfoGeneral;