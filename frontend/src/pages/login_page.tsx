import React, { useState } from "react";
import Input from "../components/common/input";
import { useNavigate } from "react-router-dom";
import { useAuthStore } from "../store/auth_store";
import { AuthService } from "../services/auth_service";



export default function LoginPage() {
    const {setUser} = useAuthStore();
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        email: "",
        password: "",
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
        const {email, password} = formData;
        setError(null)
        setLoading(true);
        try {
            const data = await AuthService.login(email, password);
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
        <h1 className="text-3xl font-bold mb-4">Log In</h1>
        <form onSubmit={handleSubmit} className="bg-white p-6 rounded shadow-md w-96">
            {error && <p className="text-red-500 mb-4">{error}</p>}
            <div className="mb-4">
                <label htmlFor="email" className="sr-only">Email</label>
            <Input handleChange={handleChange} value={formData.email}  type={"email"} id={"email"} placeholder={"Email"} className={""} />
            </div>
            <div className="mb-4">
            <label htmlFor="password" className="sr-only">Password</label>
            <Input handleChange={handleChange} value={formData.password}  type={"password"} id={"password"} placeholder={"Password"} className={""} />
            </div>
            <button disabled={loading} type="submit" className="w-full bg-blue-600 text-white py-2 rounded p-1 hover:bg-blue-700 flex items-center justify-center">
            {loading ? (
                <>
                <span className="animate-ping h-2 w-2 mr-3 rounded-full bg-white opacity-85"></span>
                <span>Loggin In...</span>
                </>
            ) : (
                <>
                <span>Login</span>
                </>
            )}
            </button>
        </form>
        </div>
    );
    }