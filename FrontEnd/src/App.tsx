import { useEffect, useState } from 'react';

import { CpuChart } from './components/CpuChart';
import { useWebSocket } from './hooks/useWebSocket';
import { Measurement } from './type';
import { MemoryChart } from './components/MemoryChart';

function App() {
  const { message } = useWebSocket(import.meta.env.VITE_API_URL);
  const [cpuData, setCpuData] = useState<Measurement[]>();
  const [memoryData, setMemoryData] = useState<Measurement[]>();

  useEffect(() => {
    setCpuData(message?.Cpu.measurements);
    setMemoryData(message?.Memory.measurements);
  }, [message]);

  return (
    <>
      <div className="container mx-auto gap-y-6">
        <CpuChart data={cpuData ?? []} />

        <MemoryChart data={memoryData ?? []} />
      </div>
    </>
  );
}

export default App;
