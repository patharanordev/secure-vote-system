import jwt from "jsonwebtoken";
import type { NextAuthOptions } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";

export const authOptions: NextAuthOptions = {
    session: {
        strategy: "jwt",
    },
    providers: [
        CredentialsProvider({
            name: "CustomProvider",
            credentials: {
                username: {
                    label: "Username",
                    placeholder: "your user name",
                },
                password: { label: "Password", type: "password" },
            },
            async authorize(credentials) {
                console.log('Gateway:', process?.env?.API_HOST)
                const res = await fetch(`${process?.env?.API_HOST}/login`, {
                    method: 'POST',
                    body: JSON.stringify(credentials),
                    headers: { "Content-Type": "application/json" }
                })

                const user = await res.json()
                // console.log('user:', user)

                // If no error and we have user data, return it
                if (res.ok && user) {
                    const decoded = jwt.verify(user.token ?? "", process.env.NEXTAUTH_SECRET as string);
                    // console.log('decoded:', decoded)

                    user['data'] = decoded;
                    
                    return user
                }
                // Return null if user data could not be retrieved
                return null
            },
        }),
    ],
    callbacks: {
        session: ({ session, token }) => {
            // console.log("Session Callback", { session, token });
            return {
                ...session,
                accessToken: token.accessToken,
                user: {
                    ...session.user,
                    id: token.id,
                    name: token.name,
                    isAdmin: token.isAdmin
                },
            };
        },

        // Ex. our user object
        // {
        //     id: '68b5adee-b583-4b25-b08c-9ed18daa3b8e',
        //     name: 'PatharaNor',
        //     admin: false,
        //     iss: 'SecVoteSys',
        //     sub: 'SecVoteSys_CustomAuth',
        //     aud: [ 'general_user' ],
        //     exp: 1688711970,
        //     nbf: 1688625570,
        //     iat: 1688625570,
        //     jti: '1'
        // }

        jwt: ({ token, user, account, profile }) => {
            // console.log("JWT Callback", { token, user, account, profile });
            if (user) {
                const u = user as unknown as any;
                const { id, name, admin, ...others } = u.data;
                token = Object.assign({}, token, { 
                    accessToken: u.token,
                    ...others,
                });

                return {
                    ...token,
                    id: id,
                    name: name,
                    isAdmin: admin,
                    accessToken: u.token
                };
            }
            return token;
        },
        async redirect({ url, baseUrl }) {
            // Allows relative callback URLs
            if (url.startsWith("/")) return `${baseUrl}${url}`
            // Allows callback URLs on the same origin
            else if (new URL(url).origin === baseUrl) return url
            return `${baseUrl}/login`
        }
    },
    pages: {
        signIn: '/login',
    },
    logger: {
        error(code, metadata) {
            console.error(code, metadata)
        },
        warn(code) {
            console.warn(code)
        },
        debug(code, metadata) {
            console.debug(code, metadata)
        }
    }
};