import { clsx } from "clsx";
import { TbDashboard, TbList, TbRobot, TbSmartHome } from "react-icons/tb";
import { FaHandshakeAngle } from "react-icons/fa6";
import {Link} from "react-router-dom";

const ICON_SIZE = "28px";
const ICON_COLOR = "white";

export interface SidebarProps {
    isOpen: boolean;
}

const Sidebar = ({ isOpen = true }: SidebarProps) => {
    return (
        <div className={clsx("bg-gray-800 w-16 flex flex-col py-5 items-center", { 'hidden': !isOpen })}>
            <button className="text-xs" aria-label="Handshake">
                <FaHandshakeAngle color={ICON_COLOR} size="32px"/>
            </button>
            <div className="flex flex-col gap-4 mt-8">
                <Link aria-label="Home" to={'/'}>
                    <TbSmartHome className="text-white" size={ICON_SIZE} />
                </Link>
                <Link aria-label="Dashboard" to={'/dashboard'}>
                    <TbDashboard className="text-white" size={ICON_SIZE}/>
                </Link>
                <Link aria-label="List" to={'/list'}>
                    <TbList className="text-white" size={ICON_SIZE}/>
                </Link>
                <Link aria-label="Robot" to={'/robot'}>
                    <TbRobot className="text-white" size={ICON_SIZE} />
                </Link>
            </div>
        </div>
    );
};

export default Sidebar;