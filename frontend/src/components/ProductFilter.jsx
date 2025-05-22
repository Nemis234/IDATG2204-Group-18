// File: src/components/ProductFilter.jsx
// Description: Dropdown filter component for selecting product category and brand

import React from 'react';

/**
 * @param {Object} props
 * @param {Array} props.categories - List of category objects { category_name }
 * @param {Array} props.brands - List of brand objects { brand_name }
 * @param {Function} props.onFilter - Callback when a filter value changes: (type, value) => void
 */
function ProductFilter({ categories, brands, onFilter }) {
  return (
    <div className="mb-6 flex flex-wrap gap-4">
      {/* Category Filter Dropdown */}
      <select 
        onChange={e => onFilter('category', e.target.value)} 
        className="border px-3 py-2 rounded"
      >
        <option value="">All Categories</option>
        {categories.map(c => (
          <option key={c.category_name} value={c.category_name}>
            {c.category_name}
          </option>
        ))}
      </select>

      {/* Brand Filter Dropdown */}
      <select 
        onChange={e => onFilter('brand', e.target.value)} 
        className="border px-3 py-2 rounded"
      >
        <option value="">All Brands</option>
        {brands.map(b => (
          <option key={b.brand_name} value={b.brand_name}>
            {b.brand_name}
          </option>
        ))}
      </select>
    </div>
  );
}

export default ProductFilter;
