import ListOfUsers from "../../../widgets/ListOfUsers";
import ListOfIssues from "../../../widgets/ListOfIssues";

export default function ListPage() {


    return (
        <div className="flex flex-col  w-full gap-2 text-base bg-gray-700 p-2 rounded-md">
            <ListOfIssues />
            <ListOfUsers />
        </div>
    );
}

