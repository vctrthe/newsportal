# newsportal-frontend

## Table of Contents
1. [Project Structure](#frontend-project-structure)
2. [How to Use](#how-to-use)
3. [Packages needed](#packages-needed)
4. [Learn More about Next.JS](#learn-more-about-nextjs)
5. [Information about Vercel](#deploy-on-vercel)

## Frontend Project Structure
```
newsportal/
└── frontend/
    ├── lib/
    │   └── axios.ts                                    -> Interceptor for getting Access Token from Browser Cookies
    │
    ├── public/
    │   ├── img/                                        -> Store Custom Website Logo
    │   │   └── .gitkeep                                -> Ensuring the folder still exists
    │   │
    │   ├── file.svg
    │   ├── globe.svg
    │   ├── next.svg
    │   ├── vercel.svg
    │   └── window.svg
    │
    ├── src/
    │   ├── app/
    │   │   ├── (auth)/                                 -> Authentication domain
    │   │   │   ├── login/
    │   │   │   │   ├── form/
    │   │   │   │   │   ├── index.tsx                   -> Login form
    │   │   │   │   │   └── validation.ts               -> Login Validation
    │   │   │   │   │
    │   │   │   │   └── page.tsx                        -> Login page
    │   │   │   │
    │   │   │   └── layout.tsx                          -> Layout for login page
    │   │   │
    │   │   ├── (website)/                              -> Main website for user non-admin
    │   │   │   ├── category/                           -> Category domain
    │   │   │   │   └── [id]/                           -> Select by Category ID
    │   │   │   │       └── page.tsx                    -> Page for displaying content by category ID
    │   │   │   │
    │   │   │   └── content-all/                        -> All Contents domain
    │   │   │       ├── detail/
    │   │   │       │   ├── [id]/                       -> Detail of a Content by ID
    │   │   │       │   │   └── page.tsx                -> Page for displaying Content Detail
    │   │   │       │   │
    │   │   │       │   └── page.tsx                    -> Page for all content details
    │   │   │       │
    │   │   │       └── page.tsx                        -> Page for listing all contents
    │   │   │
    │   │   ├── dashboard/                              -> Admin Dashboard domain
    │   │   │   ├── (home)/                             -> Dashboard home domain
    │   │   │   │   ├── category/                       -> Admin Category domain
    │   │   │   │   │   ├── components/
    │   │   │   │   │   │   ├── columns-table.tsx       -> Table settings for displaying categories
    │   │   │   │   │   │   ├── delete-category.tsx     -> Delete Category
    │   │   │   │   │   │   └── form-category.tsx       -> Form for categorie
    │   │   │   │   │   │
    │   │   │   │   │   ├── create/
    │   │   │   │   │   │   └── page.tsx                -> Page for creating a category
    │   │   │   │   │   │
    │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   └── [id]/
    │   │   │   │   │   │       └── page.tsx            -> Page for editing a category (by ID)
    │   │   │   │   │   │
    │   │   │   │   │   ├── lib/
    │   │   │   │   │   │   ├── action.ts               -> Action button setting (for edit and delete in category dashboard)
    │   │   │   │   │   │   └── validation.ts           -> Validation rules
    │   │   │   │   │   │
    │   │   │   │   │   └── page.tsx                    -> Category Dashboard Page
    │   │   │   │   │
    │   │   │   │   ├── components/
    │   │   │   │   │   ├── button-logout.tsx           -> Button setting for logout admin user
    │   │   │   │   │   └── submit-button.tsx           -> Button setting for Submitting form (global for dashboard)
    │   │   │   │   │
    │   │   │   │   ├── content/
    │   │   │   │   │   ├── components/
    │   │   │   │   │   │   ├── columns-table.tsx       -> Table setting for displaying content
    │   │   │   │   │   │   ├── delete-content.tsx      -> Delete Content
    │   │   │   │   │   │   └── form-content.tsx        -> Form for content
    │   │   │   │   │   │
    │   │   │   │   │   ├── create/
    │   │   │   │   │   │   └── page.tsx                -> Page for creating a content
    │   │   │   │   │   │
    │   │   │   │   │   ├── edit/
    │   │   │   │   │   │   └── [id]/
    │   │   │   │   │   │       └── page.tsx            -> Page for editing a content (by ID)
    │   │   │   │   │   │
    │   │   │   │   │   ├── lib/
    │   │   │   │   │   │   ├── action.ts               -> Action button setting (for edit and delete in content dashboard)
    │   │   │   │   │   │   └── validation.ts           -> Validation rules
    │   │   │   │   │   │
    │   │   │   │   │   └── page.tsx                    -> Content Dashboard Page
    │   │   │   │   │
    │   │   │   │   └── user/
    │   │   │   │       ├── components/
    │   │   │   │       │   └── form-user.tsx           -> Form for user
    │   │   │   │       │
    │   │   │   │       ├── lib/
    │   │   │   │       │   └── action.ts               -> Action settings for update password
    │   │   │   │       │
    │   │   │   │       └── page.tsx                    -> User Dashboard Page
    │   │   │   │
    │   │   │   ├── layout.tsx                          -> Layout for Login Page
    │   │   │   └── page.tsx                            -> Login Page
    │   │   │
    │   │   ├── favicon.ico
    │   │   ├── globals.css
    │   │   ├── layout.tsx                              -> Layout for Index page
    │   │   └── page.tsx                                -> Index Page
    │   │
    │   ├── components/
    │   │   ├── ui/                                     -> Gotten from Shadcn
    │   │   │   ├── alert.tsx
    │   │   │   ├── button.tsx
    │   │   │   ├── data-table.tsx
    │   │   │   ├── input.tsx
    │   │   │   ├── label.tsx
    │   │   │   ├── select.tsx
    │   │   │   ├── table.tsx
    │   │   │   └── textarea.tsx
    │   │   │
    │   │   ├── footer.tsx                              -> Index's footer setting
    │   │   └── navbar.tsx                              -> Index's navbar setting
    │   │
    │   ├── lib/
    │   │   └── utils.ts
    │   │
    │   ├── model/                                      -> Mapper for responses gotten from newsportal-be
    │   │   ├── ApiResponse.ts                          -> Default Response Interface
    │   │   ├── Category.ts                             -> Category Interface
    │   │   ├── Content.ts                              -> Content Interface
    │   │   └── User.ts                                 -> User Interface
    │   │
    │   └── middleware.ts                               -> Middleware for frontend (login checker)
    │
    ├── .env.example
    ├── components.json
    ├── eslint.config.mjs
    ├── next.config.ts.example
    ├── package-lock.json
    ├── package.json
    ├── postcss.config.mjs
    ├── README.md                                       -> Frontend-Specific documentation
    ├── tailwind.config.ts
    └── tsconfig.json 
```

## How to Use
1. Rename `next.config.ts.example` to `next.config.ts` and `.env.example` to `.env`. Fill the appropriate credentials for each files.
2. Install dependencies using (cd to the folder first):
```
npm install
# or
yarn install
# or
pnpm install
# or
bun install
```
3. Run the development server:

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
# or
bun dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## Packages needed
### Shadcn UI
```bash
npx shadcn@latest init
npx shadcn@latest add *
```
*) _add components, for example `button`_

## Learn More About Next.js

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/app/building-your-application/deploying) for more details.
