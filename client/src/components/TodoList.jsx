import { useTodos } from '../contexts/TodoContext'
import { useState } from 'react'

export default function TodoList() {
  const { todos, loading, error, addTodo, deleteTodo, toggleTodo } = useTodos()
  const [newTodo, setNewTodo] = useState('')

  const handleSubmit = (e) => {
    e.preventDefault()
    if (!newTodo.trim()) return
    addTodo({ title: newTodo, completed: false })
    setNewTodo('')
  }

  if (loading) return <div className="loading">Loading...</div>
  if (error) return <div className="error">Error: {error}</div>

  return (
    <div className="todo-container">
      <h1>Todo List</h1>
      
      <form onSubmit={handleSubmit} className="todo-form">
        <input
          type="text"
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
          placeholder="Add new todo..."
          className="todo-input"
        />
        <button type="submit" className="add-button">Add</button>
      </form>

      <ul className="todo-list">
        {todos.map(todo => (
          <li key={todo.id} className={`todo-item ${todo.completed ? 'completed' : ''}`}>
            <span 
              className="todo-title"
              style={{ textDecoration: todo.completed ? 'line-through' : 'none' }}
            >
              {todo.title}
            </span>
            <div className="todo-actions">
              <button 
                onClick={() => toggleTodo(todo.id)} 
                className="done-button"
              >
                {todo.completed ? 'Undo' : 'Done'}
              </button>
              <button 
                onClick={() => deleteTodo(todo.id)} 
                className="delete-button"
              >
                Delete
              </button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  )
}