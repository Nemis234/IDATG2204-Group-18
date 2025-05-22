// Core imports
import React, { useState } from "react";
import { useNavigate, NavLink } from "react-router-dom";

// Auth context for login state
import { useAuth } from "../context/AuthContext";

// Login page component
function Login() {
  // Form input state
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  // Error handling state
  const [error, setError] = useState("");

  // Navigation hook
  const navigate = useNavigate();

  // Login method from auth context
  const { login } = useAuth();

  // Handles login form submission
  const handleSubmit = async (e) => {
    e.preventDefault(); // Prevent default form behavior
    setError(""); // Clear any previous error messages

    try {
      // Step 1: Authenticate user and receive token/user ID
      const res = await fetch("http://localhost:8080/users/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password })
      });

      // If authentication fails
      if (!res.ok) {
        setError("Invalid credentials");
        return;
      }

      const authData = await res.json(); // Contains user_id and auth_token

      // Step 2: Fetch user profile with token
      const profileRes = await fetch(`http://localhost:8080/users/${authData.user_id}`, {
        headers: { Authorization: authData.auth_token }
      });

      // If profile fetch fails
      if (!profileRes.ok) {
        setError("Failed to load profile");
        return;
      }

      const profile = await profileRes.json();

      // Step 3: Merge login and profile data and store in context
      const fullUser = { ...authData, ...profile };
      login(fullUser); // Update global auth context

      // Redirect to homepage
      navigate("/");
    } catch (err) {
      // Catch and log unexpected errors
      console.error("Login error:", err);
      setError("Server error. Try again later.");
    }
  };

  return (
    <div className="p-6 max-w-md mx-auto">
      <h2 className="text-2xl font-bold mb-4">Login</h2>

      {/* Display error message if present */}
      {error && <p className="text-red-600 mb-2">{error}</p>}

      {/* Login form */}
      <form onSubmit={handleSubmit} className="space-y-4">
        <input
          type="email"
          placeholder="Email"
          className="w-full border p-2 rounded"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Password"
          className="w-full border p-2 rounded"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button 
          className="w-full bg-blue-600 text-white py-2 rounded" 
          type="submit"
        >
          Login
        </button>
      </form>

      {/* Navigation to registration page */}
      <div className="text-center mt-4">
        <span className="text-sm text-gray-600">Don’t have an account?</span>{" "}
        <NavLink to="/register" className="text-blue-500 hover:underline">
          Register
        </NavLink>
      </div>
    </div>
  );
}

export default Login;
