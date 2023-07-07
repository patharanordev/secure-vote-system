import { ThemeProvider } from "@mui/material"
import { SessionProvider } from "next-auth/react"
import { AppProps } from 'next/app'
import { theme } from '#/theme'

import CssBaseline from '@mui/material/CssBaseline';
import '../normalize.css'

export default function App({ 
  Component, 
  pageProps: { session, ...pageProps },
}: AppProps) {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <SessionProvider session={session}>
        <Component {...pageProps} />
      </SessionProvider>
    </ThemeProvider>
  )
}
