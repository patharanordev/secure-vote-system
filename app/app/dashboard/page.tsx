"use client"

import VoteList from "#/ui/vote/vote-list"
import { signIn } from "next-auth/react";
import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";

const DashboardPage = () => {

  const { data: session, status } = useSession();
  const [authStatus, setAuthStatus] = useState<string>(status)
  const [timer, setTimer] = useState<NodeJS.Timeout|null>(null)
  
  // useEffect(() => {
  //   const getStatus = () => authStatus
  //   const timeout = setTimeout(() => {
  //     console.log('session:', session)
  //     if (getStatus() === "loading") {
  //         console.log('redirect...')
  //         signIn();
  //     }
  //   }, 10000)
    
  //   setTimer(timeout)
    
  // }, [session])

  useEffect(() => {
      console.log('session status:', status)
      setAuthStatus(status)

      if (status === "unauthenticated") {
        signIn('credentials', { callbackUrl:'/dashboard' });
      } 
      // else if (status === "authenticated") {
      //   if (timer) {
      //     clearTimeout(timer)
      //     setTimer(null)
      //     console.log('cancel timer...')
      //   }
      // }
  }, [status])

  return (
    authStatus == "authenticated"
    ?
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <VoteList />
      </div>
    :
      "Loading"
  );
}

export default DashboardPage;