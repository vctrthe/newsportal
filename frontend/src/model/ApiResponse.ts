export interface Meta {
    status: boolean;
    message: string;
}

export interface ApiResponse<T> {
    data: T;
    meta: Meta;
}