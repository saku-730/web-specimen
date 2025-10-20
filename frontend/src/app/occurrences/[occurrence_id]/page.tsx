// src/app/occurrences/[occurrence_id]/page.tsx

import { cookies } from 'next/headers'; // サーバーサイドでクッキーを読むために必要
import { notFound } from 'next/navigation'; // データが見つからない時に404ページを表示する
import type { OccurrenceDetailResponse } from '@/types/occurrence';

// --- データを取得する非同期関数 ---
async function getOccurrenceData(id: string): Promise<OccurrenceDetailResponse | null> {
  const token = cookies().get('token')?.value; // クッキーからトークンを取得
  const apiUrl = `${process.env.NEXT_PUBLIC_API_BASE_URL}/occurrences/${id}`;

  if (!token) {
    // トークンがない場合 (ミドルウェアで弾かれるはずだけど念のため)
    console.error("サーバーコンポーネントで認証トークンが見つかりません。");
    // 本来はここでログインページにリダイレクトさせるなどの処理が良い
    return null;
  }

  try {
    // Goサーバーにデータをリクエスト！
    const res = await fetch(apiUrl, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
      cache: 'no-store', // 動的なデータなのでキャッシュしない
    });

    if (!res.ok) {
      if (res.status === 404) {
        notFound(); // データが見つからなければ404ページを表示
      }
      // それ以外のエラー
      throw new Error(`Failed to fetch data. Status: ${res.status}`);
    }

    return await res.json();
  } catch (error) {
    console.error("Fetch Error:", error);
    // エラーが発生した場合もnullを返す（またはエラーページを表示）
    return null;
  }
}

// --- ラベルと値を表示するヘルパー部品 ---
function DetailItem({ label, value }: { label: string; value?: string | number | null }) {
  if (value === null || value === undefined || value === '') {
    return null; // 値が空なら何も表示しない
  }
  return (
    <div className="py-2">
      <dt className="text-sm font-medium text-gray-500">{label}</dt>
      <dd className="mt-1 text-sm text-gray-900">{String(value)}</dd>
    </div>
  );
}

