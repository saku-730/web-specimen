// src/components/search/SearchResults.tsx

type Props = {
  data: {
    occurrence_results: any[];
    metadata: {
      total_results: number;
      current_page: number;
      per_page: number;
      total_pages: number;
    };
  } | null;
};

const SearchResults = ({ data }: Props) => {
  if (!data) {
    return <p className="text-center text-gray-500">Please enter search criteria and click "Search".</p>;
  }

  if (data.occurrence_results.length === 0) {
    return <p className="text-center text-gray-500">No results found.</p>;
  }

  const { occurrence_results, metadata } = data;

  return (
    <div>
      <p className="text-sm text-gray-600 mb-4">
        Showing page {metadata.current_page} of {metadata.total_pages} ({metadata.total_results} total results)
      </p>
      
      {/* --- 結果をテーブルで表示 --- */}
      <div className="overflow-x-auto">
        <table className="w-full text-sm text-left">
          <thead className="text-xs text-gray-500 uppercase bg-gray-50">
            <tr>
              <th className="px-6 py-3">ID</th>
              <th className="px-6 py-3">Species</th>
              <th className="px-6 py-3">Project</th>
              <th className="px-6 py-3">User</th>
              <th className="px-6 py-3">Date</th>
            </tr>
          </thead>
          <tbody>
            {occurrence_results.map((item: any) => (
              <tr key={item.occurrence_id} className="bg-white border-b hover:bg-gray-50">
                <td className="px-6 py-4 font-medium">{item.occurrence_id}</td>
                <td className="px-6 py-4">{item.classification?.species}</td>
                <td className="px-6 py-4">{item.project_name}</td>
                <td className="px-6 py-4">{item.user_name}</td>
                <td className="px-6 py-4">{new Date(item.created_at).toLocaleDateString()}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      
      {/* ここに後でページネーションのボタンを実装するのだ */}
    </div>
  );
};

export default SearchResults;
