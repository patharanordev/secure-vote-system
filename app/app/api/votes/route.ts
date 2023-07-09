import { NextResponse } from "next/server";

export async function GET(req: Request) {
    try {
        const res = await fetch(`${process.env.API_HOST}/api/v1/votes`, {
            method: req.method,
            headers: req.headers
        })

        const data = await res.json();
        console.log('votes api:', data)

        if (data?.status === 401) {
            return NextResponse.redirect(`${process.env.BASE_HOST}/login`);
        } else {
            return new NextResponse(JSON.stringify(data));
        }

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