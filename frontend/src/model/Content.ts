export interface Content {
    id: number;
    title: string;
    excerpt: string;
    description: string;
    image: string;
    tags: string[];
    status: string;
    category_name: string;
    created_at: string;
    author: string;
    created_by_name: string;
    created_by_id: number;
    category_id: number;
}

export interface ContentRequest {
    title: string;
    excerpt: string;
    description: string;
    image: string;
    tags: string[];
    status: string;
    created_by_id: number;
    category_id: number;
}