import React, { useState, useEffect } from 'react';
import { useAuth } from '../context/AuthContext';

// Checkout component handles cart review and order submission
function Checkout() {
  // Form fields for user input
  const [form, setForm] = useState({ name: '', address: '', card: '' });

  // Track if order was successfully submitted
  const [submitted, setSubmitted] = useState(false);

  // User cart, with product details
  const [cart, setCart] = useState([]);

  // Authenticated user data
  const { user } = useAuth();

  /**
   * Fetches the authenticated user's cart and enriches it with product data
   */
  useEffect(() => {
    const fetchCart = async () => {
      // Guard: skip fetch if user is not loaded
      if (!user?.user_id || !user?.auth_token) return;

      try {
        // Step 1: Get cart items
        const res = await fetch(`http://localhost:8080/users/${user.user_id}/cart`, {
          method: 'GET',
          headers: {
            Authorization: user.auth_token,
          },
        });

        if (!res.ok) {
          const errorText = await res.text();
          console.error('Failed to fetch cart:', errorText);
          return;
        }

        const data = await res.json();

        // Step 2: Enrich each item with product name & price
        const detailedCart = await Promise.all(
          (Array.isArray(data) ? data : []).map(async (item) => {
            try {
              const productRes = await fetch(`http://localhost:8080/products/${item.product_id}`);
              if (!productRes.ok) throw new Error('Product not found');

              const product = await productRes.json();
              return {
                ...item,
                name: product.name,
                price: product.price,
              };
            } catch {
              return { ...item, price: 0 };
            }
          })
        );

        setCart(detailedCart);
      } catch (err) {
        console.error('Error fetching cart:', err);
      }
    };

    fetchCart();
  }, [user]);

  /**
   * Updates form state as user types
   */
  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  /**
   * Submits the order to the backend
   */
  const handleSubmit = async (e) => {
    e.preventDefault();

    // Validate fields
    if (!form.name || !form.address || !form.card || cart.length === 0) {
      alert('Please complete all fields and add items to cart.');
      return;
    }

    // Build simplified order payload
    const orderItems = cart.map((item) => ({
      product_id: item.product_id,
      quantity: item.quantity || 1,
    }));

    try {
      const res = await fetch('http://localhost:8080/orders', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: user.auth_token,
        },
        body: JSON.stringify(orderItems),
      });

      if (res.ok) {
        setSubmitted(true);
        setCart([]);
      } else {
        const error = await res.text();
        console.error('Order failed:', error);
        alert(`Failed to place order: ${error}`);
      }
    } catch (err) {
      console.error('Network error placing order:', err);
      alert('Order error. Check backend connection.');
    }
  };

  // Calculate total cost of cart
  const total = cart.reduce(
    (sum, item) => sum + (item.price || 0) * (item.quantity || 1),
    0
  );

  return (
    <div className="p-4 max-w-xl mx-auto">
      <h2 className="text-xl font-bold mb-4">Checkout</h2>

      {/* Show confirmation after successful order */}
      {submitted ? (
        <div className="bg-green-100 text-green-800 p-3 rounded">
          Order placed successfully!
        </div>
      ) : (
        // Order form
        <form onSubmit={handleSubmit} className="space-y-4">
          <input
            type="text"
            name="name"
            placeholder="Full Name"
            className="border p-2 w-full"
            value={form.name}
            onChange={handleChange}
            required
          />
          <input
            type="text"
            name="address"
            placeholder="Shipping Address"
            className="border p-2 w-full"
            value={form.address}
            onChange={handleChange}
            required
          />
          <input
            type="text"
            name="card"
            placeholder="Card Number"
            className="border p-2 w-full"
            value={form.card}
            onChange={handleChange}
            required
          />

          <div className="text-right font-semibold text-lg">
            Total: {total.toFixed(2)} NOK
          </div>

          <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded w-full">
            Place Order
          </button>
        </form>
      )}
    </div>
  );
}

export default Checkout;
