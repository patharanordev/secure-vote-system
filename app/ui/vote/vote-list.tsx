"use client"

import { VoteListPayload, VoteItemProps } from "#/types"
import MainContainer from "#/ui/container"
import VoteItem from "#/ui/vote/vote-item"
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
        <>
            {payload?.data?.map((v: VoteItemProps, i: number) => (
                <VoteItem key={i} {...v} />
            ))}
        </>
    );
}

export default VoteList;