import { AreaChart, Card } from '@tremor/react';
import { Measurement } from '../type';
interface props {
  data: Measurement[];
}
export function CpuChart({ data }: props) {
  const sortedData = data.sort((a, b) => {
    const dateA = new Date(a.date).getTime();
    const dateB = new Date(b.date).getTime();
    return dateB - dateA; //sort desc
  });
  //get 50 elements
  const firstElements = sortedData.slice(0, 50);
  return (
    <>
      <div className="w-full p-5">
        <Card>
          <p className="text-tremor-default text-tremor-content dark:text-dark-tremor-content">CPU Usage</p>
          <AreaChart
            className="mt-4 h-72"
            data={firstElements}
            index="formattedDate"
            categories={['value']}
            colors={['blue']}
            yAxisWidth={65}
          />
        </Card>
      </div>
    </>
  );
}
