"use client";

import { AlertCircle } from "lucide-react";
import axiosInstance, { setupInterceptor } from "../../../../../lib/axios";
import { User } from "@/model/User";
import { useEffect, useState } from "react";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { ApiResponse } from "@/model/ApiResponse";
import FormUserPage from "./components/form-user";

export default function UserPage() {
    setupInterceptor();
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axiosInstance.get<ApiResponse<User>>("/admin/users/profile");
                setUser(response.data.data);
                setLoading(false);
            } catch (error: any) {
                setError(error.message || "Error fetching data");
                setLoading(false);
            }
        }

        fetchData();
    }, [])

    if (loading) {
        return <div>Loading...</div>;
    }

    if (error) return (
        <Alert variant="destructive">
            <AlertCircle className="h-4 w-4"/>
            <AlertTitle> Error</AlertTitle>
            <AlertDescription>
                {error}
            </AlertDescription>
        </Alert>
    )

    return (
        <>
            <div className="flex flex-row items-center justify-between">
                <div className="my-5 text-2xl font-bold">Pengguna</div>
            </div>
            <FormUserPage defaultValues={user}/>
        </>
    );
}