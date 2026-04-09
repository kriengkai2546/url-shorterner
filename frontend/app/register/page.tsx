"use client"
import { useState } from "react"
import { useRouter } from "next/navigation"
import { register } from "@/lib/api"

export default function RegisterPage() {
    const router = useRouter()
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [error, setError] = useState("")
    const [loading, setLoading] = useState(false)

    async function handleSubmit(e: React.FormEvent) {
        e.preventDefault()
        setLoading(true)
        setError("")
        try {
            await register(email, password)
            router.push("/login")
        } catch (err: any) {
            setError(err.message)
        } finally {
            setLoading(false)
        }
    }

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-50">
            <div className="bg-white p-8 rounded-lg shadow w-full max-w-md">
                <h1 className="text-2xl font-bold mb-6">Register</h1>
                {error && <p className="text-red-500 mb-4 text-sm">{error}</p>}
                <form onSubmit={handleSubmit} className="space-y-4">
                    <div>
                        <label className="block text-sm font-medium mb-1">Email</label>
                        <input
                            type="email"
                            value={email}
                            onChange={e => setEmail(e.target.value)}
                            className="w-full border rounded px-3 py-2 text-sm"
                            required
                        />
                    </div>
                    <div>
                        <label className="block text-sm font-medium mb-1">Password</label>
                        <input
                            type="password"
                            value={password}
                            onChange={e => setPassword(e.target.value)}
                            className="w-full border rounded px-3 py-2 text-sm"
                            required
                        />
                    </div>
                    <button
                        type="submit"
                        disabled={loading}
                        className="w-full bg-blue-600 text-white py-2 rounded text-sm font-medium hover:bg-blue-700 disabled:opacity-50"
                    >
                        {loading ? "Loading..." : "Register"}
                    </button>
                </form>
                <p className="text-sm text-center mt-4">
                    มีบัญชีแล้ว? <a href="/login" className="text-blue-600">Login</a>
                </p>
            </div>
        </div>
    )
}