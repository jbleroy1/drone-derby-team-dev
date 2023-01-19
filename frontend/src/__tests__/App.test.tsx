import React from 'react';
import { render, screen } from '@testing-library/react';
import App from '../components/App';

test('renders drone-derby text', () => {
  render(<App />);
  const linkElement = screen.getByText(/drone-derby/i);
  expect(linkElement).toBeInTheDocument();
});
