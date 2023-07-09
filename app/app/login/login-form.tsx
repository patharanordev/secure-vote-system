"use client";

import { signIn } from "next-auth/react";
import { useSearchParams, useRouter } from "next/navigation";
import { ChangeEvent, useState } from "react";

export const LoginForm = () => {
    const router = useRouter();
    const [loading, setLoading] = useState(false);
    const [formValues, setFormValues] = useState({
        username: "",
        password: "",
    });
    
    const [error, setError] = useState("");

    const searchParams = useSearchParams();
    const callbackUrl = searchParams?.get("callbackUrl") ?? "/dashboard";

    const onSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            setLoading(true);
            setFormValues({ username: "", password: "" });

            const res = await signIn("credentials", {
                redirect: false,
                username: formValues.username,
                password: formValues.password,
                callbackUrl,
            });

            setLoading(false);

            console.log(res);
            if (!res?.error) {
                router.push(callbackUrl);
            } else {
                setError("invalid username or password");
            }
        } catch (error: any) {
            setLoading(false);
            setError(error);
        }
    };

    const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setFormValues({ ...formValues, [name]: value });
    };

    return (
        <form onSubmit={onSubmit}>
            <label htmlFor="username">Username</label>
            <input
                required
                type="text"
                name="username"
                value={formValues.username}
                onChange={handleChange}
                style={{ padding: "1rem" }}
            />
            <label htmlFor="password">Password</label>
            <input
                required
                type="password"
                name="password"
                value={formValues.password}
                onChange={handleChange}
                style={{ padding: "1rem" }}
            />
            <button
                style={{
                    backgroundColor: `${loading ? "#ccc" : "#3446eb"}`,
                    color: "#fff",
                    padding: "1rem",
                    cursor: "pointer",
                }}
                disabled={loading}
            >
                {loading ? "loading..." : "Sign In"}
            </button>
        </form>
    );
};