import React, { useState } from "react";

// Form component for user registration
function RegisterForm() {
  // Form state initialization
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
    first_name: "",
    last_name: "",
    address: "",
  });

  // Message state for feedback (success or error)
  const [message, setMessage] = useState("");

  /**
   * Handle input field changes
   * Updates the corresponding field in formData
   */
  const handleChange = (e) => {
    setFormData((prev) => ({
      ...prev,
      [e.target.name]: e.target.value,
    }));
  };

  /**
   * Handle form submission to register user
   */
  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const res = await fetch("http://localhost:8080/users", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData),
      });

      const text = await res.text();

      if (res.status === 201) {
        setMessage("Registration successful!");
      } else {
        setMessage(`Failed: ${text}`);
      }
    } catch (err) {
      console.error("Registration error:", err);
      setMessage("An unexpected error occurred.");
    }
  };

  return (
    <div className="max-w-md mx-auto bg-white shadow-md rounded-lg p-8">
      <h2 className="text-2xl font-bold mb-6 text-center text-gray-800">
        Register
      </h2>

      {/* Registration Form */}
      <form onSubmit={handleSubmit} className="space-y-4">
        {[
          { name: "username", label: "Username" },
          { name: "email", label: "Email", type: "email" },
          { name: "password", label: "Password", type: "password" },
          { name: "first_name", label: "First Name" },
          { name: "last_name", label: "Last Name" },
          { name: "address", label: "Address" },
        ].map(({ name, label, type = "text" }) => (
          <div key={name}>
            <label
              htmlFor={name}
              className="block text-sm font-medium text-gray-700 mb-1"
            >
              {label}
            </label>
            <input
              type={type}
              name={name}
              id={name}
              value={formData[name]}
              onChange={handleChange}
              required
              className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
        ))}

        <button
          type="submit"
          className="w-full bg-blue-600 text-white py-2 rounded-md font-semibold hover:bg-blue-700 transition"
        >
          Register
        </button>
      </form>

      {/* Feedback Message */}
      {message && (
        <p className="mt-4 text-sm text-center text-red-500">{message}</p>
      )}
    </div>
  );
}

export default RegisterForm;
