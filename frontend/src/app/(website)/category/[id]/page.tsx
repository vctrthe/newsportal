"use client";

import { ApiResponse, Pagination } from "@/model/ApiResponse";
import { useEffect, useState } from "react";
import Swal from "sweetalert2";
import Link from "next/link";
import Image from "next/image";
import { Button } from "@/components/ui/button";
import { ArrowLeft, ArrowRight } from "lucide-react";
import { Content } from "@/model/Content";
import axiosInstance from "../../../../../lib/axios";
import React from "react";

type Params = {
    id: number
}

interface ContentByCategoryPageProps {
    params: Promise<Params>;
}

export default function ContentByCategory({params}: ContentByCategoryPageProps) {
    const [content, setContent] = useState<Content | null>(null);
    const [pagination, setPagination] = useState<Pagination | null>(null);
    const [currentPage, setCurrentPage] = useState(1);
    const [categoryId, setCategoryId] = useState<number | null>(null);

    useEffect(() => {
        async function resolveParams() {
            try {
                const resolvedParams = await params;
                // console.log("Resolved params:", resolvedParams);
                setCategoryId(resolvedParams.id);
            } catch (error) {
                console.error("Error resolving params:", error)
            }
        }
        resolveParams();
    }, [params]);

    useEffect(() => {
        if (categoryId !== null) {
            fetchData(currentPage);
        }
    }, [categoryId, currentPage]);

    const fetchData = async (page:number = 1) => {
        if (categoryId === null) {
            console.error("Category ID not resolved");
            return;
        }

        try {
            const response = await axiosInstance.get<ApiResponse<Content>>(`/fe/contents/${categoryId}?limit=6&page=${page}`);
            // console.log("API response:", response.data);
            setContent(response.data.data || null);
            setPagination(response.data.pagination ?? null);
        } catch (error: any) {
            Swal.fire({
                icon: "error",
                title: "Oops!",
                text: error.message,
                toast: true,
                showConfirmButton: false,
                timer: 1500
            });
        }
    };

    useEffect(() => {
        console.log("Pagination Data:", pagination);
    }, [pagination]);


    const handlePrevClick = () => {
        if (pagination && currentPage > 1) {
            setCurrentPage(currentPage - 1);
        }
    }

    const handleNextClick = () => {
        if (pagination && currentPage < pagination.total_pages) {
            setCurrentPage(currentPage + 1);
        }
    }

    return (
        <div>
            <div className="container px-8 mx-auto xl:px-5 max-w-screen-lg py-5 lg:py-8 relative">
                <h1 className="text-center text-3xl font-semibold tracking-tight lg:text-4xl lg:leading-snug">
                    Konten
                </h1>
                <div className="text-center">
                    <p className="mt-2 text-lg">Lihat semua konten</p>
                </div>

                <div className="mt-10 grid gap-10 md:grid-cols-2 lg:gap-10 xl:grid-cols-3">
                    {content ? (
                        <div key={content.id} className="group cursor-pointer">
                            <div className="overflow-hidden rounded-md bg-gray-100 transition-all hover:scale-105">
                                <Link href={`/content-all/detail/${content.id}`} className="relative block aspect-ratio">
                                    {content.image != "" && (
                                        <Image src={content.image} alt={content.title} width={600} height={400} className="object-cover transition-all" />
                                    )}
                                    {content.image == "" && (
                                        <Image src="https://placehold.co/600x400" alt="data" className="object-cover transition-all" width={600} height={400} />
                                    )}
                                </Link>
                            </div>
                            <div>
                                <div className="flex gap-3">
                                    <Link href={`/category/${content.category_id}`}>
                                        <span className="inline-block text-sm font-medium tracking-wider uppercase mt-5 text-blue-600">
                                            {content.category_name}
                                        </span>
                                    </Link>
                                </div>
                                <h2 className="text-lg font-semibold leading-snug tracking-tight mt-2">
                                    <Link href={`/content-all/detail/${content.id}`}>
                                        <span className="bg-gradient-to-r from-green-200 to-green-100 bg-[length:0px_10px] bg-left-bottom bg-no-repeat transition-[background-size] duration-500 hover:bg-[length:100%_3px]">
                                            {content.title}
                                        </span>
                                    </Link>
                                </h2>
                                <div className="mt-3 flex items-center space-x-3 text-gray-500">
                                    <Link href={""}>
                                        <div className="flex items-center gap-3">
                                            <div className="relative h-5 w-5 flex-shrink-0">
                                                <Image src="https://i.pravatar.cc/32?u=john@example.com" alt="author" className="rounded-full object-cover" sizes="20px" width={20} height={20}/>
                                            </div>
                                                <span className="truncate text-sm">{content.author}</span>
                                        </div>
                                    </Link>
                                    <span className="text-xs text-gray-300">•</span>
                                    <time dateTime={"2024-11-30T08:30:45Z"} className="truncate text-sm">{content.created_at}</time>
                                </div>
                            </div>
                        </div>
                    ) : (
                    <p>Konten tidak ditemukan</p>
                )}
                </div>

                {pagination?.total_pages && pagination.total_pages > 1 ? (
                    <div className="mt-10 flex items-center justify-center">
                        <nav className="isolate inline-flex space-x-px rounded-md shadow-sm">
                            <Button disabled={currentPage === 1} className="relative inline-flex items-center gap-1 rounded-l-md border border-gray-300 bg-white px-3 py-2 pr-4 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20 disabled:pointer-events-none disabled:opacity-40" onClick={handlePrevClick}>
                                <ArrowLeft className="h3 w-3 stroke-1"/>
                                <span>Sebelumnya</span>
                            </Button>

                            <Button disabled={pagination.total_pages <= currentPage} className="relative inline-flex items-center gap-1 rounded-l-md border border-gray-300 bg-white px-3 py-2 pr-4 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20 disabled:pointer-events-none disabled:opacity-40" onClick={handleNextClick}>
                                <ArrowRight className="h3 w-3 stroke-1"/>
                                <span>Selanjutnya</span>
                            </Button>
                        </nav>
                    </div>
                ) : null}
            </div>
        </div>
    );
}