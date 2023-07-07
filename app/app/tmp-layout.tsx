import { NextAuthProvider } from "./providers";
import ResponsiveAppBar from '#/ui/app-bar';
import MainContainer from "#/ui/container";

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en" className="">
      <body className="">
        <NextAuthProvider>
          <MainContainer>
            <ResponsiveAppBar />
            {children}
          </MainContainer>
        </NextAuthProvider>
      </body>
    </html>
  )
}
