import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import LoginPage from "./LoginPage";
import TodoListPage from "./TodoListPage";
import Navbar from "./Navbar";

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<LoginPage />} />
                <Route path="/todolist" element={<TodoListPage />} />
            </Routes>
        </Router>
    );
}

export default App;
