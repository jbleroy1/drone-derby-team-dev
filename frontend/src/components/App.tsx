import React, { useEffect } from 'react';
import { Routes, Route } from 'react-router-dom';
import NavBar from './navigation/NavBar';
import Mappings from './mappings/Mappings';
import Diagnostics from './diagnostics/Diagnostics';
import Control from './control/Control';
import { UserAuth } from '../contexts/AuthContext';
import { Typography } from '@mui/material';

function App() {
  const { signIn, user } = UserAuth();

  useEffect(() => {
    try {
      signIn();
    } catch (error) {
      console.error(error);
    }
  });

  return (
    <>
      {user && (
        <>
          <NavBar />
          <Routes>
            <Route path="/" element={<Mappings />} />
            <Route path="controls" element={<Control />} />
            <Route path="diagnostics" element={<Diagnostics />} />
          </Routes>
          <div className="fixed bottom-0 w-full">
            <Typography variant="caption" component="div" className="text-center">
              Signed in Anonymously ({user.uid})
            </Typography>
          </div>
        </>
      )}
    </>
  );
}

export default App;
