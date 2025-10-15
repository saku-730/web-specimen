// middleware.ts

import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

// このconfigで、ミドルウェアをどのパスで実行するか指定するのだ
export const config = {
  matcher: [
    /*
     * 下記にマッチするパス以外、すべてのパスでミドルウェアを実行する
     * - api (API routes)
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico (favicon file)
     */
    '/((?!api|_next/static|_next/image|favicon.ico).*)',
  ],
};

export function middleware(request: NextRequest) {
  console.log(`[Middleware] middleware auth log: ${request.nextUrl.pathname}`);
  const { pathname } = request.nextUrl;

  // 1. /loginページ自体は、チェックの対象外にするのだ
  if (pathname.startsWith('/login')) {
    return NextResponse.next();
  }

  // 2. リクエストからトークンクッキーを取得する
  const token = request.cookies.get('token');

  // 3. トークンがなければ、/loginページに追い返す（リダイレクト）！
  if (!token) {
    const loginUrl = new URL('/login', request.url);
    return NextResponse.redirect(loginUrl);
  }

  // 4. トークンがあれば、何もしないでそのまま通すのだ
  return NextResponse.next();
}
