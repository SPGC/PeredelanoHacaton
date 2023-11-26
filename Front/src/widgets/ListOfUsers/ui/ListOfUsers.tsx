import {useEffect, useState} from 'react';
import {DataGrid, GridColDef} from "@mui/x-data-grid";
import axios from "axios";

const columns: GridColDef[] = [
    { field: 'id', headerName: 'ID', width: 70 },
    { field: 'name', headerName: 'Name', width: 130 },
    { field: 'contact_info', headerName: 'Contact Info', width: 200 },
    { field: 'amount_of_issues', headerName: 'Issues Count', width: 130 },
];

const fetchData = async () => {
    try {
        const response = await axios.get('http://172.16.15.7:8080/users?page=1&limit=1000', {
        });
        console.log(response.data);

        const responseData = typeof response.data === 'string' ? JSON.parse(response.data) : response.data;

        if (responseData && responseData.data) {
            const formattedData = responseData.data.map((item: { id: number; name: string; contact_info: string; amount_of_issues: number; }) => ({
                id: item.id,
                name: item.name || 'No Name',
                contact_info: item.contact_info || 'No Contact Info',
                amount_of_issues: item.amount_of_issues || 0
            }));
            return formattedData;
        }
        return [];
    } catch (error) {
        console.error(error);
        return [];
    }
}

const ListOfUsers = () => {
    const [rows, setRows] = useState([]);

    fetchData();

    useEffect(() => {
        fetchData().then(data => {
            setRows(data);
        });
    }, []);

    return (
        <div>
            <div className="p-2 w-fit rounded-md text-base">Users</div>
            <DataGrid
                style={{ backgroundColor: 'white'}}
                rows={rows}
                columns={columns}
                pageSizeOptions={[10]}
            />
        </div>
    );
};

export default ListOfUsers;