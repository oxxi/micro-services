export interface Message {
  name: string;
  measurements: Measurement[];
}

export interface Measurement {
  value: number;
  date: Date;
  formattedDate?: string;
}

export interface CharInfo {
  Cpu: Message;
  Memory: Message;
}
