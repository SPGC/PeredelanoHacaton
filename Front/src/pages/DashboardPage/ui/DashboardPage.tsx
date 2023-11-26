import {Last24HCard, TopOneList} from "../../../widgets/Tiles";
import Last24hList from "../../../widgets/Tiles/ui/Last24hList.tsx";


const DashboardPage = () => {
    return (
        <div className="text-base bg-gray-700 p-2 rounded-md">
            <div className="flex flex-auto flex-wrap  justify-around gap-2 ">
                <Last24HCard numOfComplains={50} />
                <Last24hList
                    numOfCriticalComplains={20}
                    numOfNonUrgentComplains={20}
                    numOfCompletedComplains={10}
                />
                <TopOneList
                    topContry="Poland"
                    topOrganization="Rieffeisen"
                />

            </div>

            </div>
    );
};

export default DashboardPage;