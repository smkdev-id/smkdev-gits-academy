import React from 'react'
import ReactDOM from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import './index.css'
import Home from './Pages/Home.jsx';
import Profile from './Pages/Profile.jsx';
import App from './App.jsx';

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <Home />,
  },
  {
    path: "/profile",
    element: <Profile />,
    errorElement: <Profile />,
  },
  {
    path: "/blogs",
    element: <Profile />,
    errorElement: <Profile />,
  },
]);


ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
