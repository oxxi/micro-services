import { Card, LineChart } from '@tremor/react';
import { Measurement } from '../type';
interface props {
  data: Measurement[];
}
export function MemoryChart({ data }: props) {
  const sortedData = data.sort((a, b) => {
    const dateA = new Date(a.date).getTime();
    const dateB = new Date(b.date).getTime();
    return dateB - dateA;
  });
  //get 50 elements
  const firstElements = sortedData.slice(0, 50);
  return (
    <>
      <div className="w-full p-5">
        <Card>
          <p className="text-tremor-default text-tremor-content dark:text-dark-tremor-content">
            Memory Usage
          </p>

          <LineChart
            className="mt-4 h-72"
            data={firstElements}
            index="formattedDate"
            categories={['value']}
            colors={['indigo', 'rose']}
            yAxisWidth={60}
          />
        </Card>
      </div>
    </>
  );
}
