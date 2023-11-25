import {TbChecks, TbExclamationCircle, TbInfoSquareRounded} from "react-icons/tb";

interface Last24HListProps {
    numOfCriticalComplains: number;
    numOfNonUrgentComplains: number;
    numOfCompletedComplains: number;
}

const Last24HList = (props: Last24HListProps) => {
    const {
        numOfCriticalComplains,
        numOfNonUrgentComplains,
        numOfCompletedComplains
    } = props;

    return (
        <div className="flex flex-col
         bg-gray-800 w-60 justify-between
        h-28 rounded-md p-2 text-2xl">
            <p className="text-gray-300">TOTAL</p>
            <div className="flex w-4/5 flex-row mx-auto items-center justify-between">
                <div className="flex flex-col items-center font-bold">
                    <TbExclamationCircle color="red" size="28px"/>
                    {numOfCriticalComplains}
                </div>
                <div className="flex flex-col items-center font-bold">
                    <TbInfoSquareRounded color="orange" size="28px"/>
                    {numOfNonUrgentComplains}
                </div>
                <div className="flex flex-col items-center font-bold">
                    <TbChecks color="green" size="28px"/>
                    {numOfCompletedComplains}
                </div>
            </div>
        </div>
    );
};

export default Last24HList;