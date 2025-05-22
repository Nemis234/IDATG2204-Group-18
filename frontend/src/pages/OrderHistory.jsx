import React, { useEffect, useState } from 'react';
import { useAuth } from '../context/AuthContext';

// OrderHistory displays a list of past orders for the authenticated user
function OrderHistory() {
  const { user } = useAuth();

  // Local state to hold order list
  const [orders, setOrders] = useState([]);

  /**
   * Fetches order history from backend when user is authenticated
   */
  useEffect(() => {
    const fetchOrders = async () => {
      if (!user?.user_id || !user?.auth_token) return;

      try {
        const res = await fetch(`http://localhost:8080/users/${user.user_id}/orders`, {
          headers: {
            Authorization: user.auth_token
          }
        });

        if (!res.ok) {
          const errorText = await res.text();
          console.error("Failed to load orders:", errorText);
          return;
        }

        const data = await res.json();

        // Store orders only if valid array
        setOrders(Array.isArray(data) ? data : []);
      } catch (err) {
        console.error("Failed to load orders:", err);
      }
    };

    fetchOrders();
  }, [user]);

  return (
    <div className="p-6 max-w-3xl mx-auto">
      <h2 className="text-2xl font-bold mb-4">Order History</h2>

      {/* If no orders are present */}
      {orders.length === 0 ? (
        <p className="text-gray-600">No orders found.</p>
      ) : (
        <ul className="grid grid-cols-1 gap-4">
          {/* Render each order summary */}
          {orders.map((order) => (
            <li 
              key={order.order_id} 
              className="border rounded-lg p-4 shadow-sm hover:shadow-md transition-all"
            >
              <div className="flex justify-between items-center mb-2">
                <span className="font-semibold">Order #{order.order_id}</span>
                
                {/* Conditional color based on order status */}
                <span
                  className={`px-2 py-1 rounded text-sm ${
                    order.order_status === "Shipped"
                      ? "bg-blue-100 text-blue-800"
                      : order.order_status === "Completed"
                      ? "bg-green-100 text-green-800"
                      : "bg-gray-100 text-gray-800"
                  }`}
                >
                  {order.order_status}
                </span>
              </div>

              {/* Order details: date and total */}
              <div className="text-sm text-gray-700">
                <p>
                  <strong>Date:</strong>{" "}
                  {new Date(order.order_date).toLocaleDateString()}
                </p>
                <p>
                  <strong>Total:</strong>{" "}
                  {order.order_total || "N/A"} NOK
                </p>
              </div>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default OrderHistory;
