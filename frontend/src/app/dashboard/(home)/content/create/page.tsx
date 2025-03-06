"use client";

import { ApiResponse } from "@/model/ApiResponse";
import { Category } from "@/model/Category";
import { useEffect, useState } from "react";
import axiosInstance from "../../../../../../lib/axios";
import FormContentPage from "../components/form-content";

const CreateContentPage = () => {
    const [categories, setCategories] = useState<Category[]>([]);
    useEffect(() => {
        const fetchData = async() => {
            try {
                const response = await axiosInstance.get<ApiResponse<Category[]>>("/admin/categories");
                setCategories(response.data.data);
            } catch (error) {
                alert("Error fetching categories");
            }
        }

        fetchData();
    }, []);

    return (
        <div>
            <div className="flex flex-row items-center justify-between">
                <div className="my-5 text-2xl font-bold">Tambah Konten</div>
            </div>

            <FormContentPage type="ADD" categoryList={categories}/>
        </div>
    );
}

export default CreateContentPage;