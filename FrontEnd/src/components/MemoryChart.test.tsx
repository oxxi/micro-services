import { render, screen } from '@testing-library/react';
import { Measurement } from '../type';
import { MemoryChart } from './MemoryChart';

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

describe('<MemoryChart/>', () => {
  it('Should render Memory Chart', () => {
    render(<MemoryChart data={DATA} />);
    expect(screen.getByText('Memory Usage')).toBeDefined();
  });
});
