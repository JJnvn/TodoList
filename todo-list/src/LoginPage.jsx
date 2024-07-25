import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

function LoginPage() {
    // create error box
    const navigate = useNavigate();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleUsernameChange = (event) => {
        setUsername(event.target.value);
    };

    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    };

    const handleLogin = async () => {
        try {
            const response = await axios.post("http://localhost:8000/login", {
                username: username,
                password: password,
            });
            console.log(response.data.message);
            console.log(response);
            navigate("/todolist");
        } catch (error) {
            console.log(error.response.data);
            console.log(error);
        }
    };

    const handleSignup = async () => {
        try {
            const response = await axios.post("http://localhost:8000/signup", {
                username: username,
                password: password,
            });
            console.log(response.data.message);
            console.log(response);
        } catch (error) {
            console.log(error.response.data);
            console.log(error);
        }
    };

    return (
        <div className="login-page">
            <h1>Login page</h1>
            <h2>username</h2>
            <input
                className="input-box"
                type="text"
                placeholder="enter the username"
                onChange={handleUsernameChange}
            />
            <h2>password</h2>
            <input
                className="input-box"
                type="text"
                placeholder="enter the password"
                onChange={handlePasswordChange}
            />
            <div className="button-container">
                <button className="normal-button" onClick={handleLogin}>
                    Login
                </button>
                <button className="normal-button" onClick={handleSignup}>
                    Sign up
                </button>
            </div>
        </div>
    );
}
export default LoginPage;
