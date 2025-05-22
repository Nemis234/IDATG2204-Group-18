import React from "react";

/**
 * Displays a single item in the cart with quantity and total price
 * @param {Object} props
 * @param {Object} props.item - The cart item
 * @param {string} props.item.name - Product name
 * @param {number} props.item.price - Price per unit
 * @param {number} props.item.quantity - Quantity selected
 */
function CartItem({ item }) {
  return (
    <div className="border p-3 rounded mb-3 flex justify-between items-center">
      {/* Product name and pricing breakdown */}
      <div>
        <h3 className="font-semibold">{item.name}</h3>
        <p className="text-sm text-gray-600">
          {item.quantity} × {item.price} NOK
        </p>
      </div>

      {/* Total price for this line item */}
      <div className="text-right font-bold">
        {(item.quantity * item.price).toFixed(2)} NOK
      </div>
    </div>
  );
}

export default CartItem;
