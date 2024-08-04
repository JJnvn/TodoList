import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import LoginPage from "./LoginPage";
import TodoListPage from "./TodoListPage";
import SignupPage from "./SignupPage";
import StartupPage from "./StartupPage";

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<StartupPage />} />
                <Route path="/login" element={<LoginPage />} />
                <Route path="/signup" element={<SignupPage />} />
                <Route path="/todolist" element={<TodoListPage />} />
            </Routes>
        </Router>
    );
}

export default App;
