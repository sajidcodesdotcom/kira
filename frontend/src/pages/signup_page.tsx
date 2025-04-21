import React, { useState } from "react";
import Input from "../components/common/input";
import { useNavigate } from "react-router-dom";
import { useAuthStore } from "../store/auth_store";
import { AuthService } from "../services/auth_service";



export default function SignUpPage() {
    const {setUser} = useAuthStore();
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        fullName: "",
        username: "",
        email: "",
        password: "",
        confirmPassword: "",
    })

    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
        const {id, value} = event.target;
        setFormData(
            {
                ...formData,
                [id]: value
            }
                    )
    }

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        const {fullName, username, email, password, confirmPassword} = formData;
        setError(null)
        if (password !== confirmPassword) {
            setError ("Passwords do not match");
            return;
        }
        setLoading(true);
        try {
            const data = await AuthService.register(fullName, username, email, password);
            setUser(data.user);
            navigate("/dashboard");
        } catch (error: any) {
            setError(error.message);
            console.log("Error while signing up", error)
            
        } finally {
            setLoading(false);
        }
    }
    return (
        <div className="flex flex-col items-center justify-center overflow-hidden h-screen bg-gray-100">
        <h1 className="text-3xl font-bold mb-4">Sign Up</h1>
        <form onSubmit={handleSubmit} className="bg-white p-6 rounded shadow-md w-96">
            {error && <p className="text-red-500 mb-4">{error}</p>}
            <div className="mb-4">
                <label htmlFor="fullName" className="sr-only">Full Name</label>
            <Input  handleChange={handleChange} value={formData.fullName} type={"text"} id={"fullName"} placeholder={"Full name"} className={""} />
            </div>
            <div className="mb-4">
                <label htmlFor="username" className="sr-only">Username</label>
            <Input handleChange={handleChange} value={formData.username}  type={"text"} id={"username"} placeholder={"Username"} className={""} />
            </div>
            <div className="mb-4">
                <label htmlFor="email" className="sr-only">Email</label>
            <Input handleChange={handleChange} value={formData.email}  type={"email"} id={"email"} placeholder={"Email"} className={""} />
            </div>
            <div className="mb-4">
            <label htmlFor="password" className="sr-only">Password</label>
            <Input handleChange={handleChange} value={formData.password}  type={"password"} id={"password"} placeholder={"Password"} className={""} />
            </div>
            <div className="mb-4">
                <label htmlFor="confirmPassword" className="sr-only">Confirm Password</label>
            <Input handleChange={handleChange} value={formData.confirmPassword}  type={"password"} id={"confirmPassword"} placeholder={"Confirm Password"} className={""} />
            </div>
            <button disabled={loading} type="submit" className="w-full bg-blue-600 text-white py-2 rounded p-1 hover:bg-blue-700 flex items-center justify-center">
            {loading ? (
                <>
                <span className="animate-ping h-2 w-2 mr-3 rounded-full bg-white opacity-85"></span>
                <span>Signing Up...</span>
                </>
            ) : (
                <>
                <span>Sign Up</span>
                </>
            )}
            </button>
        </form>
        </div>
    );
    }