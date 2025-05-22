import React, { useEffect, useState } from 'react';
import { useAuth } from '../context/AuthContext';

// Admin dashboard component
function AdminDashboard() {
  // Authenticated user
  const { user } = useAuth();

  // State declarations
  const [products, setProducts] = useState([]);
  const [form, setForm] = useState({});
  const [users, setUsers] = useState([]);
  const [brands, setBrands] = useState([]);
  const [categories, setCategories] = useState([]);
  const [newBrand, setNewBrand] = useState({ name: '', description: '' });
  const [newCategory, setNewCategory] = useState({ name: '', description: '' });
  const [selectedUser, setSelectedUser] = useState('');
  const [membershipUser, setMembershipUser] = useState('');
  const [membershipLevel, setMembershipLevel] = useState('');

  // Fetch all necessary admin data
  useEffect(() => {
    fetchData();
  }, [user]);

  // Load products, users, brands, and categories in parallel
  const fetchData = async () => {
    try {
      const [productsData, usersData, brandsData, categoriesData] = await Promise.all([
        fetch('http://localhost:8080/products?page=1').then(res => res.json()),
        fetch('http://localhost:8080/users', { headers: { Authorization: user.auth_token } }).then(res => res.json()),
        fetch('http://localhost:8080/brands').then(res => res.json()),
        fetch('http://localhost:8080/categories').then(res => res.json()),
      ]);
      setProducts(productsData || []);
      setUsers(usersData || []);
      setBrands(brandsData || []);
      setCategories(categoriesData || []);
    } catch (err) {
      console.error("Error fetching admin data:", err);
    }
  };

  // Handle form field changes for new product
  const handleInputChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  // Handle inline edit field change
  const handleEditChange = (id, field, value) => {
    setProducts(products.map(p =>
      p.product_id === id
        ? { ...p, [field]: field === "price" ? parseFloat(value) : parseInt(value) }
        : p
    ));
  };

  // Add a new product
  const handleAddProduct = async (e) => {
    e.preventDefault();

    const newProduct = {
      ...form,
      price: parseFloat(form.price),
      stock: parseInt(form.stock),
    };

    try {
      await fetch('http://localhost:8080/products', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json', Authorization: user.auth_token },
        body: JSON.stringify({
          name: newProduct.name,
          description: newProduct.description,
          price: newProduct.price,
          stock_quantity: newProduct.stock,
          category_name: newProduct.category,
          brand_name: newProduct.brand,
        }),
      });

      setForm({});
      fetchData();
    } catch (err) {
      console.error("Error adding product:", err);
    }
  };

  // Delete product
  const handleDelete = (id) => {
    fetch(`http://localhost:8080/products/${id}`, {
      method: 'DELETE',
      headers: { Authorization: user.auth_token }
    }).then(() => fetchData());
  };

  // Save edited product
  const handleSave = (id) => {
    const product = products.find(p => p.product_id === id);
    fetch(`http://localhost:8080/products/${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json', Authorization: user.auth_token },
      body: JSON.stringify({
        name: product.name,
        description: product.description,
        img_url: product.img_url,
        price: product.price,
        stock_quantity: product.stock_quantity,
        category_name: product.category_name,
        brand_name: product.brand_name,
      }),
    }).then(() => fetchData());
  };

  // Add a new brand
  const addBrand = async (e) => {
    e.preventDefault();

    const payload = {
      brand_name: newBrand.name,
      brand_desc: newBrand.description
    };

    await fetch('http://localhost:8080/brands', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: user.auth_token
      },
      body: JSON.stringify(payload)
    });

    setNewBrand({ name: '', description: '' });
    fetchData();
  };

  // Add a new category
  const addCategory = async (e) => {
    e.preventDefault();

    const payload = {
      category_name: newCategory.name,
      category_desc: newCategory.description
    };

    await fetch('http://localhost:8080/categories', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: user.auth_token
      },
      body: JSON.stringify(payload)
    });

    setNewCategory({ name: '', description: '' });
    fetchData();
  };

  // Update user role to admin
  const updateUserRole = async (role) => {
    if (!selectedUser) return;

    const res = await fetch(`http://localhost:8080/users/${selectedUser}`, {
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json', Authorization: user.auth_token },
      body: JSON.stringify({ role_name: role }),
    });

    if (res.status === 204) {
      alert(`User role updated to ${role}`);
      fetchData();
    } else {
      alert('Failed to update role');
    }
  };

  // Assign or update user membership level
  const updateMembership = async () => {
    const today = new Date().toISOString();

    await fetch(`http://localhost:8080/users/${membershipUser}/member`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json', Authorization: user.auth_token },
      body: JSON.stringify({
        membership_level: membershipLevel,
        membership_start: today
      }),
    });

    alert('Membership updated!');
  };

  return (
    <div className="p-6 max-w-6xl mx-auto">
      <h2 className="text-2xl font-bold mb-6">Admin Dashboard</h2>

      {/* ---------------------- User Role Management ---------------------- */}
      <div className="mb-8 bg-gray-100 p-4 rounded space-y-3">
        <h3 className="text-lg font-semibold">Manage User Roles</h3>
        <select value={selectedUser} onChange={(e) => setSelectedUser(e.target.value)} className="p-2 border rounded w-full">
          <option value="">-- Select User --</option>
          {users.map((u) => (
            <option key={u.user_id} value={u.user_id}>
              {u.username || u.email || u.user_id}
            </option>
          ))}
        </select>
        <div className="flex gap-4">
          <button disabled={!selectedUser} onClick={() => updateUserRole("admin")} className="bg-purple-600 text-white px-4 py-2 rounded disabled:opacity-50">Make Admin</button>
        </div>
      </div>

      {/* ---------------------- Membership Management ---------------------- */}
      <div className="mb-8 bg-gray-100 p-4 rounded space-y-3">
        <h3 className="text-lg font-semibold">Set Membership Level</h3>
        <select value={membershipUser} onChange={(e) => setMembershipUser(e.target.value)} className="p-2 border rounded w-full">
          <option value="">-- Select User --</option>
          {users.map((u) => (
            <option key={u.user_id} value={u.user_id}>
              {u.username || u.email || u.user_id}
            </option>
          ))}
        </select>
        <select value={membershipLevel} onChange={(e) => setMembershipLevel(e.target.value)} className="p-2 border rounded w-full">
          <option value="">-- Select Membership Level --</option>
          <option value="Silver">Silver</option>
          <option value="Gold">Gold</option>
          <option value="Platinum">Platinum</option>
        </select>
        <button disabled={!membershipUser || !membershipLevel} className="bg-indigo-600 text-white px-4 py-2 rounded disabled:opacity-50" onClick={updateMembership}>
          Save Membership
        </button>
      </div>

      {/* ---------------------- Brand & Category Insertion ---------------------- */}
      <div className="grid grid-cols-2 gap-6 mb-8">
        {/* Add Brand Form */}
        <form onSubmit={addBrand} className="bg-gray-100 p-4 rounded">
          <h3 className="text-lg font-semibold mb-2">Add Brand</h3>
          <input value={newBrand.name} onChange={(e) => setNewBrand({ ...newBrand, name: e.target.value })} placeholder="Brand Name" required className="w-full p-2 border rounded mb-2" />
          <textarea value={newBrand.description} onChange={(e) => setNewBrand({ ...newBrand, description: e.target.value })} placeholder="Brand Description" className="w-full p-2 border rounded mb-2" />
          <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded">Add Brand</button>
        </form>

        {/* Add Category Form */}
        <form onSubmit={addCategory} className="bg-gray-100 p-4 rounded">
          <h3 className="text-lg font-semibold mb-2">Add Category</h3>
          <input value={newCategory.name} onChange={(e) => setNewCategory({ ...newCategory, name: e.target.value })} placeholder="Category Name" required className="w-full p-2 border rounded mb-2" />
          <textarea value={newCategory.description} onChange={(e) => setNewCategory({ ...newCategory, description: e.target.value })} placeholder="Category Description" className="w-full p-2 border rounded mb-2" />
          <button type="submit" className="bg-green-600 text-white px-4 py-2 rounded">Add Category</button>
        </form>
      </div>

      {/* ---------------------- Add New Product ---------------------- */}
      <form onSubmit={handleAddProduct} className="space-y-3 bg-gray-100 p-4 rounded mb-8">
        <h3 className="text-lg font-semibold mb-2">Add New Product</h3>
        <div className="grid grid-cols-2 gap-4">
          <input type="text" name="name" value={form.name || ''} onChange={handleInputChange} placeholder="Name" required className="p-2 border rounded" />
          <select name="brand" value={form.brand || ''} onChange={handleInputChange} required className="p-2 border rounded">
            <option value="">Select Brand</option>
            {brands.map((b) => (
              <option key={b.brand_name} value={b.brand_name}>{b.brand_name}</option>
            ))}
          </select>
          <select name="category" value={form.category || ''} onChange={handleInputChange} required className="p-2 border rounded">
            <option value="">Select Category</option>
            {categories.map((c) => (
              <option key={c.category_name} value={c.category_name}>{c.category_name}</option>
            ))}
          </select>
          <input type="number" name="price" value={form.price || ''} onChange={handleInputChange} placeholder="Price" required className="p-2 border rounded" />
          <input type="number" name="stock" value={form.stock || ''} onChange={handleInputChange} placeholder="Stock" required className="p-2 border rounded" />
        </div>
        <textarea name="description" value={form.description || ''} onChange={handleInputChange} placeholder="Description" required className="w-full p-2 border rounded" />
        <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded">Add Product</button>
      </form>

      {/* ---------------------- Product List Table ---------------------- */}
      <table className="w-full border-collapse text-sm">
        <thead>
          <tr className="bg-gray-200">
            <th className="border p-2">Name</th>
            <th className="border p-2">Price (NOK)</th>
            <th className="border p-2">Stock</th>
            <th className="border p-2">Actions</th>
          </tr>
        </thead>
        <tbody>
          {products.map((p) => (
            <tr key={p.product_id}>
              <td className="border p-2">{p.name}</td>
              <td className="border p-2">
                <input type="number" value={p.price} onChange={(e) => handleEditChange(p.product_id, 'price', e.target.value)} className="w-20 p-1 border rounded" />
              </td>
              <td className="border p-2">
                <input type="number" value={p.stock_quantity} onChange={(e) => handleEditChange(p.product_id, 'stock_quantity', e.target.value)} className="w-16 p-1 border rounded" />
              </td>
              <td className="border p-2">
                <button className="text-red-600 hover:underline" onClick={() => handleDelete(p.product_id)}>Delete</button>
                <button className="text-green-600 hover:underline ml-2" onClick={() => handleSave(p.product_id)}>Save</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default AdminDashboard;
