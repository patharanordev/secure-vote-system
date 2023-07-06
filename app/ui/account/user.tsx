import { Avatar, Box, IconButton, Menu, MenuItem, Tooltip, Typography } from "@mui/material";
import { Account, EUserMenu } from "#/types"
import React from "react";
import { signOut } from "next-auth/react"

const settings = [EUserMenu.Logout];

export const User = (props: Account) => {
    const [anchorElUser, setAnchorElUser] = React.useState<null | HTMLElement>(null);
  
    const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
      setAnchorElUser(event.currentTarget);
    };
  
    const handleCloseUserMenu = (menu: string) => {
        switch(menu) {
            case EUserMenu.Logout:
                signOut();
                break;
            default:
                break;
        }

        setAnchorElUser(null);
    };
    return (
        <Box sx={{ display: { xs:'flex' } }}>
            <Tooltip title="Profile">
                <IconButton sx={{ p: 0 }} onClick={handleOpenUserMenu}>
                    <Avatar alt="user-image" src="/images/avatar/2.jpg" />
                </IconButton>
            </Tooltip>
            <Menu
                sx={{ mt: '45px' }}
                id="menu-appbar"
                anchorEl={anchorElUser}
                anchorOrigin={{
                    vertical: 'top',
                    horizontal: 'right',
                }}
                keepMounted
                transformOrigin={{
                    vertical: 'top',
                    horizontal: 'right',
                }}
                open={Boolean(anchorElUser)}
                onClose={handleCloseUserMenu}
            >
                {settings.map((setting) => (
                    <MenuItem key={setting} onClick={() => handleCloseUserMenu(setting)}>
                        <Typography textAlign="center">{setting}</Typography>
                    </MenuItem>
                ))}
            </Menu>
            <Typography sx={{ p: 2 }}>
                {props?.profile?.name}
            </Typography>
        </Box>
    );
};