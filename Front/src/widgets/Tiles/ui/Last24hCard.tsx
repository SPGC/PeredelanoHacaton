interface Last24HCardProps {
    numOfComplains: number;
}

const Last24HCard = ({numOfComplains}: Last24HCardProps) => {
    return (
        <div className="flex flex-auto flex-row justify-between
        bg-gray-800 max-w-1/2 min-w-30 h-28 rounded-md
        p-2 text-2xl">
            <div className="text-gray-300">LAST 24H</div>
            <div className="flex text-6xl flex-col-reverse items-baseline h-full">
                {numOfComplains}
            </div>
        </div>
    );
};

export default Last24HCard;