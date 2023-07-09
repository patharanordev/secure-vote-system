"use client"

import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Container from '@mui/material/Container';
import { User } from '#/ui/account/user';
import { useSession } from 'next-auth/react';

type Props = {
  onAddItem?: Function
}

function ResponsiveAppBar(props: Props) {
  const { data: session } = useSession();
  const onAddItem = () => {
    if (typeof props.onAddItem === 'function') {
      props.onAddItem()
    }
  }

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
          }} onAddItem={onAddItem} />
        </Toolbar>
      </Container>
    </AppBar>
  );
}
export default ResponsiveAppBar;