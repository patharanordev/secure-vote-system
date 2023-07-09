import { NextAuthProvider } from "../providers";
import ResponsiveAppBar from '#/ui/app-bar';
import MainContainer from "#/ui/container";

export default async function DashboardLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
        <NextAuthProvider>
          <MainContainer>
            <ResponsiveAppBar />
            {children}
          </MainContainer>
        </NextAuthProvider>
  )
}
