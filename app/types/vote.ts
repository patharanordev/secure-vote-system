export type VoteItemProps = {
    id: string;
    itemName: string;
    itemDescription: string;
    userId: string;
    voteCount: number;
    onVoteSuccess?: Function;
}

export type VoteItemIDPayload = {
    id: string;
}

export type VoteListPayload = {
    status: number;
    error: string;
    data: Array<VoteItemProps>;
};

export type VoteItemPayload = {
    status: number;
    error: string;
    data: VoteInfoItem;
};

export type VotingPayload = {
    id: string;
    isUp: boolean;
};

export type VoteInfo = {
    itemDescription: string
    itemName: string
    voteCount?: number
};

export type VoteInfoItem = {
    vid: string
    info: VoteInfo
};