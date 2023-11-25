import Last24hList from "../../../shared/ui/Last24hList";
import Last24HCard from "../../../shared/ui/Last24hCard";

const DashboardPage = () => {
    return (
        <div className="text-base bg-gray-700 p-2 rounded-md">
            <div className="flex flex-row justify-around gap-2 ">
                <Last24HCard numOfComplains={50} />
                <Last24hList
                    numOfCriticalComplains={20}
                    numOfNonUrgentComplains={20}
                    numOfCompletedComplains={10}
                />
                {/* Filler */}
                <div className="flex flex-col
         bg-gray-800 w-60 justify-between
        h-28 rounded-md p-2 text-2xl">Что-то</div>
            </div>

            </div>
    );
};

export default DashboardPage;