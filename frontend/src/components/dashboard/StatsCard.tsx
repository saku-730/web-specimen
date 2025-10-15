type Props = {
  title: string;
  value: string;
  change: string;
};

const StatsCard = ({ title, value, change }: Props) => {
  return (
    <div className="bg-white p-6 rounded-lg shadow">
      <p className="text-sm text-gray-500">{title}</p>
      <p className="text-3xl font-bold mt-1">{value}</p>
      <p className="text-sm text-green-500 mt-2">{change}</p>
    </div>
  );
};

export default StatsCard;
