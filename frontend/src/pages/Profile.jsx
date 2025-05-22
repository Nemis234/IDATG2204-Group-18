import React, { useState, useEffect } from "react";
import { useAuth } from "../context/AuthContext";

// Profile page allows user to view and update their account details
function Profile() {
  const { user, updateProfile } = useAuth();

  // Form state stores editable user fields
  const [form, setForm] = useState({});

  // Tracks whether update was successfully submitted
  const [updated, setUpdated] = useState(false);

  /**
   * On user load, initialize form with user data
   * Password is not populated for security reasons
   */
  useEffect(() => {
    if (user) {
      setForm({ ...user, password: "" });
    }
  }, [user]);

  // Guard: If user isn't loaded yet
  if (!user) {
    return <div className="text-center mt-10">Loading user profile...</div>;
  }

  /**
   * Handle form input changes
   */
  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  /**
   * Handle profile update form submission
   */
  const handleSubmit = (e) => {
    e.preventDefault();

    const payload = { ...form };

    // Exclude password if left blank
    if (!form.password) {
      delete payload.password;
    }

    // Call global profile updater
    updateProfile(payload);

    // Show success message temporarily
    setUpdated(true);
    setTimeout(() => setUpdated(false), 3000);
  };

  return (
    <div className="p-6 max-w-lg mx-auto bg-white shadow-md rounded-lg">
      <h2 className="text-3xl font-bold mb-4 text-center">
        {user.first_name} {user.last_name}
      </h2>

      {/* Update success feedback */}
      {updated && (
        <p className="text-green-600 text-sm mb-4 text-center">
          Profile updated successfully!
        </p>
      )}

      {/* Profile form */}
      <form onSubmit={handleSubmit} className="space-y-4">
        <input
          name="first_name"
          value={form.first_name || ""}
          onChange={handleChange}
          className="w-full border p-2 rounded"
          placeholder="First Name"
        />
        <input
          name="last_name"
          value={form.last_name || ""}
          onChange={handleChange}
          className="w-full border p-2 rounded"
          placeholder="Last Name"
        />
        <input
          name="email"
          value={form.email || ""}
          onChange={handleChange}
          className="w-full border p-2 rounded"
          placeholder="Email"
        />
        <input
          name="address"
          value={form.address || ""}
          onChange={handleChange}
          className="w-full border p-2 rounded"
          placeholder="Address"
        />
        <input
          type="password"
          name="password"
          value={form.password || ""}
          onChange={handleChange}
          className="w-full border p-2 rounded"
          placeholder="New Password (leave blank to keep existing)"
        />

        <button
          type="submit"
          className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700"
        >
          Save Changes
        </button>
      </form>
    </div>
  );
}

export default Profile;
