import { Bell, LocateFixed, User, Languages } from 'lucide-react';

const Header = () => {
  return (
    <header className="bg-white shadow-sm p-4 flex justify-between items-center">
      <div>
        <h1 className="text-xl text-gray-900 font-semibold">Web-Specimen</h1>
        <p className="text-sm text-gray-500">Database-name</p>
      </div>
      <div className="flex items-center space-x-4">
        <button className="p-2 rounded-full hover:bg-gray-100 text-gray-600"><LocateFixed size={20} /></button>
        <button className="p-2 rounded-full hover:bg-gray-100 text-gray-600 relative">
          <Bell size={20} />
          <span className="absolute top-0 right-0 block h-2 w-2 rounded-full bg-red-500"></span>
        </button>
        <button className="p-2 rounded-full hover:bg-gray-100 text-gray-600 relative">
          <Languages size={20} />
        </button>
        <div className="flex items-center space-x-2">
          <div className="p-2 rounded-full bg-gray-200 text-gray-600"><User size={20} /></div>
          <div>
            <p className="text-sm text-gray-700 font-medium">user name</p>
            <p className="text-xs text-gray-500">user role</p>
          </div>
        </div>
      </div>
    </header>
  );
};

export default Header;
