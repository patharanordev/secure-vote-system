"use client"

import * as React from 'react';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import { createTheme, styled } from '@mui/material/styles';
import { VoteItemProps } from "#/types"

const Item = styled(Paper)(({ theme }) => ({
    ...theme.typography.body2,
    color: theme.palette.text.secondary,
    height: 120,
    lineHeight: '60px',
}));

export default function VoteItem(props: VoteItemProps) {
    return (
        <Item elevation={3}>
            <Grid container spacing={2}>
                <Grid item xs={9}>{props.itemName}</Grid>
                <Grid item xs={3}>{props.voteCount}</Grid>
                <Grid item xs={12}>{props.itemDescription}</Grid>
            </Grid>
        </Item>
    )
}