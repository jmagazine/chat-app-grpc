import "./styles/Login.css";
import Box from "@mui/icons-material";
import { TextField } from "@mui/material";
import * as proto from "../../../backend/src/chat_server/chat_grpc_web_pb.js"; // Adjust the path as needed
// import { CreateUserParams } from "../../../backend/src/chat_server/chat_pb";
// Import message class
function Home() {
  console.log(proto);
  return (
    <div className="login-container">
      <div className="title-subtitle">
        <h1>Chat App</h1>
        <h5>Powered by gRPC</h5>
      </div>
      <div className="forms">
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
          sx={{ marginTop: "10px" }}
        />
      </div>
      <div className="login-button">Login</div>
      <div className="signup-button">Already Have an Account? Sign Up.</div>
    </div>
  );
}

export default Home;
