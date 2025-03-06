"use client";

import React, { useEffect, useState } from "react";
import axiosInstance from "../../../../../../lib/axios";
import { ApiResponse } from "@/model/ApiResponse";
import { Content } from "@/model/Content";
import Swal from "sweetalert2";
import Link from "next/link";
import Image from "next/image";

type Params = {
    id: number
}

interface ContentDetailPageProps {
    params: Promise<Params>
}

export default function ContentDetail({params}: ContentDetailPageProps) {
    const resolvedParams = React.use(params);
    const [content, setContent] = useState<Content | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axiosInstance.get<ApiResponse<Content>>(`fe/contents/${resolvedParams.id}`);
                setContent(response.data.data);
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
        }

        fetchData();
    }, [resolvedParams.id]);

    return (
        <div>
            <div className="container px-8 mx-auto xl:px-5 max-w-screen py-5 lg:py-8 !pt-0">
                <div className="mx-auto max-w-screen-md">
                    <div className="flex justify-center">
                        <div className="flex gap-3">
                            <Link href={`/category/${content?.category_id}`}>
                                <span className="inline-block text-sm font-medium tracking-wider uppercase mt-5 text-blue-600">
                                    {content?.category_name}
                                </span>
                            </Link>
                        </div>
                    </div>
                    <h2 className="text-brand-primary mb-3 mt-2 text-center text-3xl font-semibold tracking-tight lg:text-4xl lg:leading-snug">
                        {content?.title}
                    </h2>
                </div>
                <div className="mt-3 flex justify-center space-x-3 text-gray-500">
                    <div className="flex items-center gap-3">
                        <div>
                            <p className="text-gray-800">
                                {content?.author}
                            </p>
                            <div className="flex items-center space-x-2 text-sm">
                                <time dateTime={"2024-11-30T15:30:45Z"} className="truncate text-sm">{content?.created_at}</time>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div className="relative z-0 mx-auto aspect-video max-w-screen-lg overflow-hidden lg:rounded-lg">
                {/* {content?.image && (
                    <Image
                        src={content.image}
                        alt={content?.title || "Default Title"}
                        className="object-cover w-screen"
                    />
                )} */}
                {/* <img src={content?.image} alt={`${content?.title}`} className="object-cover w-screen"/> */}
                <Image src={content?.image || "https://placehold.co/1280x720"} alt={`${content?.title || "Default"}`} className="object-cover w-screen" width={1280} height={720}/>
            </div>
            <div className="container px-8 mx-auto xl:px-5 max-w-screen py-5 lg:py-8">
                <article className="mx-auto max-w-screen-md">
                    <div className="prose mx-auto my-3">
                        {content?.description}
                    </div>
                    <div className="mb-7 mt-7 flex justify-center">
                        <Link href={`/content-all`} className="bg-brand-secondary/20 rounded-full px-5 py-2 text-sm text-blue-600">
                            Lihat semua konten
                        </Link>
                    </div>
                </article>
            </div>
        </div>
    )
}