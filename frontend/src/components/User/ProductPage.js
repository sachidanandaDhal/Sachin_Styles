import { useState, useEffect } from "react";

const ProductPage = () => {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/products")
      .then((res) => res.json())
      .then((data) => setProducts(data))
      .catch((err) => console.error("Error fetching products:", err));
  }, []);

  return (
    <div className="max-w-4xl mx-auto mt-10">
      <h2 className="text-3xl font-bold mb-4">Available Products</h2>
      <div className="grid grid-cols-2 gap-4">
        {products.map((product) => (
          <div key={product.id} className="p-4 border rounded-lg shadow">
            {product.images.length > 0 && (
              <img src={product.images[0]} alt={product.name} className="w-full h-40 object-cover mb-2" />
            )}
            <h2 className="text-xl font-semibold">{product.name}</h2>
            <p className="text-gray-600">{product.description}</p>
            <p className="text-lg font-bold text-blue-600">${product.price}</p>
            <p className="text-sm text-gray-500">Uploaded by Admin ID: {product.createdBy}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ProductPage;
