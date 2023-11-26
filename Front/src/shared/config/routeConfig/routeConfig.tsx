import { RouteProps } from 'react-router-dom';
import {MainPage} from "../../../pages/MainPage";
import {DashboardPage} from "../../../pages/DashboardPage";
import {ListPage} from "../../../pages/ListPage";
import {NotFoundPage} from "../../../pages/NotFoundPage";
import {RobotPage} from "../../../pages/RobotPage";


export enum AppRoutes {
    MAIN = 'main',
    DASHBOARD = 'dashboard',
    NOT_FOUND = 'not_found',
    LIST = 'list',
    ROBOT = 'robot',
}

export const RoutePath: Record<AppRoutes, string> = {
    [AppRoutes.MAIN]: '/',
    [AppRoutes.DASHBOARD]: '/dashboard',
    [AppRoutes.LIST]: '/list',
    [AppRoutes.ROBOT]: '/robot',

    // 404
    [AppRoutes.NOT_FOUND]: '*',
};

export const routeConfig: Record<AppRoutes, RouteProps> = {
    [AppRoutes.MAIN]: {
        path: RoutePath.main,
        element: <MainPage />,
    },
    [AppRoutes.DASHBOARD]: {
        path: RoutePath.dashboard,
        element: <DashboardPage />,
    },
    [AppRoutes.LIST]: {
        path: RoutePath.list,
        element: <ListPage />,
    },
    [AppRoutes.ROBOT]: {
        path: RoutePath.robot,
        element: <RobotPage />,
    },
    [AppRoutes.NOT_FOUND]: {
        path: RoutePath.not_found,
        element: <NotFoundPage />,
    },
};