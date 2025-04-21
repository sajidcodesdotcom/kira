import { create } from "zustand"
import { AuthResponse, User } from "../types/globals"
import { apiRequest } from "../services/api_client";

interface AuthStore {
    user: User | null;
    setUser: (user: User | null) => void;
    isLoggedIn: boolean;
    isLoading: boolean;
    checkAuthStatus: () => Promise<void>;
    logout: () => Promise<void>;
}

export const useAuthStore = create<AuthStore>((set) => ({
    user: null,
    isLoggedIn: false,
    isLoading: true,
    setUser: (user) => set({ user, isLoading: false, isLoggedIn: true }),
    checkAuthStatus: async () => {
        try {
            const data = await apiRequest<AuthResponse>("/api/auth/me", "GET");
            set({ user: data.user, isLoggedIn: true, isLoading: false });
        } catch (error) {
            console.error("Error checking auth status:", error);
            set({ user: null, isLoggedIn: false, isLoading: false });
        }
    },
    logout: async () => {
        try {
            await apiRequest("/api/auth/logout", "POST");
            set({ user: null, isLoggedIn: false, isLoading: false });
        } catch (error) {
            console.error("Error logging out:", error);
        }
    }
}))

const { checkAuthStatus } = useAuthStore.getState();

checkAuthStatus()