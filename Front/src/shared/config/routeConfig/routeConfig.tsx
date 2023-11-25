import { RouteProps } from 'react-router-dom';
import {MainPage} from "../../../pages/MainPage";
import {DashboardPage} from "../../../pages/DashboardPage";
import {ListPage} from "../../../pages/ListPage";
import {NotFoundPage} from "../../../pages/NotFoundPage";
import {RobotPage} from "../../../pages/RobotPage";


export enum AppRoutes {
    MAIN = 'main',
    DASHBOARD = 'about',
    NOT_FOUND = 'not_found',
    LIST = 'machinery',
    ROBOT = 'fields',
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
        path: RoutePath.about,
        element: <DashboardPage />,
    },
    [AppRoutes.LIST]: {
        path: RoutePath.machinery,
        element: <ListPage />,
    },
    [AppRoutes.ROBOT]: {
        path: RoutePath.fields,
        element: <RobotPage />,
    },
    [AppRoutes.NOT_FOUND]: {
        path: RoutePath.not_found,
        element: <NotFoundPage />,
    },
};