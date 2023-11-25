import './App.css'
import Sidebar from "./widgets/Sidebar";
import { Route, Routes } from "react-router-dom";
import MainPage from "./pages/MainPage/MainPage.tsx";
import DashboardPage from "./pages/DashboardPage/DashboardPage.tsx";
import ListPage from "./pages/ListPage/ListPage.tsx";
import RobotPage from "./pages/RobotPage/RobotPage.tsx";
import {ClerkProvider, RedirectToSignIn, SignedIn, SignedOut} from "@clerk/clerk-react";

function App() {
    if (!import.meta.env.VITE_REACT_APP_CLERK_PUBLISHABLE_KEY) {
        throw new Error("Missing Publishable Key")
    }
    const clerkPubKey = import.meta.env.VITE_REACT_APP_CLERK_PUBLISHABLE_KEY;

  return (
    <>
    <ClerkProvider publishableKey={clerkPubKey}>
      <div className="bg-gray-900 text-4xl flex flex-row">
                    <Sidebar isOpen={true} />
                    <div className="w-full p-3 h-screen text-white">
                        <SignedIn>
                            <Routes>
                                <Route path="/" element={<MainPage />} />
                                <Route path="/dashboard" element={<DashboardPage />} />
                                <Route path="/list" element={<ListPage />} />
                                <Route path="/robot" element={<RobotPage />} />
                            </Routes>
                        </SignedIn>
                        <SignedOut>
                            <RedirectToSignIn />
                        </SignedOut>

                    </div>
      </div>
    </ClerkProvider>
    </>
  )
}

export default App