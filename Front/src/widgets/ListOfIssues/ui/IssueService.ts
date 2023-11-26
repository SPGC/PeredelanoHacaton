import axios from "axios";


export const updateIssue = async (formData: any) => {
    try {
        const response = await axios.put(`https://geraback.fly.dev/issues`,
            JSON.stringify(formData));
        console.log('Response:', response);
        console.log(JSON.stringify(formData));
        window.location.reload();
    } catch (error) {
        console.log(JSON.stringify(formData));
        console.error('Error updating issue:', error);
    }
};

export const deleteIssue =  async (formData: any) => {
    try {
        const response = await axios.delete(`https://geraback.fly.dev/issues/${formData.id}`);
        console.log('Response:', response);
        window.location.reload();
    } catch (error) {
        console.error('Error deleting issue:', error);
    }
}
