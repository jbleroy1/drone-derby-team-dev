import { Divider, Typography } from '@mui/material';
import SignMappingsDataGrid from './SignMappingsDataGrid';
import React from 'react';

const Mappings = () => {
  return (
    <div className="m-3">
      <Typography color="text.secondary" variant="h6" className="ml-6 mt-4 mb-1">
        Sign Mappings
      </Typography>
      <Divider variant="middle" />
      <div className="ml-5 mr-5 mt-5">
        <SignMappingsDataGrid />
      </div>
    </div>
  );
};

export default Mappings;
