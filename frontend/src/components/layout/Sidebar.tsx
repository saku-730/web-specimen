// src/components/layout/Sidebar.tsx

"use client";

import { usePathname } from 'next/navigation'; 
import { useState } from "react";
import { Home, Search, Settings, Info, SquarePlus, X, Menu } from 'lucide-react'; // アイコンを少し変更
import SidebarLink from "./SidebarLink";

const Sidebar = () => {
  const [isOpen, setIsOpen] = useState(false);

  const pathname = usePathname();

  const navLinks = [
    { icon: <Home size={20} />, text: "Dashboard", href: "/"},
    { icon: <Search size={20} />, text: "Search", href: "/search"},
    { icon: <SquarePlus size={20} />, text: "Create", href: "/create"},
    { icon: <Settings size={20} />, text: "Settings", href: "/settings"},
    { icon: <Info size={20} />, text: "Info", href: "/info"},
  ];

 const navItems = navLinks.map((item) => ({
    ...item,
    active: item.href === pathname, 
  }));

  return (
    <aside 
      className={`
        absolute top-0 left-0 h-full z-30 bg-white shadow-lg flex flex-col 
        transition-all duration-300 ease-in-out 
        ${isOpen ? 'w-64' : 'w-20'}
      `}
    >
      <div className="flex flex-col items-start">
        <button 
          onClick={() => setIsOpen(!isOpen)} 
          className="p-6 text-gray-600 hover:text-gray-900"
        >
          {isOpen ? <X size={24} /> : <Menu size={24} />}
        </button>
      </div>

     
      <nav className="flex flex-col space-y-2 px-4">
        {navItems.map((item) => (
          <SidebarLink
            key={item.text}
            icon={item.icon}
            text={item.text}
            href={item.href}
            active={item.active}
            isOpen={isOpen}
          />
        ))}
      </nav>
    </aside>
  );
};

export default Sidebar;
