export type VoteItemProps = {
    id: string
    itemName: string
    itemDescription: string
    userId: string
    voteCount: number
}

export type VoteListPayload = {
    status: number;
    error: string;
    data: Array<VoteItemProps>;
};