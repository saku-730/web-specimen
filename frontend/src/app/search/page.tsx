// src/app/search/page.tsx

"use client";

import { useState, useEffect } from "react";
import SearchForm from "@/components/search/SearchForm";
import SearchResults from "@/components/search/SearchResults";

// APIレスポンスの型を定義するのだ
type DropdownOptions = {
  users: { user_id: number; user_name: string }[];
  projects: { project_id: number; project_name: string }[];
  // ... 他のドロップダウンの型もここに追加
};

type SearchResponse = {
  occurrence_results: any[]; // 本当はもっと厳密な型が良いのだ
  metadata: {
    total_results: number;
    current_page: number;
    per_page: number;
    total_pages: number;
  };
};


export default function SearchPage() {
  const [dropdowns, setDropdowns] = useState<DropdownOptions | null>(null);
  const [searchResults, setSearchResults] = useState<SearchResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  
  // 最初にページが開かれた時、ドロップダウンの選択肢を取得する
  useEffect(() => {
    const fetchDropdowns = async () => {
      try {
        const apiUrl = `${process.env.NEXT_PUBLIC_API_URL}/search`;
        const token = localStorage.getItem("token");
        if (!token) throw new Error("Authentication token not found.");

        const res = await fetch(apiUrl, {
          headers: { "Authorization": `Bearer ${token}` }
        });
        if (!res.ok) throw new Error("Failed to fetch dropdown options");

        const data = await res.json();
        setDropdowns(data); // バックエンドのレスポンスに合わせてキーを修正
      } catch (err: any) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };
    fetchDropdowns();
  }, []);

  // 検索が実行された時に呼ばれる関数
  const handleSearch = async (params: Record<string, string>) => {
    setLoading(true);
    setError(null);
    try {
      // クエリパラメータをURLに組み立てるのだ
      const query = new URLSearchParams(params).toString();
      const apiUrl = `${process.env.NEXT_PUBLIC_API_URL}/search?${query}`;
      const token = localStorage.getItem("token");

      const res = await fetch(apiUrl, {
        headers: { "Authorization": `Bearer ${token}` }
      });
      if (!res.ok) throw new Error("Failed to fetch search results");

      const data = await res.json();
      setSearchResults(data);
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };


  if (error) return <div className="text-red-500">Error: {error}</div>;

  return (
    <div className="p-4 md:p-8">
      <h1 className="text-2xl font-semibold mb-6 text-gray-800">Search Occurrences</h1>
      
      {/* 検索フォーム */}
      <div className="bg-white p-8 rounded-lg shadow-md mb-8">
        {loading && !dropdowns ? (
          <p>Loading form...</p>
        ) : (
          <SearchForm dropdowns={dropdowns} onSearch={handleSearch} isLoading={loading} />
        )}
      </div>

      {/* 検索結果 */}
      <div className="bg-white p-8 rounded-lg shadow-md">
        {loading ? (
          <p>Searching...</p>
        ) : (
          <SearchResults data={searchResults} />
        )}
      </div>
    </div>
  );
}
