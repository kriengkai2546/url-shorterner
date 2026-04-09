const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080"

export async function register(email: string, password: string) {
    const res = await fetch(`${API_URL}/auth/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
    })
    if (!res.ok) throw new Error(await res.text())
    return res.json()
}

export async function login(email: string, password: string) {
    const res = await fetch(`${API_URL}/auth/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
    })
    if (!res.ok) throw new Error(await res.text())
    return res.json()
}

export async function createURL(longURL: string, token: string) {
    const res = await fetch(`${API_URL}/urls`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify({ long_url: longURL }),
    })
    if (!res.ok) throw new Error(await res.text())
    return res.json()
}

export async function getMyURLs(token: string) {
    const res = await fetch(`${API_URL}/my-urls`, {
        headers: { "Authorization": `Bearer ${token}` },
    })
    if (!res.ok) throw new Error(await res.text())
    return res.json()
}

export async function deleteURL(id: number, token: string) {
    const res = await fetch(`${API_URL}/urls/${id}`, {
        method: "DELETE",
        headers: { "Authorization": `Bearer ${token}` },
    })
    if (!res.ok) throw new Error(await res.text())
}