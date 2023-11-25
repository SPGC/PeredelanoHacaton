import {TbCalendar, TbClock} from "react-icons/tb";
import {useEffect, useState} from "react";

const TileRowInfoGeneral = () => {

    const [currentDate, setCurrentDate] = useState(new Date());

    useEffect(() => {
        const timer = setInterval(() => {
            setCurrentDate(new Date());
        }, 1000);

        // Clear the interval when the component unmounts
        return () => clearInterval(timer);
    }, []);

    const formattedDate = currentDate.toLocaleDateString('en-US', {
        month: 'short', // Abbreviated month name
        day: 'numeric', // Numeric day
        year: 'numeric' // Four digit year
    });
    const formattedTime = currentDate.toLocaleTimeString('ru-RU');


    return (
    <div className="flex flex-row gap-6 h-10 p-2 w-full bg-gray-800 rounded-sm">
        <div className="flex flex-row gap-1 items-center"><TbCalendar /> {formattedDate}</div>
        <div className="flex flex-row gap-1 items-center"><TbClock /> {formattedTime}</div>
    </div>
    );
};

export default TileRowInfoGeneral;