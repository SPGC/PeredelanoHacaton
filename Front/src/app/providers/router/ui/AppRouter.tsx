import { Route, Routes } from 'react-router-dom';
import { Suspense } from 'react';
import { routeConfig } from '../../../../shared/config/routeConfig/routeConfig';
import {RegularPageSkeleton} from "../../../../shared/ui/Skeletons/ui/Skeletons.tsx";

const AppRouter = () => (

    <Routes>
        {Object.values(routeConfig).map(({ element, path }) => (
            <Route
                key={path}
                path={path}
                element={(
                    <Suspense fallback={<RegularPageSkeleton />}>
                                {element}
                    </Suspense>
                )}
            />
        ))}
    </Routes>

);

export default AppRouter;