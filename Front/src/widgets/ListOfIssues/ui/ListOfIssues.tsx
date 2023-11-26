import axios from "axios";
import {DataGrid, GridColDef} from "@mui/x-data-grid";
import {SetStateAction, useEffect, useState} from "react";
import {UniversalModal} from "../../../shared/ui/ModalWrapper/ModalWrapper.tsx";
import IssueForm from "./IssueFormEdit.tsx";
import {updateIssue} from "./IssueService.ts";

const columns: GridColDef[] = [
    {field: 'id', headerName: 'ID', width: 70},
    {field: 'status', headerName: 'Status', width: 130},
    {field: 'description', headerName: 'Description', width: 200},
    {field: 'organisation_id', headerName: 'Organization ID', width: 130},
    {field: 'validated', headerName: 'Validated', width: 80},
    {field: 'organisation_name', headerName: 'Org Name', width: 140},
    {field: 'organisation_country', headerName: 'Org Country', width: 140},
];

type issueRow = {
    id: number;
    status: string;
    description: string;
    organisation_id: string;
    validation: boolean;
    organisation_name: string;
    organisation_country: string;
}

const fetchData = async () => {
    const maxAttempts = 8;
    let attempt = 0;
    let response;

    while (attempt < maxAttempts) {
        try {
            response = await axios.get('https://geraback.fly.dev/issues?page=1&limit=1000');
            if (response.status === 200) {
                const responseData = typeof response.data === 'string' ? JSON.parse(response.data) : response.data;
                if (responseData && responseData.data) {
                    const formattedData = responseData.data.map((item: any) => ({
                        id: item.id || 'No ID',
                        status: item.status,
                        description: item.description,
                        organisation_id: item.organisation_id || 0,
                        validated: item.validation || false,
                        organisation_name: item.organisation_name || 'No Name',
                        organisation_country: item.organisation_country || 'NULL',
                    }));
                    return formattedData;
                }
                return [];
            }
        } catch (error) {
            console.error(`Attempt ${attempt + 1} failed:`, error);
            attempt++;
            if (attempt >= maxAttempts) {
                console.error("Max attempts reached, failing...");
                return [];
            }
        }
    }
};


const ListOfIssues = () => {
    const [rows, setRows] = useState([]);
    const [open, setOpen] = useState(false);
    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);

    const [selectedRow, setSelectedRow] = useState<issueRow | null>(null);

    const handleRowClick = (params: { row: SetStateAction<issueRow | null>; }) => {
        setSelectedRow(params.row);
        handleOpen();
    };


    useEffect(() => {
        fetchData().then(data => {
            setRows(data);
        });
    }, []);



    return (
        <div>
            <div className="p-2 w-fit rounded-md text-base">Issues</div>

            <UniversalModal
                open={open}
                onClose={handleClose}
                title="ISSUES"
            children={
                <div>
                    <div className="p-2 w-fit rounded-md text-xl">ISSUE</div>
                    <div className="bg-gray-700 p-2 rounded-md">
                        <p>ID: {selectedRow?.id}</p>
                        <p>Status: {selectedRow?.status}</p>
                        <p>Desc: {selectedRow?.description}</p>
                        <p>Organization ID: {selectedRow?.organisation_id}</p>
                        <p>Validated: {selectedRow?.validation}</p>
                        <p>Validated: {selectedRow?.organisation_country}</p>
                        <p>Validated: {selectedRow?.organisation_name}</p>
                    </div>
                    <IssueForm initialData={selectedRow} onSubmit={updateIssue} />
                </div>} />

            <DataGrid

                style={{ backgroundColor: 'white'}}
                rows={rows}
                columns={columns}
                onRowClick={handleRowClick}
            />
        </div>
    );
};

export default ListOfIssues;