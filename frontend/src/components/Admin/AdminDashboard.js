import { useState } from "react";
// import AdminNavbar from "./AdminNavbar";

const AdminDashboard = () => {
  const [formData, setFormData] = useState({
    name: "",
    description: "",
    price: "",
    images: [null, null, null, null], // Array for 4 image slots
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleImageChange = (e, index) => {
    const file = e.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = () => {
        const updatedImages = [...formData.images];
        updatedImages[index] = reader.result; // Convert image to Base64
        setFormData({ ...formData, images: updatedImages });
      };
    }
  };

  const handleAddProduct = async (e) => {
    e.preventDefault();

    const token = localStorage.getItem("token");
    if (!token) {
      alert("You must be logged in as an admin to add a product.");
      return;
    }

    const productData = {
      name: formData.name,
      description: formData.description,
      price: formData.price,
      images: formData.images.filter((img) => img !== null), // Send only selected images
    };

    const response = await fetch("http://localhost:8080/products", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(productData),
    });

    const data = await response.json();
    if (response.ok) {
      alert("Product added successfully!");
    } else {
      alert(`Error: ${data.error}`);
    }
  };

  return (
    <>
      {/* <AdminNavbar /> */}
    <div className="max-w-md mx-auto mt-10">
      <h2 className="text-3xl font-bold mb-4">Admin Dashboard</h2>
      <form onSubmit={handleAddProduct} className="space-y-4">
        <input type="text" name="name" placeholder="Name" onChange={handleChange} className="w-full p-2 border rounded" required />
        <input type="text" name="description" placeholder="Description" onChange={handleChange} className="w-full p-2 border rounded" required />
        <input type="number" name="price" placeholder="Price" onChange={handleChange} className="w-full p-2 border rounded" required />

        {/* Image Upload Boxes */}
        <div className="grid grid-cols-2 gap-4">
          {formData.images.map((image, index) => (
            <div key={index} className="relative w-32 h-32 border-2 border-dashed border-gray-400 flex items-center justify-center">
              {image ? (
                <img src={image} alt="Preview" className="w-full h-full object-cover" />
              ) : (
                <label className="absolute inset-0 flex items-center justify-center cursor-pointer">
                  <span className="text-gray-400">+</span>
                  <input type="file" accept="image/*" className="hidden" onChange={(e) => handleImageChange(e, index)} />
                </label>
              )}
            </div>
          ))}
        </div>

        <button type="submit" className="w-full bg-blue-600 text-white p-2 rounded">Add Product</button>
      </form>
    </div>
    </>
  );
};

export default AdminDashboard;
