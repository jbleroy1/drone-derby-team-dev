import React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import getCollection from '../../firebase/getCollection';

function DroneDiagnostics() {
  type droneDiagnostics = {
    droneMacAddress: string;
    status: string;
    version: string;
    battery: number;
  };

  const [rows, setRows] = React.useState<droneDiagnostics[]>([]);
  getCollection(setRows, rows, 'DroneDiagnostics');

  console.log(rows);

  return (
    <>
      {rows[0] && (
        <Card>
          <CardContent>
            <Typography className="mb-4 text-base font-medium" gutterBottom>
              Drone Status
            </Typography>
            <Typography className="text-base" color="text.primary">
              Drone Mac Address
            </Typography>
            <Typography className="text-sm mb-2" color="text.secondary">
              {rows[0].droneMacAddress}
            </Typography>
            <Typography className="text-base" color="text.primary">
              Status
            </Typography>
            <Typography className="text-sm mb-2" color="text.secondary">
              {rows[0].status}
            </Typography>
            <Typography className="text-base" color="text.primary">
              Version
            </Typography>
            <Typography className="text-sm mb-2" color="text.secondary">
              {rows[0].version}
            </Typography>
            <Typography className="text-base" color="text.primary">
              Battery
            </Typography>
            <Typography className="text-sm" color="text.secondary">
              {rows[0].battery}%
            </Typography>
          </CardContent>
        </Card>
      )}
    </>
  );
}

export default DroneDiagnostics;
