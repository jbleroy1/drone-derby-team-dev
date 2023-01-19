import React from 'react';
import Divider from '@mui/material/Divider';
import { Typography } from '@mui/material';
import CoralDiagnostics from './CoralDiagnostics';
import DroneDiagnostics from './DroneDiagnostics';
import Operations from './Operations';

const Diagnostics = () => {
  return (
    <div className="m-3">
      <Typography color="text.secondary" variant="h6" className="ml-6 mt-4 mb-1">
        Diagnostics & Operations
      </Typography>
      <Divider variant="middle" />
      <div className="justify-center">
        <div className="flex justify-center gap-8 mt-4">
          <div className="w-2/4 ml-10">
            <DroneDiagnostics />
          </div>
          <div className="w-2/4 mr-10">
            <CoralDiagnostics />
          </div>
        </div>
      </div>
      <div className="justify-center">
        <div className="flex justify-center gap-8 mt-4">
          <div className="w-2/4 ml-10">
            <Operations />
          </div>
          <div className="w-2/4 mr-10">
            <Operations />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Diagnostics;
