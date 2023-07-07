import MainContainer from "#/ui/container";

export default async function LoginLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en" className="">
      <body className="">
        <MainContainer>
          {children}
        </MainContainer>
      </body>
    </html>
  )
}
