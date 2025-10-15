// src/app/api/login/route.ts

import { NextResponse } from 'next/server';
import { cookies } from 'next/headers';

export async function POST(request: Request) {
  try {
    const body = await request.json();
    const { email, password } = body;

    // 1. Goのバックエンドにログイン情報を中継する
    const apiRes = await fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    });

    if (!apiRes.ok) {
      return NextResponse.json({ error: 'Authentication failed' }, { status: 401 });
    }

    const data = await apiRes.json();
    const token = data.token; // Goサーバーから返されたトークン

    // 2. トークンを安全なHttpOnlyクッキーに保存する！
    cookies().set('token', token, {
      httpOnly: true, // JavaScriptからアクセスできなくなり、安全！
      secure: process.env.NODE_ENV === 'production',
      maxAge: 60 * 60 * 24, // 1日間有効
      path: '/',
    });

    return NextResponse.json({ message: 'Login successful' });

  } catch (error) {
    return NextResponse.json({ error: 'Internal Server Error' }, { status: 500 });
  }
}
