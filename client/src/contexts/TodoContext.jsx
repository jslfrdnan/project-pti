import { createContext, useState, useEffect, useContext } from 'react'
import { todoAPI } from '../api'

const TodoContext = createContext()

export function TodoProvider({ children }) {
  const [todos, setTodos] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  const refreshTodos = async () => {
    try {
      setLoading(true)
      const response = await todoAPI.getAll()
      // Transform API response to match expected format
      const transformedTodos = response.map(item => ({
        id: item.data.id,
        title: `${item.data.todo}`,
        completed: false,
        originalData: item.data // Keep original data if needed
      }))
      setTodos(transformedTodos)
      setError(null)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const addTodo = async (todo) => {
    try {
      // Transform todo to match your API's expected format
      const apiTodo = {
        todo: todo.title
      }
      console.log('Sending to API:', apiTodo)
      const response = await todoAPI.create(apiTodo)
      setTodos([...todos, {
        id: response.data.id,
        title: `${response.data.todo}`,
        completed: false
      }])
    } catch (err) {
      setError(err.message)
    }
  }

  const deleteTodo = async (id) => {
    try {
      await todoAPI.delete(id)
      setTodos(todos.filter(todo => todo.id !== id))
    } catch (err) {
      setError(err.message)
    }
  }

  const toggleTodo = (id) => {
  setTodos(todos.map(todo => 
    todo.id === id ? { ...todo, completed: !todo.completed } : todo
  ));
  }

  useEffect(() => {
    refreshTodos()
  }, [])

  return (
    <TodoContext.Provider value={{
      todos,
      loading,
      error,
      addTodo,
      deleteTodo,
      toggleTodo,
      refreshTodos
    }}>
      {children}
    </TodoContext.Provider>
  )
}

export const useTodos = () => useContext(TodoContext)