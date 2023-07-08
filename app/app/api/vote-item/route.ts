import { NextResponse } from "next/server";

export async function GET(req: Request) {
    try {
        const query = req.url.split('?')[1] ?? '';
        const res = await fetch(`${process.env.API_HOST}/api/v1/vote-item?${query}`, {
            method: req.method,
            headers: req.headers,
        })

        const data = await res.json();
        console.log('vote-item api:', data)
        
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