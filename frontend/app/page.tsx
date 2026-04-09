import Link from "next/link"

export default function Home() {
    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-50">
            <div className="text-center">
                <h1 className="text-4xl font-bold mb-4">URL Shortener</h1>
                <p className="text-gray-500 mb-8">สร้าง short URL ได้ง่ายๆ</p>
                <div className="flex gap-4 justify-center">
                    <Link
                        href="/login"
                        className="bg-blue-600 text-white px-6 py-2 rounded font-medium hover:bg-blue-700"
                    >
                        Login
                    </Link>
                    <Link
                        href="/register"
                        className="border border-blue-600 text-blue-600 px-6 py-2 rounded font-medium hover:bg-blue-50"
                    >
                        Register
                    </Link>
                </div>
            </div>
        </div>
    )
}