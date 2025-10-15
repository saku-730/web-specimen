// src/app/page.tsx

import StatsCard from "@/components/dashboard/StatsCard";
import RecentDataTable from "@/components/dashboard/RecentDataTable";
import { DollarSign, Users, Package, Activity } from "lucide-react";

export default function HomePage() {
  return (
    <div className="space-y-8">
      {/* --- 統計カードのセクション --- */}
      <div>
        <h1 className="text-2xl font-semibold mb-4">Dashboard</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <StatsCard 
            title="Total Data" 
            value="" 
            change="" 
            icon={<Package size={24} className="text-blue-500" />} 
          />
          <StatsCard 
            title="" 
            value="" 
            change="" 
            icon={<Users size={24} className="text-green-500" />} 
          />
          <StatsCard 
            title="" 
            value="" 
            change="" 
            icon={<Activity size={24} className="text-yellow-500" />} 
          />
          <StatsCard 
            title="" 
            value="" 
            change="" 
            icon={<DollarSign size={24} className="text-red-500" />} 
          />
        </div>
      </div>

      {/* --- 最近のデータのテーブル --- */}
      <RecentDataTable />
    </div>
  );
}
