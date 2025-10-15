// src/components/auth/LoginForm.tsx

"use client"; // ユーザーの入力を扱うので、クライアントコンポーネントにするのだ！

import { useState } from "react";
import { useRouter } from "next/navigation"; // ページ遷移に使うのだ
import { Mail, Lock } from "lucide-react";

const LoginForm = () => {
  // フォームの入力値を記憶するための箱なのだ
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  
  const router = useRouter(); // ページ遷移のための道具

  // フォームが送信された時の処理
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault(); // ページの再読み込みを防ぐ
    setIsLoading(true);
    setError(null);

    try {
      // 1. バックエンドにPOSTリクエストを送る
      const apiUrl = `${process.env.NEXT_PUBLIC_API_URL}/login`;
      const res = await fetch(apiUrl, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });

      // もしログインに失敗したら、エラーを投げる
      if (!res.ok) {
        throw new Error("Invalid email or password. Please try again.");
      }

      // 2. レスポンスからトークンを取り出す
      const data = await res.json();
      const token = data.token; // バックエンドが { "token": "..." } という形で返すと仮定

      // 3. トークンをブラウザに保存するのだ！これが「鍵」になる
      localStorage.setItem("token", token);

      // 4. ダッシュボードページに強制的に移動させる（リダイレクト）
      router.push("/");

    } catch (err: any) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="bg-white p-8 rounded-lg shadow-md space-y-6">
      {/* エラーメッセージを表示する場所 */}
      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative">
          <p>{error}</p>
        </div>
      )}

      {/* メールアドレス入力 */}
      <div>
        <label htmlFor="email" className="block text-sm font-medium text-gray-700">
          Email address
        </label>
        <div className="mt-1 relative">
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <Mail className="h-5 w-5 text-gray-400" />
          </div>
          <input
            id="email"
            name="email"
            type="email"
            autoComplete="email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="block w-full rounded-md border-gray-300 pl-10 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
            placeholder="you@example.com"
          />
        </div>
      </div>

      {/* パスワード入力 */}
      <div>
        <label htmlFor="password" className="block text-sm font-medium text-gray-700">
          Password
        </label>
        <div className="mt-1 relative">
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <Lock className="h-5 w-5 text-gray-400" />
          </div>
          <input
            id="password"
            name="password"
            type="password"
            autoComplete="current-password"
            required
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="block w-full rounded-md border-gray-300 pl-10 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
            placeholder="••••••••"
          />
        </div>
      </div>

      {/* ログインボタン */}
      <div>
        <button
          type="submit"
          disabled={isLoading}
          className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:bg-indigo-300"
        >
          {isLoading ? "Logging in..." : "Login"}
        </button>
      </div>
    </form>
  );
};

export default LoginForm;
