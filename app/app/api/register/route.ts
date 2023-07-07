import { NextResponse } from "next/server";

export async function POST(req: Request) {
    try {
        const body = await req.json()
        console.log('body :', body)
        const res = await fetch(`${process.env.API_HOST}/signup`, {
            method: req.method,
            body: JSON.stringify(body),
            headers: { "Content-Type": "application/json" }
        })

        const { error } = await res.json();
        let response = NextResponse.redirect(`${process.env.BASE_HOST}/login`);

        if (error) {
            response = new NextResponse(error)
        }

        return response;

    } catch (error: any) {
        console.log(error)
        return new NextResponse(
            JSON.stringify({
                status: "error",
                message: error.message,
            }),
            { status: 500 }
        );
    }
}