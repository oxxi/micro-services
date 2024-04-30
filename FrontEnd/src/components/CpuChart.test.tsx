import { render, screen } from '@testing-library/react';
import { CpuChart } from './CpuChart';
import { Measurement } from '../type';

beforeAll(() => {
  global.ResizeObserver = class ResizeObserver {
    observe() {
      // do nothing
    }
    unobserve() {
      // do nothing
    }
    disconnect() {
      // do nothing
    }
  };
});

const DATA: Measurement[] = [
  {
    date: new Date(),
    value: 90,
    formattedDate: '11:05',
  },
  {
    date: new Date(),
    value: 10,
    formattedDate: '12:05',
  },
];

describe('<CpuChart/>', () => {
  it('Should render Chart', () => {
    render(<CpuChart data={DATA} />);
    expect(screen.getByText('CPU Usage')).toBeDefined();
  });
});
