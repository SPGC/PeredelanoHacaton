import './App.css'
import Sidebar from "./widgets/Sidebar";
import { Route, Routes } from "react-router-dom";
import MainPage from "./pages/MainPage/MainPage.tsx";
import DashboardPage from "./pages/DashboardPage/DashboardPage.tsx";
import ListPage from "./pages/ListPage/ListPage.tsx";
import RobotPage from "./pages/RobotPage/RobotPage.tsx";

function App() {

  return (
    <>
      <div className="bg-gray-900 text-4xl flex flex-row">

          <Sidebar isOpen={true} />
          <div className="w-full p-3 h-screen text-white">
              <Routes>
                  <Route path="/" element={<MainPage />} />
                  <Route path="/dashboard" element={<DashboardPage />} />
                  <Route path="/list" element={<ListPage />} />
                  <Route path="/robot" element={<RobotPage />} />
              </Routes>
          </div>

      </div>
    </>
  )
}

export default App
