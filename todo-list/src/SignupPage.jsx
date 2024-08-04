import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

function SignupPage() {
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

    const handleSignup = async () => {
        try {
            const response = await axios.post("http://localhost:8000/signup", {
                username: username,
                password: password,
            });
            navigate("/login");
            console.log(response.data.message);
            console.log(response);
        } catch (error) {
            console.log(error.response.data);
            console.log(error);
        }
    };

    function handleBack() {
        navigate("/");
    }

    return (
        <div className="signup-page">
            <h1>Sign-up page</h1>
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
                <button className="normal-button" onClick={handleSignup}>
                    Sign up
                </button>
                <button className="normal-button" onClick={handleBack}>
                    back
                </button>
            </div>
        </div>
    );
}
export default SignupPage;
