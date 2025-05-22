import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { useAuth } from "../context/AuthContext";

// Cart page component
function Cart() {
  // Access current user from auth context
  const { user } = useAuth();

  // Local state for cart items
  const [cartItems, setCartItems] = useState([]);

  /**
   * Fetch cart contents and product details on mount or user change
   */
  useEffect(() => {
    const fetchCart = async () => {
      if (!user?.user_id || !user?.auth_token) return;

      try {
        // Step 1: Fetch user's cart
        const res = await fetch(`http://localhost:8080/users/${user.user_id}/cart`, {
          method: "GET",
          headers: {
            Authorization: user.auth_token,
          },
        });

        if (!res.ok) {
          const errorText = await res.text();
          console.error("Failed to fetch cart:", errorText);
          return;
        }

        const rawCart = await res.json();

        // Validate that we have an array
        if (!Array.isArray(rawCart)) {
          console.error("Expected cart array, got:", rawCart);
          return;
        }

        // Step 2: Enrich cart items with product name/price
        const cartWithDetails = await Promise.all(
          rawCart.map(async (item) => {
            try {
              const productRes = await fetch(`http://localhost:8080/products/${item.product_id}`);
              if (!productRes.ok) throw new Error("Product not found");

              const product = await productRes.json();

              return {
                ...item,
                name: product.name || "Unnamed",
                price: product.price || 0,
              };
            } catch {
              // If product fetch fails, default fields
              return { ...item, name: "Unnamed", price: 0 };
            }
          })
        );

        setCartItems(cartWithDetails);
      } catch (err) {
        console.error("Network error while fetching cart:", err);
      }
    };

    fetchCart();
  }, [user]);

  /**
   * Handle removing a product from the cart
   */
  const handleRemove = async (productId) => {
    try {
      const res = await fetch(`http://localhost:8080/users/${user.user_id}/cart/${productId}`, {
        method: "DELETE",
        headers: {
          Authorization: user.auth_token,
        },
      });

      if (!res.ok) {
        const err = await res.text();
        console.error("Failed to delete item:", err);
        return;
      }

      // Remove item from local state
      setCartItems((prev) => prev.filter((item) => item.product_id !== productId));
    } catch (err) {
      console.error("Network error while deleting cart item:", err);
    }
  };

  /**
   * Handle quantity input changes
   */
  const handleQuantityChange = async (productId, value) => {
    const newQuantity = Math.max(1, Number(value) || 1); // Fallback to 1 if invalid

    // Update local state optimistically
    const updated = cartItems.map((item) =>
      item.product_id === productId ? { ...item, quantity: newQuantity } : item
    );
    setCartItems(updated);

    try {
      // Update on server
      const res = await fetch(`http://localhost:8080/users/${user.user_id}/cart/${productId}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: user.auth_token,
        },
        body: JSON.stringify({ quantity: newQuantity }),
      });

      if (!res.ok) {
        const err = await res.text();
        console.error("Failed to update quantity:", err);
      }
    } catch (err) {
      console.error("Error updating cart quantity:", err);
    }
  };

  // Calculate total cart price
  const total = cartItems.reduce(
    (sum, item) => sum + (item.price || 0) * (item.quantity || 1),
    0
  );

  return (
    <div className="p-6 max-w-3xl mx-auto">
      <h2 className="text-2xl font-bold mb-4">Your Cart</h2>

      {/* If cart is empty, display message and link */}
      {cartItems.length === 0 ? (
        <p className="text-gray-600">
          Your cart is empty.{" "}
          <Link to="/" className="text-blue-600 hover:underline">
            Browse Products
          </Link>
        </p>
      ) : (
        <div className="space-y-4">
          {/* Render each cart item */}
          {cartItems.map((item) => (
            <div
              key={item.product_id}
              className="border p-4 rounded flex justify-between items-center"
            >
              <div className="w-2/3">
                <h3 className="font-semibold">{item.name}</h3>
                <p className="text-sm text-gray-600 mb-2">
                  {item.price} NOK each
                </p>
                <div className="flex gap-2 items-center">
                  <label htmlFor={`qty-${item.product_id}`} className="text-sm">
                    Qty:
                  </label>
                  <input
                    type="number"
                    id={`qty-${item.product_id}`}
                    value={item.quantity}
                    min="1"
                    className="border p-1 w-16 text-center"
                    onChange={(e) =>
                      handleQuantityChange(item.product_id, e.target.value)
                    }
                  />
                  <button
                    onClick={() => handleRemove(item.product_id)}
                    className="ml-4 text-red-500 hover:underline text-sm"
                  >
                    Remove
                  </button>
                </div>
              </div>

              <div className="font-bold text-right w-1/3">
                {(item.price * item.quantity).toFixed(2)} NOK
              </div>
            </div>
          ))}

          {/* Cart total and checkout button */}
          <div className="border-t pt-4 text-right">
            <p className="text-lg font-bold">
              Total: {total.toFixed(2)} NOK
            </p>
            <Link
              to="/checkout"
              className="inline-block mt-4 bg-green-600 text-white px-6 py-2 rounded hover:bg-green-700"
            >
              Proceed to Checkout
            </Link>
          </div>
        </div>
      )}
    </div>
  );
}

export default Cart;
