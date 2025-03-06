import { Button } from "@/components/ui/button";
import { useFormStatus } from "react-dom";
import React from "react";

export default function SubmitButtonForm() {
    const {pending} = useFormStatus();

    return (
        <Button disabled={pending} className="w-full">Submit</Button>
    );
}