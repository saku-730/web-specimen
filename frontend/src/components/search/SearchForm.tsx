// src/components/search/SearchForm.tsx

"use client";

import { useState } from "react";

// 親コンポーネントから渡される props の型
type Props = {
  dropdowns: any; // 型は後でもっと厳密にするのだ
  onSearch: (params: Record<string, string>) => void;
  isLoading: boolean;
};

const SearchForm = ({ dropdowns, onSearch, isLoading }: Props) => {
  // フォームの入力値をまとめて管理する state
  const [params, setParams] = useState<Record<string, string>>({});

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    setParams(prev => ({ ...prev, [e.target.name]: e.target.value }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // 空のパラメータを除外して、親コンポーネントに通知する
    const filteredParams: Record<string, string> = {};
    for (const key in params) {
      if (params[key]) {
        filteredParams[key] = params[key];
      }
    }

    const searchKeys = Object.keys(filteredParams).filter(k => k !== 'page' && k !== 'per_page');
    if (searchKeys.length === 0) {
      alert("Please enter at least one search condition.");
      return; // ここで処理を中断！
    }


    onSearch(filteredParams);
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-6">
      <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6">
        {/* --- ここにたくさんの入力フォームを並べるのだ --- */}
        
        {/* ドロップダウンの例 */}
        <div>
          <label htmlFor="user_id" className="block text-sm font-medium text-gray-700">User</label>
          <select id="user_id" name="user_id" onChange={handleChange} value={params.user_id || ''} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm text-black">
            <option value="">Any</option>
            {dropdowns?.users?.map((user: any) => (
              <option key={user.user_id} value={user.user_id}>{user.user_name}</option>
            ))}
          </select>
        </div>
        
        {/* テキスト入力の例 */}
        <div>
          <label htmlFor="species" className="block text-sm font-medium text-gray-700">Species</label>
          <input type="text" id="species" name="species" onChange={handleChange} value={params.species || ''} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm text-black" />
        </div>

        {/* 日付範囲の例 */}
        <div>
          <label htmlFor="created_start" className="block text-sm font-medium text-gray-700">Created After</label>
          <input type="date" id="created_start" name="created_start" onChange={handleChange} value={params.created_start || ''} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm text-black" />
        </div>
        <div>
          <label htmlFor="created_end" className="block text-sm font-medium text-gray-700">Created Before</label>
          <input type="date" id="created_end" name="created_end" onChange={handleChange} value={params.created_end || ''} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm text-black" />
        </div>

        {/* ... 他のたくさんの入力項目も、ここに追加していくのだ ... */}

      </div>
      <div className="flex justify-end">
        <button type="submit" disabled={isLoading} className="bg-blue-600 text-white px-6 py-2 rounded-md hover:bg-blue-700 disabled:bg-blue-300">
          {isLoading ? 'Searching...' : 'Search'}
        </button>
      </div>
    </form>
  );
};

export default SearchForm;
