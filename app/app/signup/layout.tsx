import MainContainer from "#/ui/container";

export default async function RegisterLayout({
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
