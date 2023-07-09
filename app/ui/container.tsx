"use client"

import * as React from 'react';
import CssBaseline from '@mui/material/CssBaseline';
import Container from '@mui/material/Container';

type Props = {
    children?: React.ReactNode
}

export default function MainContainer(props: Props) {
    return (
        <React.Fragment>
            <CssBaseline />
            <Container fixed sx={{
                width: '100vw !important',
                height: '100vh !important',
                padding: '0px !important',
                margin: '0px !important',
                maxWidth: 'unset !important',
            }}>
                {props?.children}
            </Container>
        </React.Fragment>
    );
}