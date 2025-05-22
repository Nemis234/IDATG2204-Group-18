// File: src/components/ProductCard.jsx

import React from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext";

/**
 * Displays a product's basic information and actions.
 * Image has been removed due to unsupported backend.
 *
 * @param {Object} props
 * @param {Object} props.product - The product data to render
 */
function ProductCard({ product }) {
  const navigate = useNavigate();
  const { user } = useAuth();

  /**
   * Sends a request to add one unit of the product to the user's cart
   */
  const handleAddToCart = async () => {
    if (!user?.user_id || !user?.auth_token) {
      alert("Please log in to add items to your cart.");
      return;
    }

    try {
      const res = await fetch(`http://localhost:8080/users/${user.user_id}/cart`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: user.auth_token,
        },
        body: JSON.stringify({
          product_id: product.product_id,
          quantity: 1,
        }),
      });

      if (!res.ok) {
        const errMsg = await res.text();
        console.error("Add to cart failed:", errMsg);
        alert("Failed to add to cart.");
        return;
      }

      alert(`${product.name} added to your cart!`);
    } catch (err) {
      console.error("Network error:", err);
      alert("Error adding product to cart.");
    }
  };

  return (
    <div className="border rounded-lg p-4 hover:shadow-lg transition">
      {/* Product name */}
      <h2 className="text-lg font-semibold mb-1">{product.name}</h2>

      {/* Product description */}
      <p className="text-sm text-gray-600 mb-1">
        {product.description}
      </p>

      {/* Stock status */}
      <p className="text-sm text-gray-500 mb-1">
        In stock: {product.stock_quantity}
      </p>

      {/* Product price */}
      <p className="text-green-700 font-bold mb-2">
        {product.price} NOK
      </p>

      {/* Action buttons */}
      <div className="flex justify-between">
        <button
          className="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700"
          onClick={() => navigate(`/product/${product.product_id}`)}
        >
          Write review
        </button>
        <button
          className="bg-green-600 text-white px-3 py-1 rounded hover:bg-green-700"
          onClick={handleAddToCart}
        >
          Add to Cart
        </button>
      </div>
    </div>
  );
}

export default ProductCard;
