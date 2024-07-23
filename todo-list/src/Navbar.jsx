import React from "react";
import { useNavigate } from "react-router-dom";
import "./navbar.css";

function Navbar() {
    // add user profile

    const navigate = useNavigate();

    function handleLogout() {
        navigate("/");
    }

    return (
        <nav className="navbar">
            <h1 className="navbar-title">User Profile</h1>
            <button className="logout-button" onClick={handleLogout}>
                Logout
            </button>
        </nav>
    );
}

export default Navbar;
