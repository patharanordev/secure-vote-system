import MainContainer from "#/ui/container";

export default async function LoginLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
        <MainContainer>
          {children}
        </MainContainer>
  )
}
