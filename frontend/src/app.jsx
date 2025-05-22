// React core imports
import React from "react";
import { Routes, Route, NavLink } from "react-router-dom";

// App page components
import Home from "./pages/Home";
import ProductDetail from "./pages/ProductDetail";
import Cart from "./pages/Cart";
import Checkout from "./pages/Checkout";
import AdminDashboard from "./pages/AdminDashboard";
import OrderHistory from "./pages/OrderHistory";
import Login from "./pages/Login";
import Profile from "./pages/Profile";
import Register from "./pages/Register";

// Context for authentication state
import { useAuth } from "./context/AuthContext";
import RegisterForm from "./pages/Register";

// Main App Component
function App() {
  // Extract user and logout function from auth context
  const { user, logout } = useAuth();

  return (
    <div className="bg-gray-50 min-h-screen font-sans">
      {/* App header with branding and navigation */}
      <header className="bg-blue-700 text-white shadow">
        <div className="container mx-auto px-6 py-4 flex justify-between items-center">
          <h1 className="text-2xl font-bold tracking-tight">ElectroMart</h1>

          {/* Main Navigation Links */}
          <nav className="space-x-6">
            <NavLink to="/">Home</NavLink>
            <NavLink to="/cart">Cart</NavLink>
            <NavLink to="/admin">Admin</NavLink>

            {/* Conditionally show profile or login links based on auth status */}
            {user ? (
              <>
                <NavLink to="/profile">{user.first_name}</NavLink>
                <NavLink to="/orders">Orders</NavLink>
                <button 
                  onClick={logout} 
                  className="text-red-300 ml-2"
                >
                  Logout
                </button>
              </>
            ) : (
              <NavLink to="/login">Login</NavLink>
            )}
          </nav>
        </div>
      </header>

      {/* Main content rendered based on route */}
      <main className="container mx-auto px-6 py-8">
        <Routes>
          {/* Public Routes */}
          <Route path="/" element={<Home />} />
          <Route path="/product/:id" element={<ProductDetail />} />
          <Route path="/cart" element={<Cart />} />
          <Route path="/checkout" element={<Checkout />} />

          {/* Admin and Authenticated Routes */}
          <Route path="/admin" element={<AdminDashboard />} />
          <Route path="/orders" element={<OrderHistory />} />
          <Route path="/profile" element={<Profile />} />

          {/* Authentication Routes */}
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
        </Routes>
      </main>
    </div>
  );
}

export default App;
