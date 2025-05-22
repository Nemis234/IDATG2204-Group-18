import React, { useEffect, useState } from "react";
import ProductCard from "../components/ProductCard";

function Home() {
  const [products, setProducts] = useState([]);
  const [filtered, setFiltered] = useState([]);
  const [categories, setCategories] = useState([]);
  const [brands, setBrands] = useState([]);
  const [filters, setFilters] = useState({ category: "", brand: "" });

  // Fetch products, categories, and brands on load
  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const res = await fetch("http://localhost:8080/products?page=1");
        const data = await res.json();
        if (Array.isArray(data)) {
          setProducts(data);
          setFiltered(data);
        } else {
          console.error("Invalid products response", data);
          setProducts([]);
          setFiltered([]);
        }
      } catch (err) {
        console.error("Failed to fetch products", err);
      }
    };

    const fetchCategories = async () => {
      try {
        const res = await fetch("http://localhost:8080/categories");
        const data = await res.json();
        if (Array.isArray(data)) {
          setCategories(data);
        } else {
          console.error("Invalid categories response", data);
          setCategories([]);
        }
      } catch (err) {
        console.error("Failed to fetch categories", err);
      }
    };

    const fetchBrands = async () => {
      try {
        const res = await fetch("http://localhost:8080/brands");
        const data = await res.json();
        if (Array.isArray(data)) {
          setBrands(data);
        } else {
          console.error("Invalid brands response", data);
          setBrands([]);
        }
      } catch (err) {
        console.error("Failed to fetch brands", err);
      }
    };

    fetchProducts();
    fetchCategories();
    fetchBrands();
  }, []);

  // Handle filter change
  const handleFilterChange = (type, value) => {
    const updatedFilters = { ...filters, [type]: value };
    setFilters(updatedFilters);

    const filteredList = products.filter((product) => {
      const matchesCategory =
        !updatedFilters.category || product.category_name === updatedFilters.category;
      const matchesBrand =
        !updatedFilters.brand || product.brand_name === updatedFilters.brand;
      return matchesCategory && matchesBrand;
    });

    setFiltered(filteredList);
  };

  return (
    <div className="p-6 max-w-6xl mx-auto">
      <h2 className="text-2xl font-bold mb-4">Available Products</h2>

      {/* Filters */}
      <div className="flex gap-4 mb-6">
        <select
          value={filters.category}
          onChange={(e) => handleFilterChange("category", e.target.value)}
          className="p-2 border rounded"
        >
          <option value="">All Categories</option>
          {Array.isArray(categories) &&
            categories.map((c) => (
              <option key={c.category_name} value={c.category_name}>
                {c.category_name}
              </option>
            ))}
        </select>

        <select
          value={filters.brand}
          onChange={(e) => handleFilterChange("brand", e.target.value)}
          className="p-2 border rounded"
        >
          <option value="">All Brands</option>
          {Array.isArray(brands) &&
            brands.map((b) => (
              <option key={b.brand_name} value={b.brand_name}>
                {b.brand_name}
              </option>
            ))}
        </select>
      </div>

      {/* Product Grid */}
      {Array.isArray(filtered) && filtered.length === 0 ? (
        <p className="text-gray-600">No products match the selected filters.</p>
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
          {Array.isArray(filtered) &&
            filtered.map((product) => (
              <ProductCard key={product.product_id} product={product} />
            ))}
        </div>
      )}
    </div>
  );
}

export default Home;
