import { z } from "zod";

export const categoryFormSchema = z.object({
    title: z.string({required_error: 'Judul harus diisi'})
})