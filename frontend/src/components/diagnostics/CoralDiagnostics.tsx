import React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';

function CoralDiagnostics() {
  return (
    <>
      <Card>
        <CardContent>
          <Typography className="mb-4 text-lg font-medium" gutterBottom>
            Coral Status
          </Typography>
          <Typography className="text-base" color="text.primary">
            Drone Mac Address
          </Typography>
          <Typography className="text-sm mb-2" color="text.secondary">
            00:00:5e:00:53:af
          </Typography>
          <Typography className="text-base" color="text.primary">
            Status
          </Typography>
          <Typography className="text-sm mb-2" color="text.secondary">
            Connected
          </Typography>
          <Typography className="text-base" color="text.primary">
            Version
          </Typography>
          <Typography className="text-sm mb-2" color="text.secondary">
            1.0.3-alpha
          </Typography>
          <Typography className="text-base" color="text.primary">
            Battery
          </Typography>
          <Typography className="text-sm" color="text.secondary">
            84%
          </Typography>
        </CardContent>
      </Card>
    </>
  );
}

export default CoralDiagnostics;
