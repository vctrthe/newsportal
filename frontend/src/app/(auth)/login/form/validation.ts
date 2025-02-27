import { z } from "zod";

export const formSchema = z.object({
    email: z.string({required_error: 'Email harus diisi'}).email({message: 'Email tidak valid'}),
    password: z.string({required_error: 'Password harus diisi minimal 6 karakter'}).min(6)
})