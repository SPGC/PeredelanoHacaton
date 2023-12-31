import TileRowInfoGeneral from "../../../shared/ui/TileRowInfoGeneral";

const MainPage = () => {


    return (
        <>
            <div className="flex flex-row justify-around flex-wrap gap-2 text-base bg-gray-700 p-2 rounded-md">
                <TileRowInfoGeneral />
            </div>
            <div className="flex mt-2 flex-row justify-around flex-wrap gap-2 text-base bg-gray-700 p-2 rounded-md">
            </div>
        </>

    );
};

export default MainPage;