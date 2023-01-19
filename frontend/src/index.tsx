import React from 'react';
import ReactDOM from 'react-dom/client';
import './styles/index.css';
import { AuthContextProvider } from './contexts/AuthContext';
import { StyledEngineProvider } from '@mui/material/styles';
import { BrowserRouter } from 'react-router-dom';
import App from './components/App';

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);
root.render(
  <React.StrictMode>
    <AuthContextProvider>
      <StyledEngineProvider injectFirst>
        <BrowserRouter>
          <App />
        </BrowserRouter>
      </StyledEngineProvider>
    </AuthContextProvider>
  </React.StrictMode>
);
