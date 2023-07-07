"use client"

import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import Menu from '@mui/material/Menu';
import MenuIcon from '@mui/icons-material/Menu';
import Container from '@mui/material/Container';
import Button from '@mui/material/Button';
import MenuItem from '@mui/material/MenuItem';
import AdbIcon from '@mui/icons-material/Adb';
import { User } from '#/ui/account/user';
import { Account } from "#/types"
import { useSession } from 'next-auth/react';

function ResponsiveAppBar() {
  const { data: session } = useSession();

  return (
    <AppBar position="static">
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          <Box sx={(theme) => ({
              [theme.breakpoints.up("sm")]: {
                flexGrow: 1
              },
            })}
          />
          <User profile={{
            name: session?.user?.name,
            image: session?.user?.image,
            email: session?.user?.email
          }} />
        </Toolbar>
      </Container>
    </AppBar>
  );
}
export default ResponsiveAppBar;