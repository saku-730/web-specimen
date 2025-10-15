// src/app/create/page.tsx

import OccurrenceForm from "@/components/create/OccurrenceForm";

export default function CreatePage() {
  return (
    <div className="bg-[#dcdcdc] min-h-screen p-4 md:p-8">
      <h1 className="text-2xl font-semibold mb-6">Create New Occurrence</h1>
      <div className="bg-white p-8 rounded-lg shadow-md">
        <OccurrenceForm />
      </div>
    </div>
  );
}
