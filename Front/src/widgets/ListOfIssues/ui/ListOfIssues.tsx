import axios from "axios";
import {DataGrid, GridColDef} from "@mui/x-data-grid";
import {useEffect, useState} from "react";

const columns: GridColDef[] = [
    { field: 'id', headerName: 'ID', width: 70 },
    { field: 'status', headerName: 'Status', width: 130 },
    { field: 'description', headerName: 'Description', width: 200 },
    { field: 'organisation_id', headerName: 'Organization ID', width: 130 },
    { field: 'validated', headerName: 'Validated', width: 80 },
];

const fetchData = async () => {
    try {
        const response = await axios.get('/api/issues?page=1&limit=1000');
        console.log(response.data);

        const responseData = typeof response.data === 'string' ? JSON.parse(response.data) : response.data;

        if (responseData && responseData.data) {
            const formattedData = responseData.data.map((item: { id: number; status: string; description: string; organisation_id: string; validation: boolean; }) => ({
                id: item.id || 'No ID',
                status: item.status,
                description: item.description,
                organisation_id: item.organisation_id || 0,
                validated: item.validation || false
            }));
            return formattedData;
        }
        return [];
    } catch (error) {
        console.error(error);
        return [];
    }
}

const ListOfIssues = () => {
    const [rows, setRows] = useState([]);

    useEffect(() => {
        fetchData().then(data => {
            setRows(data);
        });
    }, []);

    return (
        <div>
            <div className="p-2 w-fit rounded-md text-base">Issues</div>
            <DataGrid
                style={{ backgroundColor: 'white'}}
                rows={rows}
                columns={columns}
                pageSizeOptions={[10]}
            />
        </div>
    );
};

export default ListOfIssues;