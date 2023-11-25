import MainPageSkeleton from "./MainPage.skeleton.tsx";
import {lazy, Suspense} from "react";

const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

const MainPageAsync = lazy(() =>
    delay(800)
        .then(() => import('./MainPage'))
);

const MainPageWithSkeleton = () => (
    <Suspense fallback={<MainPageSkeleton />}>
        <MainPageAsync />
    </Suspense>
);

export default MainPageWithSkeleton;