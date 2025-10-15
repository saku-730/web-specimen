// src/app/login/page.tsx

import LoginForm from "@/components/auth/LoginForm";

export default function LoginPage() {
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-50">
      <div className="w-full max-w-md">
        <h2 className="text-center text-3xl font-bold tracking-tight text-gray-900 mb-8">
          Login to Web-Specimen
        </h2>
        <LoginForm />
      </div>
    </div>
  );
}
