import { clsx } from "clsx";
import {TbDashboard, TbList, TbLogout, TbRobot, TbSmartHome} from "react-icons/tb";
import { FaHandshakeAngle } from "react-icons/fa6";
import { Link } from "react-router-dom";
import { useLocation } from 'react-router-dom'
import { SignOutButton } from "@clerk/clerk-react";


const ICON_SIZE = "28px";
const ICON_COLOR = "white";

export interface SidebarProps {
    isOpen: boolean;
}

const Sidebar = ({ isOpen = true }: SidebarProps) => {
    const location = useLocation();

    const isActive = (path: string) => location.pathname === path;

    return (
        <div className={clsx("bg-gray-800 justify-between w-16 flex flex-col py-5 items-center", { 'hidden': !isOpen })}>
            <div className="flex flex-col items-center w-20">
                <div className="text-xs" aria-label="Handshake">
                    <FaHandshakeAngle color={ICON_COLOR} size="36px"/>
                </div>
                <div className="flex flex-col gap-4 mt-8">
                    <Link aria-label="Home" to={'/'}>
                        <TbSmartHome className={clsx("rounded-md text-white", { 'border-b-2': isActive('/') })} size={ICON_SIZE} />
                    </Link>
                    <Link aria-label="Dashboard"  to={'/dashboard'}>
                        <TbDashboard className={clsx("rounded-md text-white", { 'border-b-2': isActive('/dashboard') })} size={ICON_SIZE}/>
                    </Link>
                    <Link aria-label="List" to={'/list'}>
                        <TbList className={clsx("rounded-md text-white", { 'border-b-2': isActive('/list') })} size={ICON_SIZE}/>
                    </Link>
                    <Link aria-label="Robot" to={'/robot'}>
                        <TbRobot className={clsx("rounded-md text-white", { 'border-b-2': isActive('/robot') })} size={ICON_SIZE} />
                    </Link>
                </div>
            </div>

            <div className="flex justify-center">
                <SignOutButton><TbLogout size="28px" color="white" /></SignOutButton>
            </div>
        </div>
    );
};

export default Sidebar;