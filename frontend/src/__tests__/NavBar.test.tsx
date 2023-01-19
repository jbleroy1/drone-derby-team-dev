import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import { BrowserRouter as Router } from 'react-router-dom';
import NavBar from '../components/navigation/NavBar';

// Verifies that the Nav bar is rendered
test('renders NavBar', () => {
  render(<NavBar />);
  const linkElement = screen.getByText(/Google Cloud Drone Derby/i);
  expect(linkElement).toBeInTheDocument();

  const buttonElement = screen.getByRole('button');
  expect(buttonElement).toBeInTheDocument();
});

// When the App Bar hamburger button is clicked the Drawer should be visible
test('renders Drawer', () => {
  render(
    <Router>
      <NavBar />
    </Router>
  );
  const buttonElement = screen.getByRole('button');

  fireEvent.click(buttonElement);
  const navItems = screen.getAllByRole('listitem');
  expect(navItems.length).toBe(3);
});

test('verify links are present in Drawer', () => {
  render(
    <Router>
      <NavBar />
    </Router>
  );
  const buttonElement = screen.getByRole('button');

  fireEvent.click(buttonElement);

  const links: HTMLAnchorElement[] = screen.getAllByRole('link');

  expect(links[0].textContent).toEqual('Sign Mappings');
  expect(links[0].href).toContain('/');

  expect(links[1].textContent).toEqual('Flight Controls');
  expect(links[1].href).toContain('/controls');

  expect(links[2].textContent).toEqual('Diagnostics & Operations');
  expect(links[2].href).toContain('/diagnostics');
});
