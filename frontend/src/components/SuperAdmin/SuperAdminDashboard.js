import React from "react";

const SuperAdminDashboard = () => {
  return (
    <div className="max-w-2xl mx-auto mt-10 p-6 bg-white shadow-lg rounded-lg">
      <h2 className="text-3xl font-bold">Super Admin Dashboard</h2>
      <p className="mt-4">Only super admins can see this page.</p>
    </div>
  );
};

export default SuperAdminDashboard;
