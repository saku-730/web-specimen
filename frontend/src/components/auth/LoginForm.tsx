// src/components/auth/LoginForm.tsx

"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Mail, Lock } from "lucide-react";

const LoginForm = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError(null);

    try {
      // 1. Next.js自身が持つAPIルートにリクエストを送るのだ！
      const res = await fetch('/api/login', { // ⬅️ 宛先を変更！
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });

      if (!res.ok) {
        throw new Error("Invalid email or password. Please try again.");
      }
      
      // 2. ログイン成功！クッキーはサーバーがセットしてくれたので、ダッシュボードに移動するだけ！
      router.push("/");
      router.refresh(); // ⬅️ サーバーの状態を更新するために追加

    } catch (err: any) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    // ... フォームの見た目(JSX)は変更なし ...
    <form onSubmit={handleSubmit} className="bg-white p-8 rounded-lg shadow-md space-y-6">
      {/* エラーメッセージ */}
      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative">
          <p>{error}</p>
        </div>
      )}
      {/* メールアドレス入力 */}
      <div>
        <label htmlFor="email" className="block text-sm font-medium text-gray-700">Email address</label>
        <div className="mt-1 relative">
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <Mail className="h-5 w-5 text-gray-400" />
          </div>
          <input id="email" name="email" type="email" autoComplete="email" required value={email} onChange={(e) => setEmail(e.target.value)} className="block w-full rounded-md pl-10 shadow-sm border border-[#808080] text-black" placeholder="you@example.com" />
        </div>
      </div>
      {/* パスワード入力 */}
      <div>
        <label htmlFor="password" className="block text-sm font-medium text-gray-700">Password</label>
        <div className="mt-1 relative">
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <Lock className="h-5 w-5 text-gray-400" />
          </div>
          <input id="password" name="password" type="password" autoComplete="current-password" required value={password} onChange={(e) => setPassword(e.target.value)} className="block w-full rounded-md pl-10 shadow-sm border border-[#808080] text-black" placeholder="••••••••" />
        </div>
      </div>
      {/* ログインボタン */}
      <div>
        <button type="submit" disabled={isLoading} className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 disabled:bg-indigo-300">
          {isLoading ? "Logging in..." : "Login"}
        </button>
      </div>
    </form>
  );
};

export default LoginForm;
