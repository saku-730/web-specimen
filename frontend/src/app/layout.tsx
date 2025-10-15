// src/app/layout.tsx

import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Sidebar from "@/components/layout/Sidebar";
import Header from "@/components/layout/Header";

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
        {/* 'relative' を追加して、子要素の 'absolute' の基準点にするのだ */}
        <div className="relative min-h-screen">
          <Sidebar /> {/* ⬅️ このサイドバーは後で「浮いた」存在になる */}

          {/* メインコンテンツのエリア。
            'pl-20' で、閉じたサイドバーの幅(w-20)と同じだけの左パディングを常に追加するのだ。
          */}
          <div className="flex-1 flex flex-col pl-20">
            <Header />
            <main className="flex-1 p-6 overflow-y-auto">
              {children}
            </main>
          </div>
        </div>
      </body>
    </html>
  );
}
