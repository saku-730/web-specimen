// src/components/create/OccurrenceForm.tsx

"use client"; // ユーザーの入力を扱うので、クライアントコンポーネントにするのだ！

import { useState, useEffect } from "react";

// まずは、APIのレスポンスと送信データの型を定義するのだ（後で別のファイルに分けると綺麗になる）
type DropdownOptions = {
  users: { user_id: number; user_name: string }[];
  projects: { project_id: number; project_name: string }[];
  languages: { language_id: number; language_common: string }[];
  observation_methods: { observation_method_id: number; observation_method_name: string }[];
  specimen_methods: { specimen_methods_id: number; specimen_methods_common: string }[];
  institutions: { institution_id: number; institution_code: string }[];
};

// ... 送信するデータの型定義もここに書く ...

const OccurrenceForm = () => {
  // フォーム全体の入力値を、この一つの state で管理するのだ
  const [formData, setFormData] = useState<any>({}); // APIから受け取るのでanyで初期化
  // ドロップダウンの選択肢を記憶する state
  const [dropdowns, setDropdowns] = useState<DropdownOptions | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // ページが読み込まれた時に、APIから初期データを取得するのだ
  useEffect(() => {
    const fetchData = async () => {
      try {
        // ここで /create のGET APIを叩く。トークンを忘れずに！
	const apiUrl = `${process.env.NEXT_PUBLIC_API_URL}/create`;
        const res = await fetch(apiUrl, {
          headers: {
            "Authorization": `Bearer ${localStorage.getItem("token")}` // 仮でlocalStorageから取得
          }
        });
        if (!res.ok) throw new Error("Failed to fetch initial data");
        
        const data = await res.json();
        setDropdowns(data.dropdown_list);
        // default_valueがあれば、それをフォームの初期値にセットする
        if (data.default_value) {
          setFormData(data.default_value);
        }
      } catch (err: any) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, []); // []が空なので、この処理は最初の1回だけ実行されるのだ

  // 入力値が変更された時に、formDataを更新する関数
  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    
    // nameに "classification.species" のようなドットが含まれている場合に対応する
    const keys = name.split('.');
    if (keys.length > 1) {
      setFormData((prev: any) => ({
        ...prev,
        [keys[0]]: {
          ...prev[keys[0]],
          [keys[1]]: value
        }
      }));
    } else {
      setFormData((prev: any) => ({
        ...prev,
        [name]: value
      }));
    }
  };
  
  // フォームが送信された時の処理
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault(); // ページの再読み込みを防ぐ
    
    try {
      // POSTリクエストをバックエンドに送る
      const res = await fetch("http://localhost:8080/api/v0_0_2/create", {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
        body: JSON.stringify(formData)
      });
      
      if (!res.ok) throw new Error("Failed to submit data");
      
      const result = await res.json();
      alert("Success! Created Occurrence ID: " + result.OccurrenceID);
      // ここで成功ページにリダイレクトする処理などを書く
      // window.location.href = `/occurrences/${result.OccurrenceID}`;

    } catch (err: any) {
      alert("Error: " + err.message);
    }
  };

  if (loading) return <div>Loading form...</div>;
  if (error) return <div className="text-red-500">Error: {error}</div>;

  return (
    <form onSubmit={handleSubmit} className="space-y-8">
      {/* --- ここから各セクションの部品を並べていくのだ --- */}

      {/* 基本情報セクション */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6 border-b pb-8">
        <div>
          <label htmlFor="user_id" className="block text-sm font-medium text-gray-700">User</label>
          <select id="user_id" name="user_id" value={formData.user_id || ''} onChange={handleChange} className="mt-1 block w-full rounded-md shadow-sm text-black border border-[#808080] caret-black">
            {dropdowns?.users.map(user => (
              <option key={user.user_id} value={user.user_id}>{user.user_name}</option>
            ))}
          </select>
        </div>
        <div>
          <label htmlFor="project_id" className="block text-sm font-medium text-gray-700">Project</label>
          <select id="project_id" name="project_id" value={formData.project_id || ''} onChange={handleChange} className="mt-1 block w-full rounded-md shadow-sm text-black border border-[#808080] caret-black">
            {dropdowns?.projects.map(p => (
              <option key={p.project_id} value={p.project_id}>{p.project_name}</option>
            ))}
          </select>
        </div>
        <div>
          <label htmlFor="created_at" className="block text-sm font-medium text-gray-700">Date</label>
          <input type="datetime-local" id="created_at" name="created_at" value={formData.created_at ? new Date(formData.created_at).toISOString().slice(0, 16) : ''} onChange={handleChange} className="mt-1 block w-full rounded-md shadow-sm text-black border border-[#808080] caret-black" />
        </div>
      </div>

      {/* 分類セクション */}
      <div className="border-b pb-8">
        <h2 className="text-lg font-semibold mb-4 text-gray-700">Classification</h2>
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div>
            <label htmlFor="classification.kingdom" className="block text-sm font-medium text-gray-700">Kingdom</label>
            <input type="text" id="classification.kingdom" name="classification.kingdom" value={formData.classification?.kingdom || ''} onChange={handleChange} className="mt-1 block w-full rounded-md shadow-sm text-black border border-[#808080] caret-black" />
          </div>
          <div>
            <label htmlFor="classification.phylum" className="block text-sm font-medium text-gray-700">Phylum</label>
            <input type="text" id="classification.phylum" name="classification.phylum" value={formData.classification?.phylum || ''} onChange={handleChange} className="mt-1 block w-full rounded-md shadow-sm text-black border border-[#808080] caret-black" />
          </div>
	  <div>
            <label htmlFor="classification.class" className="block text-sm font-medium text-gray-700">class</label>
            <input type="text" id="classification.class" name="classification.class" value={formData.classification?.class || ''} onChange={handleChange} className="mt-1 block w-full rounded-md shadow-sm text-black border border-[#808080] caret-black" />
          </div>



        </div>
      </div>

      {/* ... Observation, Specimen, Identification のセクションも同様に作る ... */}
      

      <div className="flex justify-end">
        <button type="submit" className="bg-blue-600 text-white px-6 py-2 rounded-md hover:bg-blue-700">
          Create Occurrence
        </button>
      </div>
    </form>
  );
};

export default OccurrenceForm;
