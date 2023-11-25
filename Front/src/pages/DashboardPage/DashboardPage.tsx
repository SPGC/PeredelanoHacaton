import Last24HCard from "../../shared/ui/Last24hCard/ui/Last24hCard.tsx";
import Last24hList from "../../shared/ui/Last24hList";

const DashboardPage = () => {
    return (
        <div className="text-base bg-gray-700 p-2 rounded-md">
            <div className="flex flex-row gap-2 ">
                <Last24HCard numOfComplains={50} />
                <Last24hList
                    numOfCriticalComplains={20}
                    numOfNonUrgentComplains={20}
                    numOfCompletedComplains={10}
                />
            </div>

            </div>
    );
};

export default DashboardPage;