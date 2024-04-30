import { useEffect, useRef, useState } from 'react';
import { CharInfo, Message } from '../type';
import { format } from 'date-fns';

export const useWebSocket = (url: string) => {
  const [message, setMessage] = useState<CharInfo>();
  const webSocket = useRef<WebSocket | null>(null);

  useEffect(() => {
    webSocket.current = new WebSocket(url);

    webSocket.current.onopen = () => {
      webSocket.current?.send(JSON.stringify(['1', '1', 'broadcast:lobby', 'phx_join', {}]));
    };

    webSocket.current.onmessage = (event: MessageEvent<string>) => {
      const data = JSON.parse(event.data);
      if (data[3]) {
        const eventType = data[3];
        const payload = data[4];
        let cpu: Message = { name: 'cpu', measurements: [] };

        let memory: Message = { name: 'memory', measurements: [] };
        if (eventType === 'new') {
          payload.message.map((item: Message) => {
            const formattedMeasurements = item.measurements.map((measurement) => {
              return {
                ...measurement,
                formattedDate: format(new Date(measurement.date), 'mm:ss'),
              };
            });

            if (item.name === 'CPU') {
              cpu = { name: 'cpu', measurements: formattedMeasurements };
            } else {
              memory = { name: 'memory', measurements: formattedMeasurements };
            }
          });
          const ms: CharInfo = { Cpu: cpu, Memory: memory };
          setMessage(ms);
        }
      }
    };

    //cleanup
    return () => {
      if (webSocket.current) {
        webSocket.current.close();
      }
    };
  }, [url]);

  return { message };
};
