import "./styles/Login.css";
// import Box, { SettingsPowerTwoTone } from "@mui/icons-material";
import { TextField } from "@mui/material";
import { useState } from "react";
import "dotenv";
import "js-sha256";
import { sha256 } from "js-sha256";

function Login() {
  // handle error messages
  const [errorMessage, setErrorMessage] = useState("");

  // handle transitions from login view to sign up view
  const [isLogin, setIsLogin] = useState(true);

  const API_URL = import.meta.env.VITE_GRPC_GATEWAY_URL;

  const handleSubmit = () => {
    if (isLogin) {
      Login();
    } else {
      Signup();
    }
  };

  const Login = () => {
    // validate input
    const username = (
      document.getElementById("username-field") as HTMLInputElement
    )?.value;
    const password = (
      document.getElementById("password-field") as HTMLInputElement
    )?.value;
    if (!username) {
      setErrorMessage("Please enter a username.");
    } else if (!password) {
      setErrorMessage("Please enter a password.");
    }

    // attempt to login
    try {
      const hashToken = sha256(password);
      fetch(`${API_URL}/v1/GetUser`, {
        method: "POST",
        body: JSON.stringify({
          username: username,
          hash_token: hashToken,
        }),
        headers: {
          "Content-type": "application/json; charset=UTF-8",
        },
      }).then((res) => {
        if (!res.ok) {
          setErrorMessage(`Error: ${res.text()}`);
        } else {
          res
            .json()
            .then((json) => console.log(`Welcome, ${json["username"]}`));
        }
      });
    } catch {
      setErrorMessage("Something went wrong, please try again later.");
    }
  };

  const Signup = () => {
    // validate input
    const firstName = (
      document.getElementById("first-name-field") as HTMLInputElement
    )?.value;
    const lastName = (
      document.getElementById("last-name-field") as HTMLInputElement
    )?.value;
    const username = (
      document.getElementById("username-field") as HTMLInputElement
    )?.value;
    const password = (
      document.getElementById("password-field") as HTMLInputElement
    )?.value;
    const confirmPassword = (
      document.getElementById("confirm-password-field") as HTMLInputElement
    )?.value;
    if (!firstName) {
      setErrorMessage("Please enter a first name.");
    } else if (!lastName) {
      setErrorMessage("Please enter a last name.");
    } else if (!username) {
      setErrorMessage("Please enter a username.");
    } else if (!password) {
      setErrorMessage("Please enter a password.");
    } else if (password !== confirmPassword) {
      setErrorMessage("Passwords do not match.");
    }

    // attempt to log in
    try {
      const hashToken = sha256(password);
      console.log(firstName);
      console.log(lastName);
      fetch(`${API_URL}/v1/CreateUser`, {
        method: "POST",
        body: JSON.stringify({
          first_name: firstName,
          last_name: lastName,
          username: username,
          hash_token: hashToken,
        }),
        headers: {
          "Content-type": "application/json; charset=UTF-8",
        },
      }).then((res) => {
        if (!res.ok) {
          setErrorMessage(`Error: ${res.text()}`);
        } else {
          res.json().then((json) => {
            console.log(json);
            console.log(`Welcome, ${json["user"]["username"]}`);
          });
        }
      });
    } catch {
      setErrorMessage("Something went wrong, please try again later.");
    }
  };

  const handleSwitchView = () => {
    setIsLogin((prev) => !prev);
  };

  return (
    <div id="login-container">
      <div className="title-subtitle">
        <h1>Chat App</h1>
        <h5>Powered by gRPC</h5>
        <h6 id="error-msg">{errorMessage}</h6>
      </div>
      <div id="forms">
        {!isLogin && (
          <>
            <TextField
              className="form-field"
              id="first-name-field"
              label="First Name"
              variant="filled"
              sx={{ marginTop: "10px" }}
            />
            <TextField
              className="form-field"
              id="last-name-field"
              label="Last Name"
              variant="filled"
              sx={{ marginTop: "10px" }}
            />
          </>
        )}
        <TextField
          className="form-field"
          id="username-field"
          label="Username"
          variant="filled"
          sx={{ marginTop: "10px", borderRadius: "15px" }}
        />
        <TextField
          className="form-field"
          id="password-field"
          label="Password"
          variant="filled"
          type="password"
          sx={{ marginTop: "10px" }}
        />
        {!isLogin && (
          <TextField
            className="form-field"
            id="confirm-password-field"
            label="Confirm password"
            variant="filled"
            type="password"
            sx={{ marginTop: "10px" }}
          />
        )}
      </div>
      <div id="submit-button" onClick={handleSubmit}>
        {isLogin ? "Login" : "Signup"}
      </div>
      <div id="switch-view-button" onClick={handleSwitchView}>
        {isLogin
          ? "Don't have an account? Sign up here."
          : "Already have an account? Log in here."}
      </div>
    </div>
  );
}

export default Login;
