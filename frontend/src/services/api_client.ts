
const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";


export const apiRequest = async <T>(endpoint: string, method: "GET" | "POST" | "PUT" | "DELETE", body?: any): Promise<T> => {
    const options: RequestInit = {
        method,
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
        },
        body: body ? JSON.stringify(body) : undefined,
    }

    const response = await fetch(`${API_URL}${endpoint}`, {
        ...options,
        credentials: "include", // this tells the browser to include cookies in the request
    })

    if (!response.ok) {
        const errJson = await response.json()
        if (errJson.error) {
            throw new Error(errJson.error)
        } else {
            throw new Error("Unknown error occurred")
        }
    }

    // handle diferent response types
    const contentType = response.headers.get("Content-Type");
    if (contentType != "application/json") {
        throw new Error(`Error: ${response.status} ${response.statusText} - ${contentType}`);
    }
    const data = await response.json();
    console.log("Response data: ", data);
    return data as T;
}

