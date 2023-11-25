import Box from '@mui/material/Box';
import { DataGrid, GridColDef } from '@mui/x-data-grid';

const columns: GridColDef[] = [
    { field: 'id', headerName: 'ID', width: 90 },
    {
        field: 'status',
        headerName: 'Status',
        width: 150,
    },
    {
        field: 'organization',
        headerName: 'Organization',
        width: 150,
    },
    {
        field: 'description',
        headerName: 'Description',
        width: 200,
    },
    {
        field: 'contact',
        headerName: 'Contact',
        width: 200,
    },
];

const rows = [
    { id: 1, status: 'New', description: 'Issue description here', organization: 'Yandex', contact: '+79162322222' },
    { id: 2, status: 'Completed', description: 'Another issue description', organization: 'Yandex', contact: '+79162322222' },
    { id: 3, status: 'Pending', description: 'Another issue description', organization: 'Yandex', contact: '+79162322222' },
    { id: 4, status: 'Pending', description: 'Another issue description', organization: 'Yandex', contact: '+79162322222' },
    { id: 5, status: 'Pending', description: 'Another issue description', organization: 'Yandex', contact: '+79162322222' },
    { id: 6, status: 'New', description: 'Another issue description', organization: 'Yandex', contact: '+79162322222' },
    { id: 7, status: 'Pending', description: 'Another issue description', organization: 'Yandex', contact: '+79162322222' },
    { id: 8, status: 'Pending', description: 'Another issue description', organization: 'Yandex', contact: '+79162322222' },
    { id: 9, status: 'New', description: 'Another issue description', organization: 'Yandex', contact: '+79162322222' },
    { id: 10, status: 'Pending', description: 'Another issue description', organization: 'Yandex', contact: '+79162322222' },
    { id: 11, status: 'Completed', description: 'Another issue description', organization: 'Yandex', contact: '+79162322222' },
];

export default function ListPage() {
    return (
        <div className="flex flex-col w-full gap-2 text-base bg-gray-700 p-2 rounded-md">
            <div className="p-2 bg-gray-900 w-fit rounded-md text-base">Issues</div>
            <Box sx={{ height: 500, width: '100%' }}>
                <DataGrid
                    style={{ backgroundColor: 'white' }}
                    rows={rows}
                    columns={columns}
                    initialState={{
                        pagination: {
                            paginationModel: {
                                pageSize: 10,
                            },
                        },
                    }}
                    pageSizeOptions={[5]}
                    checkboxSelection
                    disableRowSelectionOnClick
                />
            </Box>
        </div>

    );
}

