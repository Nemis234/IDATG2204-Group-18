import React, { createContext, useState, useContext, useEffect } from "react";

// Create a new context for authentication
const AuthContext = createContext();

/**
 * Provides authentication state and actions to children components
 */
export const AuthProvider = ({ children }) => {
  // Global user state (null = not authenticated)
  const [user, setUser] = useState(null);

  /**
   * Load user from localStorage on initial render (persist login)
   */
  useEffect(() => {
    const storedUser = localStorage.getItem("user");
    if (storedUser) {
      setUser(JSON.parse(storedUser));
    }
  }, []);

  /**
   * Logs in a user by saving their credentials to state and localStorage
   * @param {Object} userData - user object from backend (incl. token)
   */
  const login = (userData) => {
    localStorage.setItem("user", JSON.stringify(userData));
    setUser(userData);
  };

  /**
   * Logs out user and clears state/localStorage
   */
  const logout = () => {
    localStorage.removeItem("user");
    setUser(null);
  };

  /**
   * Updates the user profile on the backend and in local state
   * @param {Object} updates - Fields to update (e.g. name, email, password)
   */
  const updateProfile = async (updates) => {
    // Guard: prevent update if not authenticated
    if (!user?.user_id || !user?.auth_token) return;

    try {
      const res = await fetch(`http://localhost:8080/users/${user.user_id}`, {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
          Authorization: user.auth_token,
        },
        body: JSON.stringify(updates),
      });

      if (res.status === 204) {
        // Merge updates into current user state
        const newUser = { ...user, ...updates };
        localStorage.setItem("user", JSON.stringify(newUser));
        setUser(newUser);
      } else {
        const error = await res.text();
        console.error("Update failed:", error);
      }
    } catch (err) {
      console.error("Error updating profile:", err);
    }
  };

  /**
   * Expose auth data and actions to consuming components
   */
  return (
    <AuthContext.Provider value={{ user, login, logout, updateProfile }}>
      {children}
    </AuthContext.Provider>
  );
};

/**
 * Hook to access authentication state and actions
 */
export const useAuth = () => useContext(AuthContext);
