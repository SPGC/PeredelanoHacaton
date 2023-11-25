import {lazy, Suspense} from "react";
import ListPageSkeleton from "./ListPage.skeleton.tsx";

const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

const ListPageAsync = lazy(() =>
    delay(800)
        .then(() => import('./ListPage.tsx'))
);

const ListPageWithSkeleton = () => (
    <Suspense fallback={<ListPageSkeleton/>}>
        <ListPageAsync />
    </Suspense>
);

export default ListPageWithSkeleton;