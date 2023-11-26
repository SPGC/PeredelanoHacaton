import { useEffect, useState } from 'react';
import { TbCalendar, TbClock } from 'react-icons/tb';

const DateTimeDisplay = () => {
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

    return (
        <div className="flex flex-auto flex-row gap-6 h-10 whitespace-nowrap p-2 w-64 justify-center bg-gray-800 rounded-md">
            <div className="flex flex-row gap-1 items-center"><TbCalendar /> {formattedDate}</div>
            <div className="flex flex-row gap-1 items-center "><TbClock /> {formattedTime}</div>
        </div>
    );
};

export default DateTimeDisplay;