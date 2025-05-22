import React, { useState } from "react";

/**
 * A simple review form component that captures user feedback and rating.
 * @param {Function} onSubmit - Callback invoked with { review, rating } on form submit.
 */
function ReviewForm({ onSubmit }) {
  const [review, setReview] = useState(""); // Text area input
  const [rating, setRating] = useState(5);  // Default rating value

  /**
   * Handles form submission
   * Calls parent handler and resets form fields
   */
  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit({ review, rating });
    setReview("");     // Clear input after submit
    setRating(5);      // Reset rating to default
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-3">
      {/* Review text area */}
      <textarea
        className="w-full p-2 border rounded"
        placeholder="Write your review..."
        value={review}
        onChange={(e) => setReview(e.target.value)}
        required
      />

      {/* Rating dropdown (1–5) */}
      <select
        value={rating}
        onChange={(e) => setRating(Number(e.target.value))}
        className="border p-2 rounded"
      >
        {[1, 2, 3, 4, 5].map((star) => (
          <option key={star} value={star}>
            {star} Star{star > 1 && "s"}
          </option>
        ))}
      </select>

      {/* Submit button */}
      <button
        type="submit"
        className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700"
      >
        Submit
      </button>
    </form>
  );
}

export default ReviewForm;
