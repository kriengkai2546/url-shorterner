"use client"
import { useState, useEffect } from "react"
import { useRouter } from "next/navigation"
import { createURL, getMyURLs, deleteURL, API_URL } from "@/lib/api"

export default function DashboardPage() {
    const router = useRouter()
    const [urls, setURLs] = useState<any[]>([])
    const [longURL, setLongURL] = useState("")
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState("")
    const [newURL, setNewURL] = useState<any>(null)

    function getToken() {
        return localStorage.getItem("token") || ""
    }

    useEffect(() => {
        if (!getToken()) {
            router.push("/login")
            return
        }
        fetchURLs()
    }, [])

    async function fetchURLs() {
        try {
            const data = await getMyURLs(getToken())
            setURLs(data || [])
        } catch {
            router.push("/login")
        }
    }

    async function handleCreate(e: React.FormEvent) {
        e.preventDefault()
        setLoading(true)
        setError("")
        setNewURL(null)
        try {
            const data = await createURL(longURL, getToken())
            setNewURL(data)
            setLongURL("")
            fetchURLs()
        } catch (err: any) {
            setError(err.message)
        } finally {
            setLoading(false)
        }
    }

    async function handleDelete(id: number) {
        try {
            await deleteURL(id, getToken())
            fetchURLs()
        } catch (err: any) {
            setError(err.message)
        }
    }

    function handleLogout() {
        localStorage.removeItem("token")
        router.push("/login")
    }

    return (
        <div className="min-h-screen bg-gray-50 p-6">
            <div className="max-w-2xl mx-auto">
                <div className="flex justify-between items-center mb-6">
                    <h1 className="text-2xl font-bold">Dashboard</h1>
                    <button
                        onClick={handleLogout}
                        className="text-sm text-red-500 hover:underline"
                    >
                        Logout
                    </button>
                </div>

                {/* Create URL */}
                <div className="bg-white rounded-lg shadow p-6 mb-6">
                    <h2 className="font-medium mb-4">สร้าง Short URL</h2>
                    {error && <p className="text-red-500 text-sm mb-3">{error}</p>}
                    {newURL && (
                        <div className="bg-green-50 border border-green-200 rounded p-3 mb-3">
                            <p className="text-sm text-green-800 font-medium">สร้างสำเร็จ</p>
                            <a
                                href={`${API_URL}/${newURL.short_code}`}
                                target="_blank"
                                className="text-sm text-blue-600 hover:underline"
                            >
                                {`${API_URL}/${newURL.short_code}`}
                            </a>
                        </div>
                    )}
                    <form onSubmit={handleCreate} className="flex gap-2">
                        <input
                            type="url"
                            value={longURL}
                            onChange={e => setLongURL(e.target.value)}
                            placeholder="https://example.com/very-long-url"
                            className="flex-1 border rounded px-3 py-2 text-sm"
                            required
                        />
                        <button
                            type="submit"
                            disabled={loading}
                            className="bg-blue-600 text-white px-4 py-2 rounded text-sm font-medium hover:bg-blue-700 disabled:opacity-50"
                        >
                            {loading ? "..." : "Shorten"}
                        </button>
                    </form>
                </div>

                {/* URL List */}
                <div className="bg-white rounded-lg shadow p-6">
                    <h2 className="font-medium mb-4">URL ของฉัน</h2>
                    {urls.length === 0 ? (
                        <p className="text-sm text-gray-400">ยังไม่มี URL</p>
                    ) : (
                        <div className="space-y-3">
                            {urls.map(url => (
                                <div
                                    key={url.id}
                                    className="flex items-center justify-between border rounded p-3"
                                >
                                <div className="min-w-0 flex-1">
                                    <a
                                        href={`${API_URL}/${url.short_code}`}
                                        target="_blank"
                                        className="text-sm text-blue-600 font-medium hover:underline"
                                    >
                                        {API_URL}/{url.short_code}
                                    </a>
                                    <p className="text-xs text-gray-400 truncate mt-1">
                                        {url.long_url}
                                    </p>
                                </div>
                                    <button
                                        onClick={() => handleDelete(url.id)}
                                        className="text-xs text-red-500 hover:underline ml-4 flex-shrink-0"
                                    >
                                        ลบ
                                    </button>
                                </div>
                            ))}
                        </div>
                    )}
                </div>
            </div>
        </div>
    )
}