// --- メインのページコンポーネント ---
// params.occurrence_id でURLのIDが取れるのだ！
export default async function OccurrenceDetailPage({ params }: { params: { occurrence_id: string } }) {
  
  // サーバーサイドでデータを取得するのだ！
  const data = await getOccurrenceData(params.occurrence_id);

  // データが取得できなかった場合の表示
  if (!data) {
    return (
      <div className="p-8 text-center text-red-500">
        Failed to load occurrence data. Please try again later.
      </div>
    );
  }

  // 日付を見やすい形にするヘルパー関数 (必要に応じて調整)
  const formatDate = (isoString: string) => {
    if (!isoString) return 'N/A';
    try {
      return new Date(isoString).toLocaleString('ja-JP');
    } catch {
      return isoString; // パース失敗時は元の文字列
    }
  };

  return (
    <div className="bg-gray-100 min-h-screen p-8">
      <div className="max-w-4xl mx-auto bg-white shadow-lg rounded-lg overflow-hidden">
        <div className="px-6 py-4 bg-gray-50 border-b">
          <h1 className="text-xl font-semibold text-gray-800">Occurrence Details (ID: {params.occurrence_id})</h1>
        </div>

        <div className="px-6 py-4">
          {/* --- 基本情報 --- */}
          <h2 className="text-lg font-medium text-gray-700 mb-2">Basic Information</h2>
          <dl className="grid grid-cols-1 md:grid-cols-3 gap-x-4 gap-y-1 divide-y divide-gray-200">
            <DetailItem label="User" value={data.user_name} />
            <DetailItem label="Project" value={data.project_name} />
            <DetailItem label="Individual ID" value={data.individual_id} />
            <DetailItem label="Lifestage" value={data.lifestage} />
            <DetailItem label="Sex" value={data.sex} />
            <DetailItem label="Body Length" value={data.body_length} />
            <DetailItem label="Language ID" value={data.language_id} />
            <DetailItem label="Date Created" value={formatDate(data.created_at)} />
            <DetailItem label="Note" value={data.note} />
          </dl>
          <hr className="my-4"/>

          {/* --- 場所情報 --- */}
          <h2 className="text-lg font-medium text-gray-700 mb-2">Location</h2>
          <dl className="grid grid-cols-1 md:grid-cols-3 gap-x-4 gap-y-1 divide-y divide-gray-200">
            <DetailItem label="Place Name" value={data.place_name} />
            <DetailItem label="Latitude" value={data.latitude} />
            <DetailItem label="Longitude" value={data.longitude} />
          </dl>
          <hr className="my-4"/>

          {/* --- 分類情報 --- */}
          {data.classification && (
            <>
              <h2 className="text-lg font-medium text-gray-700 mb-2">Classification</h2>
              <dl className="grid grid-cols-2 md:grid-cols-4 gap-x-4 gap-y-1 divide-y divide-gray-200">
                <DetailItem label="Kingdom" value={data.classification.kingdom} />
                <DetailItem label="Phylum" value={data.classification.phylum} />
                <DetailItem label="Class" value={data.classification.class} />
                <DetailItem label="Order" value={data.classification.order} />
                <DetailItem label="Family" value={data.classification.family} />
                <DetailItem label="Genus" value={data.classification.genus} />
                <DetailItem label="Species" value={data.classification.species} />
                <DetailItem label="Others" value={data.classification.others} />
              </dl>
              <hr className="my-4"/>
            </>
          )}

          {/* --- 観察情報 (リスト) --- */}
          <h2 className="text-lg font-medium text-gray-700 mb-2">Observations ({data.observation.length})</h2>
          {data.observation.map((obs, index) => (
            <div key={obs.observation_id} className={`py-4 ${index < data.observation.length - 1 ? 'border-b' : ''}`}>
              <dl className="grid grid-cols-1 md:grid-cols-3 gap-x-4 gap-y-1">
                <DetailItem label="Observer" value={obs.observation_user} />
                <DetailItem label="Method" value={obs.observation_method_name} />
                <DetailItem label="Observed At" value={formatDate(obs.observed_at)} />
              </dl>
              <DetailItem label="Behavior" value={obs.behavior} />
            </div>
          ))}
          <hr className="my-4"/>

          {/* --- 標本情報 (リスト) --- */}
          <h2 className="text-lg font-medium text-gray-700 mb-2">Specimens ({data.specimen.length})</h2>
          {data.specimen.map((spec, index) => (
             <div key={spec.specimen_id} className={`py-4 ${index < data.specimen.length - 1 ? 'border-b' : ''}`}>
              <dl className="grid grid-cols-1 md:grid-cols-3 gap-x-4 gap-y-1">
                <DetailItem label="Preparator" value={spec.specimen_user} />
                <DetailItem label="Method" value={spec.specimen_methods_common} />
                 <DetailItem label="Date Prepared" value={formatDate(spec.created_at)} />
                <DetailItem label="Institution" value={spec.institution_code} />
                <DetailItem label="Collection ID" value={spec.collection_id} />
              </dl>
            </div>
          ))}
          <hr className="my-4"/>

          {/* --- 同定情報 (リスト) --- */}
          <h2 className="text-lg font-medium text-gray-700 mb-2">Identifications ({data.identification.length})</h2>
           {data.identification.map((ident, index) => (
             <div key={ident.identification_id} className={`py-4 ${index < data.identification.length - 1 ? 'border-b' : ''}`}>
              <dl className="grid grid-cols-1 md:grid-cols-3 gap-x-4 gap-y-1">
                <DetailItem label="Identifier" value={ident.identification_user} />
                <DetailItem label="Date Identified" value={formatDate(ident.identified_at)} />
              </dl>
              <DetailItem label="Source Info" value={ident.source_info} />
            </div>
          ))}
          <hr className="my-4"/>

          {/* --- 添付ファイル情報 (リスト) --- */}
          <h2 className="text-lg font-medium text-gray-700 mb-2">Attachments ({data.attachments.length})</h2>
          <ul className="list-disc pl-5 text-sm space-y-1">
            {data.attachments.map(att => (
              <li key={att.attachment_id}>
                {/* ここは後で実際のファイルへのリンクにするのだ */}
                <span className="text-blue-600 hover:underline">{att.file_name || att.file_path}</span>
                {att.note && <span className="text-gray-500 text-xs ml-2">({att.note})</span>}
              </li>
            ))}
          </ul>

        </div>
      </div>
    </div>
  );
}
