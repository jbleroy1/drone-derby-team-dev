import React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import LaunchIcon from '@mui/icons-material/Launch';
import Link from '@mui/material/Link';
import getCollection from '../../firebase/getCollection';

function Operations() {
  type operationsTypes = {
    logging: string;
    monitoring: string;
    slos: string;
  };

  const [operations, setOperations] = React.useState<operationsTypes[]>([]);
  getCollection(setOperations, operations, 'Operations');

  return (
    <>
      {operations[0] && (
        <Card>
          <CardContent>
            <Typography className="mb-4 text-base font-medium" gutterBottom>
              Operations
            </Typography>
            <Typography className="text-base mb-2" color="text.primary">
              <Link href={operations[0].logging} target="_blank" color="inherit" underline="none">
                Logging <LaunchIcon fontSize="small"></LaunchIcon>
              </Link>
            </Typography>
            <Typography className="text-base mb-2" color="text.primary">
              <Link
                href={operations[0].monitoring}
                target="_blank"
                color="inherit"
                underline="none"
              >
                Monitoring <LaunchIcon fontSize="small"></LaunchIcon>
              </Link>
            </Typography>
            <Typography className="text-base mb-2" color="text.primary">
              <Link href={operations[0].slos} target="_blank" color="inherit" underline="none">
                SLOs <LaunchIcon fontSize="small"></LaunchIcon>
              </Link>
            </Typography>
          </CardContent>
        </Card>
      )}
    </>
  );
}

export default Operations;
