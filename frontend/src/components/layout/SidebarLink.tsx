// src/components/layout/SidebarLink.tsx

import { ReactNode } from "react";

type Props = {
  icon: ReactNode;
  text: string;
  href: string;
  active: boolean;
  isOpen: boolean; // 親から開閉状態を受け取る
};

const SidebarLink = ({ icon, text, href, active, isOpen }: Props) => {
  return (
    <a
      href={href}
      className={`
        flex items-center p-3 rounded-lg transition-colors
        ${active
          ? 'bg-gray-200 text-gray-900 font-semibold'
          : 'text-gray-600 hover:bg-gray-100'
        }
      `}
    >
      {icon}
      {/* isOpenがtrueの時だけ、文字を表示するのだ */}
      {isOpen && <span className="ml-4 whitespace-nowrap">{text}</span>}
    </a>
  );
};

export default SidebarLink;
