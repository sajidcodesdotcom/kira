import { User } from "./models";

export type AuthResponse = {
    token: string;
    user: User;
}