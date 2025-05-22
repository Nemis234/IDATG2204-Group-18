// File: src/pages/ProductDetail.jsx
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import ReviewForm from "../components/ReviewForm";
import { useAuth } from '../context/AuthContext';

// ProductDetail page displays detailed info about a selected product
function ProductDetail() {
  // Extract product ID from route params
  const { id } = useParams();

  // Auth context to access current user
  const { user } = useAuth();

  // Product state
  const [product, setProduct] = useState(null);

  /**
   * Handle submission of product reviews
   * @param {Object} reviewData - Contains rating and review text
   */
  const handleReviewSubmit = async (reviewData) => {
    console.log("Review submitted:", reviewData);

    // Guard: prevent unauthenticated review submission
    if (!user?.user_id || !user?.auth_token) {
      alert("You must be logged in to submit a review.");
      return;
    }

    try {
      await fetch(`http://localhost:8080/products/${id}/reviews`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": user.auth_token
        },
        body: JSON.stringify({
          user_id: user.user_id,
          rating: reviewData.rating,
          comment: reviewData.review
        })
      });

      alert("Review submitted!");
    } catch (err) {
      console.error("Failed to submit review:", err);
      alert("Could not submit review. Please try again.");
    }
  };

  /**
   * Fetch product details on mount or when ID changes
   */
  useEffect(() => {
    fetch(`http://localhost:8080/products/${id}`)
      .then(res => res.json())
      .then(data => setProduct(data))
      .catch(err => console.error("Failed to fetch product details:", err));
  }, [id]);

  // Show loading state while fetching product data
  if (!product) return <p>Loading...</p>;

  return (
    <div className="p-4">
      <h2 className="text-xl font-bold">{product.name}</h2>

      {/* Product description and price */}
      <p>{product.description}</p>
      <p className="text-lg font-semibold mt-2">
        {product.price} NOK
      </p>

      {/* Review submission section */}
      <h3 className="text-lg font-semibold mt-6 mb-2">Leave a Review</h3>
      <ReviewForm onSubmit={handleReviewSubmit} />
    </div>
  );
}

export default ProductDetail;
