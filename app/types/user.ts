export type UserProfile = {
    name: string | null | undefined
    image: string | null | undefined
    email: string | null | undefined
}

export type Account = {
    profile: UserProfile,
    onAddItem?: Function
}