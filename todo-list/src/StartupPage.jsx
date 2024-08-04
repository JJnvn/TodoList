import { useNavigate } from "react-router-dom";

function StartupPage() {
    const navigate = useNavigate();

    function handleLogin() {
        navigate("/login");
    }

    function handleSignup() {
        navigate("/signup");
    }

    return (
        <>
            <div className="startup-page">
                <h1>Startup Page</h1>
                <div className="button-container">
                    <button className="normal-button" onClick={handleLogin}>
                        Log in
                    </button>
                    <button className="normal-button" onClick={handleSignup}>
                        Sign up
                    </button>
                </div>
            </div>
        </>
    );
}
export default StartupPage;
