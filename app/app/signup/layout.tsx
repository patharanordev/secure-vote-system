import MainContainer from "#/ui/container";

export default async function RegisterLayout({
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
