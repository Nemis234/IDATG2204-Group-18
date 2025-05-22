// React core libraries for rendering the app
import React from "react";
import ReactDOM from "react-dom/client";

// React Router for client-side routing
import { BrowserRouter } from "react-router-dom";

// Main App component
import App from "./app";

// Global CSS styles
import "./styles/index.css";

// Authentication context provider
import { AuthProvider } from "./context/AuthContext";

// Mount the React app to the root DOM element
ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    {/* Global auth state provider for the entire app */}
    <AuthProvider>
      {/* Enables routing using HTML5 history API */}
      <BrowserRouter>
        {/* Root application component */}
        <App />
      </BrowserRouter>
    </AuthProvider>
  </React.StrictMode>
);
