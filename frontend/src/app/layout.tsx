// src/app/layout.tsx
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Sidebar from "@/components/layout/Sidebar"; // サイドバーをインポート
import Header from "@/components/layout/Header";   // ヘッダーをインポート

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Web-Specimen",
  description: "Specimen Management Dashboard",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body className={`${inter.className} bg-gray-100`}>
        <div className="flex h-screen">
          <Sidebar /> {/* ⬅️ 左側のサイドバーを配置 */}
          <div className="flex-1 flex flex-col">
            <Header /> {/* ⬅️ 上部のヘッダーを配置 */}
            <main className="flex-1 p-6 overflow-y-auto">
              {children} {/* ⬅️ ここに page.tsx の中身が入るのだ */}
            </main>
          </div>
        </div>
      </body>
    </html>
  );
}
