import  { useState } from 'react';
import {TbTrash} from "react-icons/tb";
import {deleteIssue} from "./IssueService.ts";

// @ts-ignore
const IssueForm = ({ initialData, onSubmit }) => {
    const [formData, setFormData] = useState({
        id: initialData.id || 0,
        status: initialData.status || '',
        description: initialData.description || '',
        organisation_id: initialData.organisation_id || '',
        validation: initialData.validation || false,
        organisation_name: initialData.organisation_name || '',
        organisation_country: initialData.organisation_country || '',
    });

    const handleChange = (e: any) => {
        const { name, value, type, checked } = e.target;
        setFormData({
            ...formData,
            [name]: type === 'checkbox' ? checked : value,
        });
    };

    const handleDelete = (e: any) => {
        e.preventDefault();
        deleteIssue(formData);
    }

    const handleSubmit = (e: any) => {
        e.preventDefault();
        onSubmit(formData);
        // Здесь вы можете вызвать функцию, которая отправляет formData на сервер
    };

    return (
        <form onSubmit={handleSubmit} className="flex flex-col w-96">
                <label className="p-2 mt-2 bg-gray-600 rounded-md flex justify-between flex-row ">
                    ID:
                    <input
                        type="number"
                        name="id"
                        value={formData.id}
                        onChange={handleChange}
                        disabled
                    />
                </label>
            <label className="p-2 mt-2 bg-gray-600 rounded-md flex justify-between flex-row ">
                Status:
                <input
                    className="text-black"
                    type="text"
                    name="status"
                    value={formData.status}
                    onChange={handleChange}
                />
            </label>
            <label className="p-2 mt-2 bg-gray-600 rounded-md flex justify-between flex-row ">
                Description:
                <textarea
                    className="text-black"
                    name="description"
                    value={formData.description}
                    onChange={handleChange}
                />
            </label>
            <label className="p-2 mt-2 bg-gray-600 rounded-md flex justify-between flex-row ">
                Organization Name:
                <textarea
                    className="text-black"
                    name="organisation_name"
                    value={formData.organisation_name}
                    onChange={handleChange}
                />
            </label>
            <label className="p-2 mt-2 bg-gray-600 rounded-md flex justify-between flex-row ">
                Organization Country:
                <textarea
                    className="text-black"
                    name="organisation_country"
                    value={formData.organisation_country}
                    onChange={handleChange}
                />
            </label>
            <label className="p-2 mt-2 bg-gray-600 rounded-md flex justify-between flex-row ">
                Org ID:
                <input
                    disabled
                    className="text-black"
                    type="text"
                    name="organisation_id"
                    value={formData.organisation_id}
                    onChange={handleChange}
                />
            </label>
            <label className="p-2 mt-2 bg-gray-600 rounded-md flex justify-between flex-row ">
                Validated:
                <input
                    type="checkbox"
                    name="validated"
                    checked={formData.validation}
                    onChange={handleChange}
                />
            </label>
            <div className="flex flex-row justify-between items-center">
                <button onClick={handleDelete} className="bg-red-500 p-2 mt-3 rounded-md"><TbTrash size="24px"/></button>
                <button className="bg-orange-500 w-fit p-2 mt-3 ml-auto rounded-md" type="submit">Update</button>
            </div>
        </form>
    );
};

export default IssueForm;