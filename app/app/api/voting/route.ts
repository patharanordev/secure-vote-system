import { NextResponse } from "next/server";

export async function PATCH(req: Request) {
    try {
        const body = await req.json()
        console.log('body :', body)
        const res = await fetch(`${process.env.API_HOST}/api/v1/voting`, {
            method: req.method,
            headers: req.headers,
            body: JSON.stringify(body)
        })

        const data = await res.json();
        console.log('voting api:', data)
        
        return new NextResponse(JSON.stringify(data));

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