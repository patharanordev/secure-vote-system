import { getToken } from "next-auth/jwt";
import { NextRequest, NextResponse } from "next/server";

export { default } from "next-auth/middleware";

export const config = {
  // matcher: ["/dashboard"],
  matcher: ["/((?!register|signup|api|login).*)"],
};

export async function middleware(request: NextRequest) {
  const token = await getToken({
    req: request,
    secret: process.env.NEXTAUTH_SECRET,
    cookieName: process.env.COOKIE_NAME,
  });

  // redirect user without access to login
  if (token?.token && Date.now() / 1000 < token?.exp) {
    return NextResponse.redirect("/login");
  }

  // // redirect user without admin access to login
  // if (!token?.isAdmin) {
  //   return NextResponse.redirect("/login");
  // }

  return NextResponse.next();
}