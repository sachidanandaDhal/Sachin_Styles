import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import Login from "./pages/Login";
import Register from "./pages/Register";
import ProductPage from "./components/User/ProductPage";
import AdminDashboard from "./components/Admin/AdminDashboard";
import SuperAdminDashboard from "./components/SuperAdmin/SuperAdminDashboard";

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/products" element={<ProductPage />} />
        <Route path="/admin-dashboard" element={<AdminDashboard />} />
        <Route path="/super-dashboard" element={<SuperAdminDashboard />} />
      </Routes>
    </Router>
  );
}

export default App;
