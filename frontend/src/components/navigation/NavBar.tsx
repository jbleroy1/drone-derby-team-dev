import React, { useState } from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import Drawer from '@mui/material/Drawer';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemText from '@mui/material/ListItemText';
import { ListItemIcon } from '@mui/material';
import ForkRightTwoToneIcon from '@mui/icons-material/ForkRightTwoTone';
import ControlCameraTwoToneIcon from '@mui/icons-material/ControlCameraTwoTone';
import MonitorHeartOutlinedIcon from '@mui/icons-material/MonitorHeartOutlined';
import { Link } from 'react-router-dom';

const NavBar = () => {
  const [isOpen, setIsOpen] = useState(false);

  const toggleDrawer = () => () => {
    setIsOpen((prev) => !prev);
  };

  const list = () => (
    <Box sx={{ width: 300, mt: 8 }} role="presentation" onClick={toggleDrawer()}>
      <List>
        <Link to="/">
          <ListItem disablePadding>
            <ListItemButton>
              <ListItemIcon>
                <ForkRightTwoToneIcon />
              </ListItemIcon>
              <ListItemText primary={'Sign Mappings'} />
            </ListItemButton>
          </ListItem>
        </Link>
        <Link to="/controls">
          <ListItem disablePadding>
            <ListItemButton>
              <ListItemIcon>
                <ControlCameraTwoToneIcon />
              </ListItemIcon>
              <ListItemText primary={'Flight Controls'} />
            </ListItemButton>
          </ListItem>
        </Link>
        <Link to="/diagnostics">
          <ListItem disablePadding>
            <ListItemButton>
              <ListItemIcon>
                <MonitorHeartOutlinedIcon />
              </ListItemIcon>
              <ListItemText primary={'Diagnostics & Operations'} />
            </ListItemButton>
          </ListItem>
        </Link>
      </List>
    </Box>
  );

  return (
    <>
      <Box sx={{ flexGrow: 1 }}>
        <AppBar position="relative" sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}>
          <Toolbar>
            <IconButton
              onClick={toggleDrawer()}
              size="large"
              edge="start"
              color="inherit"
              aria-label="menu"
              sx={{ mr: 2 }}
            >
              <MenuIcon />
            </IconButton>
            <Typography variant="h6" component="div" className="text-center" sx={{ flexGrow: 1 }}>
              Google Cloud Drone Derby
            </Typography>
          </Toolbar>
        </AppBar>
      </Box>
      <Drawer anchor={'left'} open={isOpen} onClose={toggleDrawer()}>
        {list()}
      </Drawer>
    </>
  );
};

export default NavBar;
