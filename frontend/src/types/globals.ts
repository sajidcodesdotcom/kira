export interface User {
    id: string;
    full_name: string;
    email: string;
    password: string;
    username: string;
    role: string;
    avatar_url: string;
    created_at: string;
    updated_at: string;
}

export interface Project {
    id: string;
    name: string;
    description: string;
    owner_id: string;
    status: string;
    created_at: string;
    updated_at: string;
}

export interface Task {
    id: string;
    title: string;
    description: string;
    project_id: string;
    assignee_id: string;
    status: string;
    priority: string;
    due_date: string;
    created_at: string;
    updated_at: string;
}