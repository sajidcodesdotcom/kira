import {  useNavigate } from "react-router-dom";
import { useAuthStore } from "../hooks/auth_store";

interface protectedRouteProps {
    children: React.ReactNode;
}

export default function ProtectedRoute({ children }: protectedRouteProps) {
    const navigate = useNavigate();
    const { isLoggedIn, isLoading } = useAuthStore();

    if (isLoading) {
        return (
            <div className="flex justify-center items-center h-screen">
            <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
            <span className="ml-3 text-lg font-medium text-gray-700">Loading...</span>
            </div>
        );
    }

    if (!isLoggedIn) {
        navigate("/login");
    }


    return <>{children}</>;
}