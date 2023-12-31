"use client"

import { VoteListPayload, VoteItemProps, VoteInfo } from "#/types"
import VoteItem from "#/ui/vote/vote-item"
import { Grid } from "@mui/material";
import { useEffect, useState } from "react";

type Props = {
    list: VoteItemProps[]
    onVoteSuccess?: Function
    onClickEdit?: Function
    onClickDelete?: Function
}

const VoteList = (props: Props) => {
    const [result, setResult] = useState<VoteItemProps[]>([]);

    const onVoteSuccess = (vid: string) => {
        if (typeof props.onVoteSuccess === 'function') {
            props.onVoteSuccess(vid)
        }
    }

    const onClickEdit = (voteInfo: VoteInfo) => {
        if (typeof props.onClickEdit === 'function') {
            props.onClickEdit(voteInfo)
        }
    }

    const onClickDelete = (vid: string) => {
        if (typeof props.onClickDelete === 'function') {
            props.onClickDelete(vid)
        }
    }

    useEffect(() => {
        if (props.list && props.list.length > 0) {
            setResult(props.list)
        }
    }, [props.list])

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
            {result?.map((v: VoteItemProps, i: number) => (
                <Grid item key={`grid-${v.id}`} xs={12} sm={6} md={4}
                    sx={(theme) => ({
                        padding: 1,
                        [theme.breakpoints.down('sm')]: {
                            marginTop: 2,
                        }
                    })}
                >
                    <VoteItem key={v.id} {...v} 
                        onVoteSuccess={onVoteSuccess} 
                        onClickEdit={onClickEdit}
                        onClickDelete={onClickDelete}
                    />
                </Grid>
            ))}
        </Grid>
    );
}

export default VoteList;