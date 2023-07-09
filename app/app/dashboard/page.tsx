"use client"

import { VoteInfo, VoteItemPayload } from "#/types";
import ResponsiveAppBar from "#/ui/app-bar";
import AddItemDialog from "#/ui/dialog/dlg-add-item";
import VoteList from "#/ui/vote/vote-list"
import { signIn } from "next-auth/react";
import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";

const DashboardPage = () => {

  const { data: session, status } = useSession();
  const [authStatus, setAuthStatus] = useState<string>(status)
  const [isOpenAddItem, setOpenAddItem] = useState<boolean>(false)
  const [token, setToken] = useState<string|null>(null);

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

  const onAddItem = () => {
    console.log('dashboard page, onAddItem :', true)
    setOpenAddItem(true)
  }

  const onSaveVoteInfo = (data: VoteInfo) => {
    console.log('Vote info:', data)
    const isError = createItem(data);

    if (!isError) {
      setOpenAddItem(false)
    }
  }

  const onCancelVoteInfo = (data: VoteInfo) => {
    console.log('Vote info:', data)
    setOpenAddItem(false)
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
          <VoteList />
        </div>
        <AddItemDialog 
          isOpen={isOpenAddItem}
          onSave={onSaveVoteInfo}
          onCancel={onCancelVoteInfo}
        />
      </>
    :
      "Loading"
  );
}

export default DashboardPage;