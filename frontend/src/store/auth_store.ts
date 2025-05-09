import { create } from "zustand"
import { AuthResponse, User } from "../types/models"
import { apiRequest } from "../services/api_client";
import { AuthService } from "../services/auth_service";

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
            const data = await AuthService.getCurrentUser();
            set({ user: data.user, isLoggedIn: true, isLoading: false });
        } catch (error) {
            console.error("Error checking auth status:", error);
            set({ user: null, isLoggedIn: false, isLoading: false });
        }
    },
    logout: async () => {
        try {
            await AuthService.logout();
            set({ user: null, isLoggedIn: false, isLoading: false });
        } catch (error) {
            console.error("Error logging out:", error);
        }
    }
}))

const { checkAuthStatus } = useAuthStore.getState();

checkAuthStatus()