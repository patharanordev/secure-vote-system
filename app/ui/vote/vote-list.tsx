"use client"

import { VoteListPayload, VoteItemProps } from "#/types"
import VoteItem from "#/ui/vote/vote-item"
import { Grid } from "@mui/material";
import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";

const VoteList = () => {
    const { data: session } = useSession();
    const [payload, setPayload] = useState<VoteListPayload|null>(null);

    const load = async (token: string) => {
        const payload: VoteListPayload = await fetch('/api/votes', {
            method: "GET",
            headers: { "Authorization": `Bearer ${token}` }
        }).then((res) => res.json());

        console.log('VoteList session:', token);
        console.log('VoteList payload:', payload);
        setPayload(payload);
    }

    useEffect(() => {
        if (session?.accessToken) {
            load(session.accessToken)
        }
    }, [session])

    return (
        <Grid container sx={(theme) => ({
            m: 2,
            marginTop: 20,
            
            [theme.breakpoints.down('md')]: {
                marginTop: 4
            },

            [theme.breakpoints.down('sm')]: {
                marginTop: 0,
                padding: 2
            }
        })}>
            {payload?.data?.map((v: VoteItemProps, i: number) => (
                <Grid item key={`grid-${v.id}`} xs={12} sm={6} md={4}
                    sx={(theme) => ({
                        padding: 1,
                        [theme.breakpoints.down('sm')]: {
                            marginTop: 2,
                        }
                    })}
                >
                    <VoteItem key={v.id} {...v} />
                </Grid>
            ))}
        </Grid>
    );
}

export default VoteList;