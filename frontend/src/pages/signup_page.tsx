import React, { useState } from "react";
import Input from "../components/common/input";
import { useNavigate } from "react-router-dom";
import { useAuthStore } from "../store/auth_store";
import { AuthService } from "../services/auth_service";
import { Link } from "react-router-dom";

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
    const [showPassword, setShowPassword] = useState(false);
    const [showConfirmPassword, setShowConfirmPassword] = useState(false);

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
        <>
            <div className="m-auto w-full max-w-md h-screen grid place-items-center">
                <div className="bg-white rounded-xl p-8 w-full shadow-2xl overflow-hidden">
                    <div className="p-8">

                        <h1 className="text-3xl font-semibold text-gray-800 text-center mb-6">Create Account</h1>
                        
                        {error && (
                            <div className="mb-6 p-3 bg-red-50 border-l-4 border-red-500 text-red-700 rounded">
                                <p>{error}</p>
                            </div>
                        )}
                        
                        <form onSubmit={handleSubmit}>
                            <div className="space-y-5">
                                <div>
                                    <label htmlFor="fullName" className="block text-sm font-medium text-gray-700 mb-1">Full Name</label>
                                    <Input 
                                        handleChange={handleChange} 
                                        value={formData.fullName} 
                                        type="text" 
                                        id="fullName" 
                                        placeholder="John Doe" 
                                        className="w-full" 
                                    />
                                </div>
                                
                                <div>
                                    <label htmlFor="username" className="block text-sm font-medium text-gray-700 mb-1">Username</label>
                                    <Input 
                                        handleChange={handleChange} 
                                        value={formData.username} 
                                        type="text" 
                                        id="username" 
                                        placeholder="johndoe" 
                                        className="w-full" 
                                    />
                                </div>

                                <div>
                                    <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-1">Email</label>
                                    <Input 
                                        handleChange={handleChange} 
                                        value={formData.email} 
                                        type="email" 
                                        id="email" 
                                        placeholder="you@example.com" 
                                        className="w-full" 
                                    />
                                </div>
                                
                                <div className="relative">
                                    <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-1">Password</label>
                                    <Input 
                                        handleChange={handleChange} 
                                        value={formData.password} 
                                        type={showPassword ? "text" : "password"}
                                        id="password" 
                                        placeholder="Your password" 
                                        className="w-full" 
                                    />
                                    <button
                                        type="button"
                                        className="absolute right-0 top-8 flex items-center pr-3 text-gray-500"
                                        aria-label="Toggle password visibility"
                                        onClick={() => setShowPassword(!showPassword)}
                                    >
                                        {showPassword ? (
                                            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                                                <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
                                                <path fillRule="evenodd" d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z" clipRule="evenodd" />
                                            </svg>
                                        ) : (   
                                            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                                                <path fillRule="evenodd" d="M3.707 2.293a1 1 0 00-1.414 1.414l14 14a1 1 0 001.414-1.414l-1.473-1.473A10.014 10.014 0 0019.542 10C18.268 5.943 14.478 3 10 3a9.958 9.958 0 00-4.512 1.074l-1.78-1.781zm4.261 4.26l1.514 1.515a2.003 2.003 0 012.45 2.45l1.514 1.514a4 4 0 00-5.478-5.478z" clipRule="evenodd" />
                                                <path d="M12.454 16.697L9.75 13.992a4 4 0 01-3.742-3.741L2.335 6.578A9.98 9.98 0 00.458 10c1.274 4.057 5.065 7 9.542 7 .847 0 1.669-.105 2.454-.303z" />
                                            </svg>
                                        )}
                                    </button>
                                </div>
                                
                                <div className="relative">
                                    <label htmlFor="confirmPassword" className="block text-sm font-medium text-gray-700 mb-1">Confirm Password</label>
                                    <Input 
                                        handleChange={handleChange} 
                                        value={formData.confirmPassword} 
                                        type={showConfirmPassword ? "text" : "password"}
                                        id="confirmPassword" 
                                        placeholder="Confirm your password" 
                                        className="w-full" 
                                    />
                                    <button
                                        type="button"
                                        className="absolute right-0 top-8 flex items-center pr-3 text-gray-500"
                                        aria-label="Toggle password visibility"
                                        onClick={() => setShowConfirmPassword(!showConfirmPassword)}
                                    >
                                        {showConfirmPassword ? (
                                            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                                                <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
                                                <path fillRule="evenodd" d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z" clipRule="evenodd" />
                                            </svg>
                                        ) : (   
                                            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                                                <path fillRule="evenodd" d="M3.707 2.293a1 1 0 00-1.414 1.414l14 14a1 1 0 001.414-1.414l-1.473-1.473A10.014 10.014 0 0019.542 10C18.268 5.943 14.478 3 10 3a9.958 9.958 0 00-4.512 1.074l-1.78-1.781zm4.261 4.26l1.514 1.515a2.003 2.003 0 012.45 2.45l1.514 1.514a4 4 0 00-5.478-5.478z" clipRule="evenodd" />
                                                <path d="M12.454 16.697L9.75 13.992a4 4 0 01-3.742-3.741L2.335 6.578A9.98 9.98 0 00.458 10c1.274 4.057 5.065 7 9.542 7 .847 0 1.669-.105 2.454-.303z" />
                                            </svg>
                                        )}
                                    </button>
                                </div>
                                
                                <button 
                                    disabled={loading} 
                                    type="submit" 
                                    className="w-full flex justify-center py-3 px-4 border border-transparent rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors disabled:opacity-75"
                                >
                                    {loading ? (
                                        <div className="flex items-center">
                                            <svg className="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                                                <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                                            </svg>
                                            <span>Creating account...</span>
                                        </div>
                                    ) : (
                                        <span>Sign Up</span>
                                    )}
                                </button>
                            </div>
                        </form>
                        
                        <div className="mt-6 text-center text-sm">
                            <p className="text-gray-600">
                                Already have an account? <Link to="/login" className="font-medium text-blue-600 hover:text-blue-500">Sign in</Link>
                            </p>
                        </div>
                    </div>

                    </div>
                </div>
        </>
    );
}