import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import {BrowserRouter} from "react-router-dom";
import {ErrorBoundary} from "react-error-boundary";
import PageError from "./pages/PageError/PageError.tsx";

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
      <BrowserRouter>
          <ErrorBoundary fallback={<PageError />}>
            <App />
          </ErrorBoundary>
      </BrowserRouter>
  </React.StrictMode>,
)
