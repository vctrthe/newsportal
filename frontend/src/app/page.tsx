"use client";
import { Content } from "@/model/Content";
import Image from "next/image";
import { useEffect, useState } from "react";
import axiosInstance from "../../lib/axios";
import { ApiResponse } from "@/model/ApiResponse";
import Swal from "sweetalert2";
import Link from "next/link";

export default function Home() {
  const [contents, setContents] = useState<Content[]>([]);
  const [sliceData, setSliceData] = useState<Content[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axiosInstance.get<ApiResponse<Content[]>>("/fe/contents?limit=8")
        const sliced = response.data.data.slice(0,2);
        setContents(response.data.data);
        setSliceData(sliced);
        setContents((prevContents) => prevContents.slice(2));
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
  }, []);
  return (
    <div>
      <div className="grid gap-10 md:grid-cols-2 lg:gap-10">
        {sliceData.map((content, index) => (
          <div key={index} className="group cursor-pointer">
            <div className="overflow-hidden rounded-md bg-gray-100 transition-all hover:scale-105">
              <Link className="relative block aspect-video" href={""}>
              {content.image != "" && (
                <Image src={content.image} alt={content.title} width={600} height={400} className="object-cover transition-all"/>
              )}
              </Link>
            </div>
            <div>
              <div>
                <div className="flex gap-3">
                  <Link href={`/category/${content.category_id}`}>
                    <span className="inline-block text-sx font-medium tracking-wider upercase mt-5 text-blue-600">{content.category_name}</span>
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
                  <time dateTime={"2024-11-30 08:30:45+07:00"} className="truncate text-sm">{content.created_at}</time>
                </div>
              </div>
            </div>
          </div>
        ))}
      </div>

      <div className="mt-10 grid gap:10 md:grid-cols-2 lg:gap-10 xl:grid-cols-3">
        {contents.map((content, index) => (
          <div key={index} className="group cursor-pointer">
            <div className="overflow-hidden rounded-md bg-gray-100 transition-all hover:scale-105">
              <Link className="relative block aspect-ratio" href={`/content-all/detail/${content.id}`}>
                {content.image != "" && (
                  <Image src={content.image} alt={content.title} width={600} height={400} className="object-cover transition-all" />
                )}
                {content.image == "" && (
                  <Image src="https://placehold.co/600x400" alt="data" className="object-cover transition-all" width={600} height={400}/>
                )}
              </Link>
            </div>
            <div>
              <div>
                <div className="flex gap-3">
                  <Link href={`/category/${content.category_id}`}>
                    <span className="inline-block text-sx font-medium tracking-wider upercase mt-5 text-blue-600">{content.category_name}</span>
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
                  <time dateTime={"2024-11-30 08:30:45+07:00"} className="truncate text-sm">{content.created_at}</time>
                </div>
              </div>
            </div>
          </div>
        ))}
      </div>

      <div className="mt-10 flex justify-center">
        <Link href={"/content-all"} className="relative inline-flex items-center gap-1 rounded-md border-gray-300 bg-white px-3 py-2 pl-4 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20">
          <span>
            See All Posts
          </span>
        </Link>
      </div>
    </div>
  );
}
