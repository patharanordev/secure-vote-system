"use client"

import * as React from 'react';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import { styled } from '@mui/material/styles';
import { 
    VoteItemProps, 
    VoteItemPayload,
    VotingPayload 
} from "#/types"

import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import IconButton from '@mui/material/IconButton';
import Stack from '@mui/material/Stack';
import { useSession } from 'next-auth/react';
import { useEffect, useState } from 'react';

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
    const { data: session } = useSession();
    const [voteCount, setVoteCount] = useState<number>(props.voteCount);
    const [payload, setPayload] = useState<VoteItemPayload|null>(null);
    const [token, setToken] = useState<string|null>(null);

    const getItemById = async (id: string) => {
        const payload: VoteItemPayload = await fetch(`/api/vote-item?id=${id}`, {
            headers: { "Authorization": `Bearer ${token}` },
            method: "GET",
        }).then((res) => res.json());

        console.log('Vote item by id response:', payload);

        setVoteCount(payload?.data?.info?.voteCount ?? 0);
    }

    const vote = async (payload: VotingPayload) => {
        const res: VotingPayload = await fetch('/api/voting', {
            method: "PATCH",
            body: JSON.stringify({ ...payload }),
            headers: { 
                "Content-Type": "application/json", 
                "Authorization": `Bearer ${token}`
            }
        }).then((res) => res.json());

        // console.log('VoteItem session:', token);
        console.log('VoteItem res:', res);
        setPayload(res);
        getItemById(props.id);
    }

    useEffect(() => {
        if (session?.accessToken) {
            setToken(session.accessToken)
        }
    }, [session])

    const onUpVote = () => {
        vote({ id:props.id, isUp:true })
    }

    const onDownVote = () => {
        vote({ id:props.id, isUp:false })
    }

    return (
        <Item elevation={3} sx={{px:1}}>
            <Grid container spacing={1}>
                <Grid item xs={3}>
                    <Content sx={{
                        textAlign: 'center',
                        fontSize: 20,
                        display: "flex",
                        justifyContent: "center",
                    }}>
                        <Stack spacing={0}>
                            <IconButton aria-label="up" onClick={onUpVote}>
                                <KeyboardArrowUpIcon />
                            </IconButton>
                            <b>{voteCount ?? 0}</b>
                            <IconButton aria-label="down" onClick={onDownVote}>
                                <KeyboardArrowDownIcon />
                            </IconButton>
                        </Stack>
                    </Content>
                </Grid>
                <Grid item xs={9}>
                    <Stack spacing={1}>
                        <Content sx={{
                            fontSize: 20,
                        }}>{props.itemName}</Content>
                        <Content sx={{
                            overflowY: "scroll",
                            height: 120,
                            fontSize: 14,
                        }}>{props.itemDescription}</Content>
                    </Stack>
                </Grid>
            </Grid>
        </Item>
    )
}