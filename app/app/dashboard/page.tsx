"use client"

import { EModifyVoteInfoMode, VoteInfo, VoteItemIDPayload, VoteItemPayload, VoteItemProps, VoteListPayload } from "#/types";
import ResponsiveAppBar from "#/ui/app-bar";
import ModifyVoteItemDialog, { DEFAULT_VOTE_INFO } from "#/ui/dialog/dlg-modify-vote-item";
import VoteList from "#/ui/vote/vote-list"
import { signIn } from "next-auth/react";
import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";

const DashboardPage = () => {

  const { data: session, status } = useSession();
  const [authStatus, setAuthStatus] = useState<string>(status)
  const [isOpenModifyItem, setOpenModifyItem] = useState<boolean>(false)
  const [modeModifyItem, setModeModifyItem] = useState<EModifyVoteInfoMode>(EModifyVoteInfoMode.Add)
  const [token, setToken] = useState<string|null>(null);
  const [payload, setPayload] = useState<VoteListPayload|null>(null);
  const [defaultVoteInfo, setDefaultVoteInfo] = useState<VoteInfo>(DEFAULT_VOTE_INFO)

  const createItem = async (payload: VoteInfo) => {
    const res: VoteItemPayload = await fetch('/api/vote-item', {
        method: "POST",
        body: JSON.stringify({ ...payload }),
        headers: { 
            "Content-Type": "application/json", 
            "Authorization": `Bearer ${token}`
        }
    }).then((res) => res.json());

    console.log('Create vote item res:', res);

    return res.error
  }

  const editItem = async (payload: VoteInfo) => {
    const res: VoteItemPayload = await fetch('/api/vote-item', {
        method: "PATCH",
        body: JSON.stringify({ ...payload }),
        headers: { 
            "Content-Type": "application/json", 
            "Authorization": `Bearer ${token}`
        }
    }).then((res) => res.json());

    console.log('Edit vote item res:', res);

    return res.error
  }

  const deleteItem = async (payload: VoteItemIDPayload) => {
    const res: VoteItemPayload = await fetch('/api/vote-item', {
        method: "DELETE",
        body: JSON.stringify({ ...payload }),
        headers: { 
            "Content-Type": "application/json", 
            "Authorization": `Bearer ${token}`
        }
    }).then((res) => res.json());

    console.log('Delete vote item res:', res);

    return res.error
  }

  const load = async (token: string) => {
    const payload: VoteListPayload = await fetch('/api/votes', {
        method: "GET",
        headers: { "Authorization": `Bearer ${token}` }
    }).then((res) => res.json());

    // console.log('VoteList session:', token);
    console.log('VoteList payload:', payload);
    setPayload(payload);
}

  const onAddItem = () => {
    console.log('dashboard page, onAddItem :', true)
    setModeModifyItem(EModifyVoteInfoMode.Add)
    setOpenModifyItem(true)
  }

  const onSaveVoteInfo = async (data: VoteInfo) => {
    console.log('Save vote info:', data)

    let isError = null
    if (modeModifyItem === EModifyVoteInfoMode.Add) {
      isError = await createItem(data);
    } else if (modeModifyItem === EModifyVoteInfoMode.Edit) {
      isError = await editItem(data);
    }

    if (!isError) {
      setOpenModifyItem(false)
      load(token ?? "")
    }

    // Clear to default
    setModeModifyItem(EModifyVoteInfoMode.Add)
    setDefaultVoteInfo(DEFAULT_VOTE_INFO)
  }

  const onCancelVoteInfo = (data: VoteInfo) => {
    setOpenModifyItem(false)
  }

  const onClickEdit = async (voteInfo: VoteInfo) => {
    console.log('Edit vote info:', voteInfo)
    setDefaultVoteInfo(voteInfo)
    setModeModifyItem(EModifyVoteInfoMode.Edit)

    setTimeout(() => {
      setOpenModifyItem(true)
    }, 250)

    // const isError = await deleteItem({ props });

    // if (!isError) {
    //   load(token ?? "")
    // }
  }

  const onClickDelete = async (id: string) => {
    console.log('Delete vote info:', id)
    const isError = await deleteItem({ id });

    if (!isError) {
      load(token ?? "")
    }
  }

  const onVoteSuccess = (id: string) => {
    load(token ?? "")
  }

  useEffect(() => {
      console.log('session status:', status)
      setAuthStatus(status)

      if (status === "unauthenticated") {
        signIn('credentials', { callbackUrl:'/dashboard' });
      }
  }, [status])

  useEffect(() => {
    if (session?.accessToken) {
        setToken(session.accessToken)
        load(session.accessToken ?? "")
    }
  }, [session])

  return (
    authStatus == "authenticated"
    ?
      <>
        <ResponsiveAppBar onAddItem={onAddItem} />
        <div
          style={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <VoteList 
            list={ payload?.data ?? [] } 
            onVoteSuccess={onVoteSuccess}
            onClickEdit={onClickEdit}
            onClickDelete={onClickDelete}
          />
        </div>
        <ModifyVoteItemDialog 
          isOpen={isOpenModifyItem}
          onSave={onSaveVoteInfo}
          onCancel={onCancelVoteInfo}
          defaultProps={defaultVoteInfo}
        />
      </>
    :
      "Loading"
  );
}

export default DashboardPage;