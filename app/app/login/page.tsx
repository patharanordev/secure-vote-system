import { LoginForm } from "./login-form";

export default function LoginPage() {
  return (
    <div
      style={{
        display: "flex",
        height: "70vh",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <div>
        <h1>Login</h1>
        <LoginForm />
      </div>
    </div>
  );
}
