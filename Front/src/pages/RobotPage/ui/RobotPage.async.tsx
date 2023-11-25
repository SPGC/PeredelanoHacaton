import {lazy, Suspense} from "react";
import RobotPageSkeleton from "./RobotPage.skeleton.tsx";

const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

const RobotPageAsync = lazy(() =>
    delay(800)
        .then(() => import('./RobotPage.tsx'))
);

const RobotPageWithSkeleton = () => (
    <Suspense fallback={<RobotPageSkeleton/>}>
        <RobotPageAsync />
    </Suspense>
);

export default RobotPageWithSkeleton;