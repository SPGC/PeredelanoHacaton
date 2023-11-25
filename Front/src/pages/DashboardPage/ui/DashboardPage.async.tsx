import {lazy, Suspense} from "react";
import DashboardPageSkeleton from "./DashboardPage.skeleton.tsx";

const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

const DashboardPageAsync = lazy(() =>
    delay(800)
        .then(() => import('./DashboardPage.tsx'))
);

const DashboardPageWithSkeleton = () => (
    <Suspense fallback={<DashboardPageSkeleton/>}>
        <DashboardPageAsync />
    </Suspense>
);

export default DashboardPageWithSkeleton;