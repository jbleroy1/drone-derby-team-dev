import React from 'react';
import Divider from '@mui/material/Divider';
import { Typography } from '@mui/material';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import RotateLeftIcon from '@mui/icons-material/RotateLeft';
import RotateRightIcon from '@mui/icons-material/RotateRight';

const Control = () => {
  return (
    <div className="m-3">
      <Typography color="text.secondary" variant="h6" className="ml-6 mt-4 mb-1">
        Flight Controls
      </Typography>
      <Divider variant="middle" />
      <div className="flex justify-center gap-8 mt-4">
        <div className="w-6/12">
          <Card>
            <CardContent>
              <div className="flex gap-2">
                <div className="w-2/4">
                  <Card className="h-full flex justify-center">
                    <CardContent>
                      <div className="flex h-full justify-center gap-2">
                        <div>
                          <Button
                            className="text-center h-full"
                            variant="contained"
                            color="primary"
                          >
                            <RotateLeftIcon />
                          </Button>
                        </div>
                        <div>
                          <Button
                            className="text-center h-full"
                            variant="contained"
                            color="primary"
                          >
                            <RotateRightIcon />
                          </Button>
                        </div>
                      </div>
                    </CardContent>
                  </Card>
                </div>
                <div className="w-2/4">
                  <Card>
                    <CardContent>
                      <div className="flex flex-col items-center gap-2 mt-4">
                        <Button className="w-48 text-center" variant="contained" color="success">
                          Take Off
                        </Button>
                        <Button className="w-48 text-center" variant="contained" color="info">
                          Land
                        </Button>

                        <Button className="w-48 text-center" variant="contained" color="error">
                          Stop
                        </Button>
                      </div>
                    </CardContent>
                  </Card>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
};

export default Control;
