interface Last24HCardProps {
    numOfComplains: number;
}

const Last24HCard = ({numOfComplains}: Last24HCardProps) => {
    return (
        <div className="flex flex-row justify-between bg-gray-800 w-60 h-28 rounded-md p-2 text-2xl">
            <div className="text-gray-300">LAST 24H</div>
            <div className="flex font-black text-6xl flex-col-reverse items-baseline h-full">
                {numOfComplains}
            </div>
        </div>
    );
};

export default Last24HCard;