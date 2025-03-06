"use client"

import { FC, useEffect, useState } from "react"
import axiosInstance, { setupInterceptor } from "../../../../../../../lib/axios"
import { Category } from "@/model/Category"
import { ApiResponse } from "@/model/ApiResponse"
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert"
import { AlertCircle } from "lucide-react"
import FormCategoryPage from "../../components/form-category"

type Params = {
    id: number
}

interface EditCategoryPageProps {
    params: Params
}

const EditCategoryPage: FC<EditCategoryPageProps> = ({ params }) => {
    setupInterceptor();
    
    const [category, setCategory] = useState<Category | null>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                setLoading(false);
                const response = await axiosInstance.get<ApiResponse<Category>>(`/admin/categories/${params.id}`);
                setCategory(response.data.data);
            } catch (error: any) {
                setError(error.message || "Error fetching data");
                setLoading(false);
            }
        }

        fetchData();
    }, [params.id])

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
        <div>
            <div className="flex flex-row items-center justify-between">
                <div className="my-5 text-2xl font-bold">Ubah Kategori</div>
            </div>

            <FormCategoryPage type="EDIT" defaultValues={category}/>
        </div>
    );
}

export default EditCategoryPage;