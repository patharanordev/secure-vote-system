import { NextAuthProvider } from "../providers";
import MainContainer from "#/ui/container";

export default async function DashboardLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
        <NextAuthProvider>
          <MainContainer>
            {children}
          </MainContainer>
        </NextAuthProvider>
  )
}
