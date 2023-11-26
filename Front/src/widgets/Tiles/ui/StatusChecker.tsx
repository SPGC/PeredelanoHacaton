import {useEffect, useState} from 'react';
import axios from 'axios';
import { TbReload, TbRobot, TbServer } from 'react-icons/tb';

const StatusChecker = () => {
    const [serverStatus, setServerStatus] = useState('offline');
    // const [databaseStatus, setDatabaseStatus] = useState('offline');
    const [botStatus, setBotStatus] = useState('offline');

    const checkServerStatus = async () => {
        try {
            const response = await axios.get('/api/');
            setServerStatus(response.status === 200 ? 'online' : 'offline');
        } catch (error) {
            setServerStatus('offline');
        }
    };

    // TODO: Implement database status checker, if possible without risks

    // const checkDatabaseStatus = async () => {
    //     try {
    //         const response = await axios.get('/');
    //         setDatabaseStatus(response.status === 200 ? 'online' : 'offline');
    //     } catch (error) {
    //         setDatabaseStatus('offline');
    //     }
    // };
    const checkAllStatuses = async () => {
        await Promise.all([checkServerStatus(), checkBotStatus()]);
    };

    const checkBotStatus = async () => {
        try {
            const response = await axios.get('/checkBot');
            setBotStatus(response.status === 200 ? 'online' : 'offline');
        } catch (error) {
            setBotStatus('offline');
        }
    };

    useEffect(() => {
        checkAllStatuses();
    }, []);

    return (
        <div className="flex flex-auto flex-row justify-center gap-6 h-10 p-2 w-fit bg-gray-800 rounded-md">
            <div className="flex flex-row gap-1 items-center ">
                Status:
                <TbServer size="24px" color={serverStatus === 'online' ? 'green' : 'red'} />
                |
                <TbRobot size="24px" color={botStatus === 'online' ? 'green' : 'red'} />
                <br/>
                <button onClick={checkAllStatuses}><TbReload /></button>
            </div>
        </div>
    );
};

export default StatusChecker;