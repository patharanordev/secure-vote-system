"use client"

import * as React from 'react';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import { styled } from '@mui/material/styles';
import { VoteItemProps } from "#/types"

const Item = styled(Paper)(({ theme }) => ({
    ...theme.typography.body2,
    color: theme.palette.text.secondary,
    height: 196,
    lineHeight: '60px',
}));

const Content = styled(Paper)(({ theme }) => ({
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'left',
    color: theme.palette.text.secondary,
    backgroundColor: "#e7e7e7",
}));

export default function VoteItem(props: VoteItemProps) {
    return (
        <Item elevation={3} sx={{px:1}}>
            <Grid container spacing={1}>
                <Grid item xs={9}>
                    <Content sx={{
                        fontSize: 20,
                    }}>{props.itemName}</Content>
                </Grid>
                <Grid item xs={3}>
                    <Content sx={{
                        textAlign: 'center',
                        fontSize: 20,
                        display: "flex",
                        justifyContent: "center",
                    }}><b>{props.voteCount}</b></Content>
                </Grid>
                <Grid item xs={12}>
                    <Content sx={{
                        overflowY: "scroll",
                        height: 120,
                        fontSize: 14,
                }}>{props.itemDescription}</Content>
                </Grid>
            </Grid>
        </Item>
    )
}