import {TbMoodSad} from "react-icons/tb";

interface Last24HCardProps {
    topContry: string;
    topOrganization: string;
}

const TopOneList = (props: Last24HCardProps) => {
    const {
        topContry,
        topOrganization
    } = props;

    return (
        <div className="flex flex-auto flex-col gap-1
        bg-gray-800 max-w-1/2 min-w-30 h-28 rounded-md
        p-2 text-md">
            <div className="text-gray-300 flex flex-row justify-between">
                <p className="flex flex-row items-center gap-2"><TbMoodSad size="20px"/> COUNTRY</p>
                <p className="font-bold">{topContry}</p>
            </div>
            <div className="text-gray-300 flex flex-row justify-between">
                <p className="flex flex-row items-center gap-2"><TbMoodSad size="20px" /> ORGANIZATION</p>
                <p className="font-bold">{topOrganization}</p>
            </div>

        </div>
    );
};

export default TopOneList;