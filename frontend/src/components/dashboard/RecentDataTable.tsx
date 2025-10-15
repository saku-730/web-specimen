import { Settings } from 'lucide-react';

const RecentDataTable = () => {
  const data = [
    { name: "", date: "", phone: "", location: "", registered: true },
    { name: "", date: "", phone: "", location: "", registered: false },
    { name: "", date: "", phone: "", location: "", registered: true },
    { name: "", date: "", phone: "", location: "", registered: false },
  ];

  return (
    <div className="bg-white p-6 rounded-lg shadow">
      <div className="flex justify-between items-center mb-4">
	<h2 className="text-lg font-semibold text-gray-800">Recent Data</h2>
	<button className="p-2 rounded-full hover:bg-gray-100 text-gray-600"><Settings size={20} /></button>
      </div>
      <table className="w-full text-sm text-left">
	<thead className="text-xs text-gray-500 uppercase bg-gray-50">
	  <tr>
	    <th scope="col" className="px-6 py-3">Name</th>
	    <th scope="col" className="px-6 py-3">Order Date</th>
	    <th scope="col" className="px-6 py-3">Phone Number</th>
	    <th scope="col" className="px-6 py-3">Location</th>
	    <th scope="col" className="px-6 py-3">Registered</th>
	    <th scope="col" className="px-6 py-3"></th>
	  </tr>
	</thead>
	<tbody>
	  {data.map((item, index) => (
	    <tr key={index} className="bg-white border-b">
	      <td className="px-6 py-4 font-medium text-gray-900">{item.name}</td>
	      <td className="px-6 py-4 text-gray-600">{item.date}</td>
	      <td className="px-6 py-4 text-gray-600">{item.phone}</td>
	      <td className="px-6 py-4 text-gray-600">{item.location}</td>
	      <td className="px-6 py-4">{item.registered ? "Yes" : "No"}</td>
	      <td className="px-6 py-4 text-right space-x-4">
		<a href="#" className="font-medium text-blue-600 hover:underline">Options</a>
		<a href="#" className="font-medium text-blue-600 hover:underline">Details</a>
	      </td>
	    </tr>
	  ))}
	</tbody>
      </table>
      {/* ここに後でページネーションコンポーネントを追加するのだ */}
    </div>
  );
};

export default RecentDataTable;
