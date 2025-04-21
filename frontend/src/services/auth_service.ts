import { AuthResponse } from "../types/api"
import { apiRequest } from "./api_client"

export const AuthService = {
    login: (email: string, password: string) => {
        return apiRequest<AuthResponse>("/api/auth/login", "POST", { email, password })
    },

    register: (fullName: string, username: string, email: string, password: string) => {
        return apiRequest<AuthResponse>("/api/auth/register", "POST", { full_name: fullName, username, email, password })
    },

    logout: () => apiRequest("/api/auth/logout", "POST"),

    getCurrentUser: () => {
        return apiRequest<AuthResponse>("/api/auth/me", "GET")
    },
}