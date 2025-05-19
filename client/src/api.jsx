const API_BASE = import.meta.env.VITE_API_BASE || '/api'

export const todoAPI = {
  getAll: async () => {
    const response = await fetch(`${API_BASE}/todo`)
    if (!response.ok) throw new Error('Failed to fetch todos')
    return await response.json()
  },
  
  create: async (todo) => {
    const response = await fetch(`${API_BASE}/todo/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(todo),
    })
    if (!response.ok) throw new Error('Failed to create todo')
    return await response.json()
  },
  
  delete: async (id) => {
    const response = await fetch(`${API_BASE}/todo/${id}`, {
      method: 'DELETE',
    })
    if (!response.ok) throw new Error('Failed to delete todo')
    return await response.json()
  }
}