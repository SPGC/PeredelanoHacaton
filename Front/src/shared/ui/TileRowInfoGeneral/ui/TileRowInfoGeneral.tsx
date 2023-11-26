import {DateAndTime, StatusChecker} from "../../../../widgets/Tiles";


const TileRowInfoGeneral = () => {

    return (
        <>
            <div className="flex flex-row flex-auto gap-6 h-10 p-2 w-fit bg-gray-800 justify-center rounded-md">
                <div className="flex whitespace-nowrap text-xl font-black flex-row gap-1 items-center">
                    ADMIN PANEL
                </div>
            </div>
            <DateAndTime />
            <StatusChecker />
            <div className="flex flex-row gap-6 h-10 p-2 w-fit bg-gray-800 rounded-md">
                <div className="flex flex-row gap-1 items-center text-xs">
                    v.0.0.4
                </div>
            </div>
        </>
    );
};

export default TileRowInfoGeneral;