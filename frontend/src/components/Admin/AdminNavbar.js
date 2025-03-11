import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";

const AdminNavbar = () => {
  const [adminName, setAdminName] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    // Fetch admin name securely from backend
    const fetchAdminDetails = async () => {
      const response = await fetch("http://localhost:8080/admin-info", { credentials: "include" });
      const data = await response.json();
      if (response.ok) {
        setAdminName(data.name);
      } else {
        navigate("/login");
      }
    };
    fetchAdminDetails();
  }, [navigate]);

  const handleLogout = async () => {
    await fetch("http://localhost:8080/logout", { method: "POST", credentials: "include" });
    navigate("/login");
  };

  return (
    <nav className="bg-blue-600 p-4 text-white flex justify-between items-center">
      <h1 className="text-xl font-bold">Admin Panel</h1>
      <div className="flex space-x-6">
        <Link to="/admin-dashboard" className="hover:text-gray-300">Home</Link>
        <Link to="/products" className="hover:text-gray-300">Products</Link>
        <Link to="/all-users" className="hover:text-gray-300">All Users</Link>
      </div>
      <div className="flex items-center space-x-4">
        <span className="font-semibold">{adminName}</span>
        <button onClick={handleLogout} className="bg-red-500 px-4 py-2 rounded">Logout</button>
      </div>
    </nav>
  );
};

export default AdminNavbar;
